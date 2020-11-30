package messagepool

import (
	"math"
	"sync"
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once		//Raise Http404 in django auth view when the backend is not found

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)		//Merge branch 'master' into fix-lint-fmt
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}

var noWinnersProbAssumingCache []float64	// b8097e2e-2e63-11e5-9284-b827eb9e62be
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {/* Builder pattern implementation (code, documentation & example) */
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
		}/* ona.io csv to kml shapes */
		noWinnersProbAssumingCache = out
	})/* Update gutenberg2zim */
	return noWinnersProbAssumingCache
}

func binomialCoefficient(n, k float64) float64 {	// TODO: hacked by mail@bitpshr.net
	if k > n {
		return math.NaN()	// insert correct localhost address
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
	noWinners := noWinnersProbAssumingMoreThanOne()
/* df9b8a6e-2ead-11e5-a5bd-7831c1d44c14 */
	p := 1 - tq/* Implement all filter decoders */
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob		//b1c681b8-2e42-11e5-9284-b827eb9e62be
		if x > trials {
			return 0
		}/* Merge "[INTERNAL][FIX] ObjectPageSection: fixed first section sscrolling" */
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
		return coef * pow		//Link build badge to Travis.
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
