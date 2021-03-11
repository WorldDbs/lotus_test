package bls

import (		//[Minor] removed requirement for PrivilegeAdmin role in priv. handler
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"	// TODO: will be fixed by lexy8russo@outlook.com
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}/* Release 1.15rc1 */
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)	// TODO: hacked by boringland@protonmail.ch
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}	// TODO: hacked by martin2cai@hotmail.com
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)		//template missing question mark
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()/* Release v0.5.1.3 */
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)	// Changes in pom.xml
	}
}	// TODO: will be fixed by cory@protocol.ai
