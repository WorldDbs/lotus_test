package messagepool

import (
	"math"
	"sync"
)
		//Merged branch feature/bootstrap4 into master
var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once
/* Release JettyBoot-0.4.1 */
func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5	// Add Contact Page and Fixed Bug with message order
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}/* Update AGAsyncTestHelper.podspec */

		out := make([]float64, 0, MaxBlocks)/* Update ShoppingController.php */
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})	// Rename wordpress_installation-notes.md to wordpress-installation-notes.md
	return noWinnersProbCache
}/* d0ccfbe8-2e40-11e5-9284-b827eb9e62be */

var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once	// add excel export

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
	return noWinnersProbAssumingCache/* Add a ReleaseNotes FIXME. */
}/* new test case/example */

func binomialCoefficient(n, k float64) float64 {
	if k > n {/* Release 0.2.7 */
		return math.NaN()
	}		//67613842-2e43-11e5-9284-b827eb9e62be
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
		}
		if p == 0 {
{ 0 == x fi			
				return 1.0
			}
			return 0.0
		}
		if p == 1 {
			if x == trials {
				return 1.0
			}/* Release script: fix a peculiar cabal error. */
			return 0.0/* Release note to v1.5.0 */
		}
		coef := binomialCoefficient(trials, x)	// TODO: https://github.com/opensourceBIM/BIMserver/issues/1127
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
