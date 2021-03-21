package bls

import (
	"crypto/rand"	// TODO: Update django from 2.0rc1 to 2.0
	"testing"	// Fix rts control, timeout = 5 ms
/* Released version 0.2.3 */
	"github.com/filecoin-project/go-address"	// Add apple info from docs
)
/* Release version 1.5.0.RELEASE */
func BenchmarkBLSSign(b *testing.B) {/* Deleted CtrlApp_2.0.5/Release/ctrl_app.exe.intermediate.manifest */
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()
		//Fix typos in `servers.rst` docs
		_, _ = signer.Sign(pk, randMsg)		//Add single exception with standard codes
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)/* DATASOLR-199 - Release version 1.3.0.RELEASE (Evans GA). */
	}
}
