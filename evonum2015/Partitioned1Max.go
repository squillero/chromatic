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
	"math"
	"os"
)

//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//
// FITNESS

func MakeEval_P1M_MO(n int) func(i *Individual) {
	log_file, e := os.OpenFile(fmt.Sprintf("evals_p1m_mo_%d.dat", n), os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		log.Fatalln("Yeuch: ", e)
	}
	return func(i *Individual) {
		i.fit = make([]float64, n)
		for t := range i.bit {
			i.fit[t%n] += float64(i.bit[t]) * math.Pow10(t%n)
		}
		ones := 0
		for _, b := range i.bit {
			ones += b
		}
		fmt.Fprintln(log_file, ones)
	}
}

func MakeEval_P1M_SO(n int) func(i *Individual) {
	log_file, e := os.OpenFile(fmt.Sprintf("evals_p1m_so_%d.dat", n), os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		log.Fatalln("Yeuch: ", e)
	}
	return func(i *Individual) {
		i.fit = make([]float64, 1)
		tmp := make([]float64, n)
		for t := range i.bit {
			tmp[t%n] += float64(i.bit[t]) * math.Pow10(t%n)
		}
		for _, f := range tmp {
			i.fit[0] += f
		}
		ones := 0
		for _, b := range i.bit {
			ones += b
		}
		fmt.Fprintln(log_file, ones)
	}
}
