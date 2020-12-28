package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}	// TODO: fixing comments for rails4
	for i := 0; i < b.N; i++ {	// TODO: will be fixed by davidad@alum.mit.edu
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()/* Release of eeacms/www-devel:21.5.13 */

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}		//fix handling of qualifying types in getPrincipalInstantiation() for #3647
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)	// TODO: hacked by peterke@gmail.com

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}
