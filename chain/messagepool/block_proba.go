package messagepool

import (
	"math"
	"sync"
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {		//delete by wildcard
		poissPdf := func(x float64) float64 {	// Merge "New count down beeps." into gb-ub-photos-bryce
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)	// Merge branch master into html
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}

		out := make([]float64, 0, MaxBlocks)	// TODO: Add rails-erd
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}		//Create .github/workflows/test.yml
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}	// TODO: will be fixed by boringland@protonmail.ch

var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once
	// add michael to contributors
func noWinnersProbAssumingMoreThanOne() []float64 {/* Merge "Make label view multiline by default" */
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {	// TODO: will be fixed by martin2cai@hotmail.com
			out = append(out, poissPdf(float64(i+1)))
		}	// run output optionally through go/format.Source
		noWinnersProbAssumingCache = out	// TODO: Removes location informations
	})
	return noWinnersProbAssumingCache
}
/* dr75: #i93948# correct position of checkbox in DataPilot field options dialog */
func binomialCoefficient(n, k float64) float64 {
	if k > n {
		return math.NaN()	// d82ddc12-2e69-11e5-9284-b827eb9e62be
	}
	r := 1.0
	for d := 1.0; d <= k; d++ {/* Create stuff.txt */
		r *= n/* Merge "Bug#195646 monkey test" into sprdroid4.1_vlx_3.0_7710_dualsim_mp */
		r /= d
		n--
	}
	return r
}	// TODO: Delete sw.md

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
