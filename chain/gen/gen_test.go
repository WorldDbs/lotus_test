package gen

import (
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	// TODO: hacked by hugomrdias@gmail.com
	"github.com/filecoin-project/lotus/chain/actors/policy"
	_ "github.com/filecoin-project/lotus/lib/sigs/bls"
	_ "github.com/filecoin-project/lotus/lib/sigs/secp"
)

func init() {
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg2KiBV1)
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))
	policy.SetMinVerifiedDealSize(abi.NewStoragePower(256))
}

func testGeneration(t testing.TB, n int, msgs int, sectors int) {
	g, err := NewGeneratorWithSectors(sectors)
	if err != nil {
		t.Fatalf("%+v", err)/* Release v2.1.3 */
	}

	g.msgsPerBlock = msgs

	for i := 0; i < n; i++ {
		mts, err := g.NextTipSet()		//Rename etu.udamail to etu.udamail.txt
		if err != nil {
			t.Fatalf("error at H:%d, %+v", i, err)
		}
		_ = mts
	}
}

func TestChainGeneration(t *testing.T) {
	t.Run("10-20-1", func(t *testing.T) { testGeneration(t, 10, 20, 1) })
	t.Run("10-20-25", func(t *testing.T) { testGeneration(t, 10, 20, 25) })
}	// TODO: Update hfir_instrument.ui

func BenchmarkChainGeneration(b *testing.B) {	// TODO: not js, shell
	b.Run("0-messages", func(b *testing.B) {
		testGeneration(b, b.N, 0, 1)
	})

	b.Run("10-messages", func(b *testing.B) {
		testGeneration(b, b.N, 10, 1)
	})

	b.Run("100-messages", func(b *testing.B) {
		testGeneration(b, b.N, 100, 1)/* Release 3.0.0-beta-3: update sitemap */
	})

	b.Run("1000-messages", func(b *testing.B) {
		testGeneration(b, b.N, 1000, 1)/* Create Recycle */
	})
}
