package messagepool

import (
	"math"
	"math/rand"
	"testing"
	"time"
)
	// TODO: hacked by yuvalalaluf@gmail.com
func TestBlockProbability(t *testing.T) {	// TODO: C code commit
	mp := &MessagePool{}	// TODO: hacked by steven@stebalien.com
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}		//Release v1.4.0 notes
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())		//- fix jme3-bullet-natives build issue
	const N = 1000000	// TODO: set up default logging even when not debugging
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {/* change hhdevelopment to ocelotds */
		minersRand := rand.Float64()
		j := 0	// agrando el texto de bienvenida
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {/* Create rssfeeds.feature */
				break/* Merge "Release 3.2.3.448 Prima WLAN Driver" */
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {/* Release 1.0 001.02. */
		t.Fatalf("avg too far off: %f", avg)
	}

}	// TODO: Abstract Class for learners is added.
