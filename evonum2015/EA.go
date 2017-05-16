//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\# -*-go-*-
//              //                                                          //
//   #####      //   Chromatic Selection                                    //
//  ######      //   by Giovanni Squillero <giovanni.squillero@polito.it>   //
//  ###   \     //                                                          //
//   ##G  c\    //   An Oversimplified approach to exploit comparison-based //
//   #     _\   //   optimizers for commensurable multi-objective problems  //
//   |  _/      //   See the paper @ EvoNUM <http://www.evostar.org/2015/>  //
//              //                                                          //
//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//
//
// Comments and criticisms (either constructive or not) are always welcomed!
//
//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//
//
// Version 0.1 (winter 2014)
// See <https://github.com/squillero/chromatic> for the latest version

package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//
// INDIVIDUALS

type Individual struct {
	bit   []int
	fit   []float64
	score float64
}

func (i Individual) String() string {
	s := fmt.Sprintf("%3d) ", i.score)
	for b := range i.bit {
		s += fmt.Sprintf("%d", i.bit[b])
	}
	s += fmt.Sprintf(" / %v", i.fit)
	return s
}

func RandomIndividual(b int) *Individual {
	var i Individual

	i.bit = make([]int, b)
	for t := range i.bit {
		i.bit[t] = rand.Int() % 2
	}
	i.fit = nil
	i.score = 0

	return &i
}

func Mutate(i *Individual) Individual {
	var o Individual
	o.bit = make([]int, len(i.bit))
	copy(o.bit, i.bit)
	b := rand.Int() % len(o.bit)
	o.bit[b] = (o.bit[b] + 1) % 2
	return o
}

func Crossover(i1, i2 *Individual) Individual {
	var o Individual
	o.bit = make([]int, len(i1.bit))
	for b := range o.bit {
		if rand.Float64() < 0.5 {
			o.bit[b] = i1.bit[b]
		} else {
			o.bit[b] = i2.bit[b]
		}
	}

	return o
}

//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//
// POPULATION

type Population struct {
	bits  int
	eval  func(i *Individual)
	ind   []*Individual
	score func(p *Population)
}

func (p *Population) Init(b int, e func(i *Individual), s func(p *Population)) {
	p.eval = e
	p.score = s
	p.bits = b
	p.ind = make([]*Individual, MU)
	for t := range p.ind {
		p.ind[t] = RandomIndividual(p.bits)
	}
	p.Eval()
}

func (p *Population) Eval() {
	for i := range p.ind {
		if p.ind[i].fit == nil {
			p.eval(p.ind[i])
		}
	}
	p.score(p)
	sort.Sort(p)
}

func (p *Population) Begat() {
	p.ind = p.ind[0:MU]
	for t := 0; t < LAMBDA; t++ {
		p1 := rand.Int() % MU
		p2 := rand.Int() % MU
		o := Crossover(p.ind[p1], p.ind[p2])
		om := Mutate(&o)
		p.ind = append(p.ind, &om)
	}
	p.Eval()
}

func (p *Population) Debug() {
	log.Println("Dumping population:")
	for _, i := range p.ind {
		log.Println("+  ", i)
	}
}

func (p *Population) Len() int {
	return len(p.ind)
}
func (p *Population) Swap(a, b int) {
	p.ind[a], p.ind[b] = p.ind[b], p.ind[a]
}
func (p *Population) Less(a, b int) bool {
	return p.ind[b].score < p.ind[a].score
}
func Score_Fit0(p *Population) {
	for i := range p.ind {
		p.ind[i].score = p.ind[i].fit[0]
	}
}
func Score_Chromatic(p *Population) {
	for i := range p.ind {
		score := 0
		for i2 := 0; i2 < GAMMA; i2++ {
			c := rand.Int() % len(p.ind)
			res := CCompare(p.ind[i].fit, p.ind[c].fit)
			if res > 0 {
				score += 3
			} else if res == 0 {
				score += 1
			}
		}
		p.ind[i].score = float64(score)
	}
}
