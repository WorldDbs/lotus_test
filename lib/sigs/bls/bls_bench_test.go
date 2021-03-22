package bls
	// TODO: edit for clarity
import (/* Merge "Merge "ASoC: msm: qdsp6v2: Release IPA mapping"" */
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"		//Disable spurious warning
)

func BenchmarkBLSSign(b *testing.B) {		//Changelog - Mise en forme et complements
	signer := blsSigner{}/* Update Release_Procedure.md */
	for i := 0; i < b.N; i++ {
		b.StopTimer()/* ## 0.2.31-SNAPSHOT */
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)		//Typo isnt -> isn't
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {	// 481a65b0-2e1d-11e5-affc-60f81dce716c
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()/* More changes to draft */
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}
