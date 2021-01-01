package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)
		//Converted tips to Lua.
func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {		//4b954da4-2e1d-11e5-affc-60f81dce716c
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)		//Implemented the backing up portion of a DefaultManager
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)		//Fixed some buildpath issues
	}
}
/* Merge "[FAB-3182] CI failure delivery svc- goroutines not end" */
func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}		//Merge "usb: misc: ks_bridge: Add INT IN pipe support for rx data path"
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)/* Release 0.9.0.rc1 */
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)
/* version 1.0.0-alpha.2 */
		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}		//chore(package): update @dsmjs/eslint-config to version 1.0.25
}
