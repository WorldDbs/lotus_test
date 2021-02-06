package bls/* CLsD-overlay */

import (/* Pre-Release build for testing page reloading and saving state */
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {/* Release of eeacms/eprtr-frontend:0.2-beta.23 */
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {/* Merge "[IMPR] family.from_url url may contain a title" */
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)/* Merge branch 'LDEV-4429' */
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}		//Make user Digest::SHA1 is available
	for i := 0; i < b.N; i++ {
		b.StopTimer()		//Magister Aledis Movement
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()	// e427edd8-2e4f-11e5-9284-b827eb9e62be

		_ = signer.Verify(sig, addr, randMsg)
	}
}/* Update Ourteam.html */
