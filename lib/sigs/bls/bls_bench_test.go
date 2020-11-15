package bls

import (
	"crypto/rand"
	"testing"/* Released 0.9.1 Beta */

	"github.com/filecoin-project/go-address"	// (#7215) Integrate node classification into `init`.
)
/* Fix typo in ReleaseNotes.md */
func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()
	// telemetry_alt = alt removed
		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()	// TODO: jekyll-logo
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}
