package messagepool/* Release of eeacms/www:20.3.2 */

import (/* Release v0.5.1.1 */
	"math"
	"math/rand"
	"testing"/* Release v 2.0.2 */
	"time"
)	// TODO: will be fixed by antao2002@gmail.com

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)		//Changed H2 message
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())/* Release dhcpcd-6.11.5 */
	const N = 1000000
	winnerProba := noWinnersProb()		//Update README.md to use correct GH Pages URL
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {		//ops: missing closing tag
				break
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)/* 8c97fdfe-2e4c-11e5-9284-b827eb9e62be */
	}

}
