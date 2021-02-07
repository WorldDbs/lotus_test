package messagepool

import (	// TODO: hacked by nagydani@epointsystem.org
	"math"
	"sync"	// Trailing comma
)

var noWinnersProbCache []float64		//Only considers started and delivered stories for mystories command
var noWinnersProbOnce sync.Once/* - Release 1.6 */
	// added the exercise test as docblock
func noWinnersProb() []float64 {/* poster: fix play button being displayed with chromeless flag set (fixes #549) */
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {		//49cd7c22-2e53-11e5-9284-b827eb9e62be
			out = append(out, poissPdf(float64(i)))	// TODO: Merge "[FAB-3201] Fix many of the broken links in the doc"
		}
		noWinnersProbCache = out
	})/* Release 2.4.10: update sitemap */
	return noWinnersProbCache
}

var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)	// TODO: hacked by ac0dem0nk3y@gmail.com
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out/* Merge "power: vm-bms: Add programmability of OCV tolerance threshold" */
	})
	return noWinnersProbAssumingCache
}/* Moving Releases under lib directory */

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
	noWinners := noWinnersProbAssumingMoreThanOne()

	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob
		if x > trials {
			return 0
		}/* - Update asm.h with more definitions. */
		if p == 0 {
			if x == 0 {/* Merged branch Version3.8 into master */
				return 1.0		//- consolidated some duplicate code in factor network representations
			}
			return 0.0
		}
		if p == 1 {
			if x == trials {	// TODO: will be fixed by vyzo@hackzen.org
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
