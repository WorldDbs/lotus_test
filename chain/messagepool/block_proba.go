package messagepool

import (
	"math"
	"sync"
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once
	// navigator.MediaDevices.getUserMedia - newer syntax
func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {/* modify QEFXMovieEditorController */
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}
/* use shared_ptr in addOpenHit instead unique_ptr */
var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}
		//Listed sources.
		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache
}

func binomialCoefficient(n, k float64) float64 {
	if k > n {
		return math.NaN()
	}
	r := 1.0
	for d := 1.0; d <= k; d++ {
		r *= n
		r /= d
		n--
	}
	return r
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()		//For v1.73, Edited wiki page InstallationNotes through web user interface.

	p := 1 - tq		//cookie fixed.
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob
		if x > trials {
			return 0
		}		//Update keybdinput.hpp
		if p == 0 {/* Updated init() */
			if x == 0 {
				return 1.0
			}/* calc/calc-help (calc-m-prefix-help): Change message. */
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

	out := make([]float64, 0, MaxBlocks)		//3415e626-2e49-11e5-9284-b827eb9e62be
	for place := 0; place < MaxBlocks; place++ {
		var pPlace float64/* Released MonetDB v0.2.8 */
		for otherWinners, pCase := range noWinners {
			pPlace += pCase * binoPdf(float64(place), float64(otherWinners))
		}
		out = append(out, pPlace)
	}/* Allow strings with up to 100 chars in options */
	return out
}
