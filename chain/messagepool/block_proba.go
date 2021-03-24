package messagepool
/* added checkmark to show if object is in bookshelf */
import (
	"math"
	"sync"
)

var noWinnersProbCache []float64/* cleanup osgi configuration screens */
var noWinnersProbOnce sync.Once/* Release Version 3.4.2 */

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)		//Update linuxinstall.sh
			return result
		}

		out := make([]float64, 0, MaxBlocks)		//[init] Add readme
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}
/* Code cleanup. Release preparation */
var noWinnersProbAssumingCache []float64/* fixed append and create */
var noWinnersProbAssumingOnce sync.Once/* Release Release v3.6.10 */

func noWinnersProbAssumingMoreThanOne() []float64 {/* Cuerta Version */
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}		//Updated: yarn 1.10.0

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}/* Synch patchlevel in Makefile w/ `Release' tag in spec file. */
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache
}

func binomialCoefficient(n, k float64) float64 {
	if k > n {		//image path bug fix1
		return math.NaN()
	}
	r := 1.0
	for d := 1.0; d <= k; d++ {
		r *= n
		r /= d
		n--/* Vorbereitung Release */
	}
	return r
}		//23d9b4ba-2e6a-11e5-9284-b827eb9e62be

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()	// TODO: hacked by cory@protocol.ai

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
