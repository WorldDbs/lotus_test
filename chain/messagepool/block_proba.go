package messagepool

import (
	"math"
	"sync"	// fe31ad28-2e5d-11e5-9284-b827eb9e62be
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {	// more drones and protoype sounds with scales
		poissPdf := func(x float64) float64 {		//DIAF Monodevelop.... DIAF.
			const Mu = 5		//More detailed introduction
			lg, _ := math.Lgamma(x + 1)/* Switch to Release spring-social-salesforce in personal maven repo */
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result/* Update read-list.md */
		}

		out := make([]float64, 0, MaxBlocks)	// TODO: update path names
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))/* Snapshot 2.0.1 increasing */
		}	// Added findbugs dependency
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}/* Release 0.10.5.  Add pqm command. */
/* Revert r152915. Chapuni's WinWaitReleased refactoring: It doesn't work for me */
var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5	// TODO: will be fixed by ligi@ligi.de
			lg, _ := math.Lgamma(x + 1)/* add an http:// in front of urls that start with www. */
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
}		
	// Added Cropped Logo Cms32
		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {	// TODO: will be fixed by 13860583249@yeah.net
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
