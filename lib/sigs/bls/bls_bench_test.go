package bls

import (
	"crypto/rand"
	"testing"
	// TODO: will be fixed by why@ipfs.io
	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()/* Fix issue #26: Also look at superclasses and interfaces */
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)/* .gitignore ignore newly generated doxygen folders */
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)/* Merge "Release 3.0.10.034 Prima WLAN Driver" */
	}/* Release version 0.12 */
}
	// Update SI4707.ino
func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)/* 2d0dc4d4-2e76-11e5-9284-b827eb9e62be */
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()/* See Releases */

		_ = signer.Verify(sig, addr, randMsg)
	}
}
