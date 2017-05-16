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
	"os"
)

const ()

//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//
// FITNESS

func MakeEval_MK_MO(n int) func(i *Individual) {
	Value := make([][MK_BITS]float64, n)
	var Weight [MK_BITS]float64
	var totWeight float64
	rand.Seed(int64(n))
	for i := 0; i < n; i++ {
		Weight[i] = rand.Float64()
		totWeight += Weight[i]
		for j := 0; j < MK_BITS; j++ {
			Value[i][j] = 1000 * rand.Float64() / float64(n)
		}
	}
	MaxWeight := totWeight/2 + totWeight*rand.Float64()/2

	logFile, e := os.OpenFile(fmt.Sprintf("evals_mk_mo_%d.dat", n), os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		log.Fatalln("Yeuch: ", e)
	}

	return func(i *Individual) {
		i.fit = make([]float64, n)
		var w float64
		for b := range i.bit {
			if i.bit[b] > 0 {
				w += Weight[b]
				for f := range Value {
					i.fit[f] += Value[f][b]
				}
			}
		}
		if w > MaxWeight {
			for f := range i.fit {
				i.fit[f] = 0
			}
		}
		var realFitness float64
		for f := range i.fit {
			realFitness += i.fit[f]
		}
		fmt.Fprintln(logFile, realFitness)
	}
}

func MakeEval_MK_SO(n int) func(i *Individual) {
	Value := make([][MK_BITS]float64, n)
	var Weight [MK_BITS]float64
	var totWeight float64
	rand.Seed(int64(n))
	for i := 0; i < n; i++ {
		Weight[i] = rand.Float64()
		totWeight += Weight[i]
		for j := 0; j < MK_BITS; j++ {
			Value[i][j] = 1000 * rand.Float64() / float64(n)
		}
	}
	MaxWeight := totWeight/2 + totWeight*rand.Float64()/2

	logFile, e := os.OpenFile(fmt.Sprintf("evals_mk_so_%d.dat", n), os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		log.Fatalln("Yeuch: ", e)
	}

	return func(i *Individual) {
		i.fit = make([]float64, 1)
		tmp := make([]float64, n)
		var w float64
		for b := range i.bit {
			if i.bit[b] > 0 {
				w += Weight[b]
				for f := range Value {
					tmp[f] += Value[f][b]
				}
			}
		}
		if w > MaxWeight {
			for f := range tmp {
				tmp[f] = 0
			}
		}
		for f := range tmp {
			i.fit[0] += tmp[f]
		}
		fmt.Fprintln(logFile, i.fit[0])
	}
}

//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//
// SORT ON "SCORE"

func Score_MK_Base(p *Population) {
	for i := range p.ind {
		p.ind[i].score = 0
		for _, f := range p.ind[i].fit {
			p.ind[i].score += f
		}
	}
}
