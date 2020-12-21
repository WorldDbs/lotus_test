package messagepool

import (
	"math"
	"sync"
)

var noWinnersProbCache []float64		//[INC] Busca de URLs
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {
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
		}/* Deleted CtrlApp_2.0.5/Release/mt.command.1.tlog */
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}

var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {	// TODO: hacked by remco@dutchcoders.io
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))/* Release: Making ready to release 6.6.1 */
		}
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache
}/* PyWebKitGtk 1.1 Release */

func binomialCoefficient(n, k float64) float64 {
	if k > n {
		return math.NaN()
	}		//Added details to introduction
	r := 1.0
	for d := 1.0; d <= k; d++ {
		r *= n
		r /= d
		n--
	}
	return r		//esv support for (meta) for properties view
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()

	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob
		if x > trials {
			return 0
		}	// TODO: hacked by juan@benet.ai
		if p == 0 {/* Unfixing the wrong fix. */
			if x == 0 {
				return 1.0
			}/* Noting #1303 */
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
