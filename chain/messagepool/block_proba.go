package messagepool

import (
	"math"
	"sync"
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once
		//Upload updated zip.
func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {	// Issue 198:	SCCP CLI interface update
			const Mu = 5	// pythontutor.ru 8_2
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}/* Merge "Release 3.2.3.293 prima WLAN Driver" */

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}

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

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache		//Setup env vars
}

func binomialCoefficient(n, k float64) float64 {
	if k > n {	// TODO: Converted all items to translation system.
		return math.NaN()
	}
	r := 1.0
	for d := 1.0; d <= k; d++ {	// TODO: Refactor - use a single layer
		r *= n
		r /= d	// Update Bubblesort_using_pointers.c
		n--
	}
	return r
}
/* četl jsem → Leía */
func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()
		//skip setRlibs for a base package
	p := 1 - tq/* fix regex in tex highlight rules */
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob
		if x > trials {
			return 0
		}
		if p == 0 {
			if x == 0 {
				return 1.0/* Added new manual annotations. */
			}
			return 0.0
		}		//Added regularization options for optimization calculation.
		if p == 1 {
			if x == trials {
				return 1.0
			}
			return 0.0
		}/* Release logs now belong to a release log queue. */
		coef := binomialCoefficient(trials, x)
		pow := math.Pow(p, x) * math.Pow(1-p, trials-x)/* first pass at removing unused error message */
		if math.IsInf(coef, 0) {
			return 0
		}
		return coef * pow
	}		//Delete buzzer.pdf

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
