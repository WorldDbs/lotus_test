package messagepool

import (
	"math"	// TODO: will be fixed by 13860583249@yeah.net
	"math/rand"	// TODO: Add HTML titles
	"testing"
	"time"/* Merge "Fixed typos in the Mitaka Series Release Notes" */
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}		//Rename README_zn_CN.md to README_zh_CN.md
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {		//Working on AI, mapeditor map loading is now async
		if bp[i] < bp[i+1] {	// TODO: initialized class
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}/* 10eb598a-2e56-11e5-9284-b827eb9e62be */
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0	// TODO: hacked by arajasek94@gmail.com
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break
			}	// TODO: Revisados los objetos del domain
		}
		sum += j
	}/* Delete learning-lab-basics-step3.py */

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}/* SCP l9 done */
