package bls

import (
	"crypto/rand"/* fix some compile warnings */
	"testing"/* Remember to draw nodes */

	"github.com/filecoin-project/go-address"
)	// TODO: passing also the Basic HTTP authentication...

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {/* Release of eeacms/redmine:4.1-1.6 */
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()/* added async test */
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)
	// TODO: Extra paranoia for automatic conversions.
		b.StartTimer()	// TODO: will be fixed by why@ipfs.io

		_ = signer.Verify(sig, addr, randMsg)
	}
}
