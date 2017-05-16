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
	"log"
	"math/rand"
	"time"
)

const (
	MU     = 30
	LAMBDA = 20
	GAMMA  = 10 * (MU + LAMBDA)
	TAU    = 3

	MAX_GENERATIONS = 1000
	P1M_BITS        = 1000
	MK_BITS         = 1000
)

func main() {
	log.SetFlags(0)
	log.Println("Chromatic Selection -- Oversimplified multi-objective optimization")
	log.Println("(!) winter 2014 by Giovanni Squillero <giovanni.squillero@polito.it>")
	log.Println("")
	log.SetFlags(log.Lmicroseconds)

	s := time.Now().UTC().UnixNano()
	rand.Seed(s)
	log.Println("Random seed =", s)

	// MULTI-OBJECTIVE KNAPSACKS

	testMK := []int{1, 2, 5, 10, 20, 50, 100}

	for _, step := range testMK {
		var pop Population
		log.Println("Running Multi-objective knapsack/aggregate fitness", step)
		pop.Init(MK_BITS, MakeEval_MK_SO(step), Score_Fit0)
		for g := 0; g < MAX_GENERATIONS; g++ {
			pop.Begat()
		}
	}
	for _, step := range testMK {
		var pop Population
		log.Println("Running Multi-objective knapsack/chromatic", step)
		pop.Init(MK_BITS, MakeEval_MK_MO(step), Score_Chromatic)
		for g := 0; g < MAX_GENERATIONS; g++ {
			pop.Begat()
		}
	}

	// PARTITIONED 1-MAX TEST

	testP1M := []int{1, 2, 5, 10, 20, 50, 100}

	for _, step := range testP1M {
		var pop Population
		log.Println("Running experiment Partitioned 1-Max/aggregate fitness", step)
		pop.Init(P1M_BITS, MakeEval_P1M_SO(step), Score_Fit0)
		for g := 0; g < MAX_GENERATIONS; g++ {
			pop.Begat()
		}
	}
	for _, step := range testP1M {
		var pop Population
		log.Println("Running experiment Partitioned 1-Max/chromatic", step)
		pop.Init(P1M_BITS, MakeEval_P1M_MO(step), Score_Chromatic)
		for g := 0; g < MAX_GENERATIONS; g++ {
			pop.Begat()
		}
	}
}
