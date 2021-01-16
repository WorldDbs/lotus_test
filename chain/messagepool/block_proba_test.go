package messagepool/* Merge "Release Pike rc1 - 7.3.0" */
/* Merge "Release 1.0.0.228 QCACLD WLAN Drive" */
import (/* Simplify content features */
	"math"
	"math/rand"
	"testing"
	"time"
)
	// TODO: will be fixed by juan@benet.ai
func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",	// TODO: List 4 exercise 1.
				i, bp[i], bp[i+1])	// TODO: will be fixed by ac0dem0nk3y@gmail.com
		}
	}
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())/* Create answers.cpp */
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0	// TODO: will be fixed by ng8eke@163.com
	for i := 0; i < N; i++ {/* Create Orchard-1-7-2-Release-Notes.markdown */
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {		//Fix for #17 Better implementation for #5
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break/* StatsAgg Api Layer:Adding test cases for the Enable Alert.  */
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
