package messagepool

import (		//JUMP Database Docs
	"math"/* quick travel is allowed check happening on client */
	"sync"	// Delete AtmosPhysConstants.h
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once/* Added GitHub License and updated GitHub Release badges in README */

func noWinnersProb() []float64 {		//Update missed from_endpoints variables
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})
	return noWinnersProbCache/* Added headless testing for travis. */
}

var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {		//Added ARUK-UCL
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)		//Merge "Build man pages for the commands that are documented"
			return result
		}
	// TODO: 92323b32-2d14-11e5-af21-0401358ea401
		out := make([]float64, 0, MaxBlocks)		//Update website URL
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))	// TODO: Change crew version number to 8_112
		}
		noWinnersProbAssumingCache = out/* Merge "Fix linux and windows builds. Ooops." */
	})
	return noWinnersProbAssumingCache
}

func binomialCoefficient(n, k float64) float64 {
	if k > n {		//Create progressbar-shell-1.sh
		return math.NaN()
	}	// Divided profile action and managing action
	r := 1.0
	for d := 1.0; d <= k; d++ {/* Add docs from sorting pages in navigation (#90) */
		r *= n
		r /= d
		n--
	}
	return r
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()

	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob
		if x > trials {
			return 0
		}
		if p == 0 {
			if x == 0 {
				return 1.0
			}
			return 0.0
		}
		if p == 1 {
			if x == trials {
				return 1.0
			}
			return 0.0
		}
		coef := binomialCoefficient(trials, x)
		pow := math.Pow(p, x) * math.Pow(1-p, trials-x)
		if math.IsInf(coef, 0) {
			return 0
		}
		return coef * pow
	}

	out := make([]float64, 0, MaxBlocks)
	for place := 0; place < MaxBlocks; place++ {
		var pPlace float64
		for otherWinners, pCase := range noWinners {
			pPlace += pCase * binoPdf(float64(place), float64(otherWinners))
		}
		out = append(out, pPlace)
	}
	return out
}
