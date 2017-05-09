//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\# -*-go-*-
//              //                                                          //
//   #####      //   Chromatic Selection                                    //
//  ######      //   by Giovanni Squillero <giovanni.squillero@polito.it>   //
//  ###   \     //                                                          //
//   ##G  c\    //   An Oversimplified approach to exploit comparison-based //
//   #     _\   //   optimizers for commensurable multi-objective problems  //
//   |  _/      //   https://github.com/squillero/chromatic                 //
//              //                                                          //
//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//
// This code is licensed under a BSD license
//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//
// VERSION HISTORY
// - v0.1 : winter 2014

package main

import (
	"log"
	"math"
	"math/rand"
)

//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//
// CHROMATIC COMPARISON CORE
// Notez bien: this code is *not* optimized (indeed, it should be considered
// more pseudo-code, than code).
// Comments and criticisms (either constructive or not) are welcomed

func CCompare(fit1, fit2 []float64) int {
	chroma := make([]float64, len(fit1))
	blues := 0.0

	for i := range fit1 {
		if fit1[i] == 0 && fit2[i] == 0 {
			chroma[i] = 0
		} else {
			delta := math.Max(fit1[i], fit2[i]) - math.Min(fit1[i], fit2[i])
			norm := math.Max(fit1[i], fit2[i])
			chroma[i] = delta / norm
		}
		blues += chroma[i]
	}
	var c int
	if blues == 0 {
		return 0
	}
	blues *= rand.Float64()
	c = -1
	for blues >= 0 {
		c++
		blues -= chroma[c]
	}

	// paranoia check
	if c < 0 {
		log.Fatalln("PANIC: CC underflowed", c, chroma)
	}

	if fit1[c] > fit2[c] {
		return +1
	} else if fit1[c] < fit2[c] {
		return -1
	} else {
		return 0
	}
}
