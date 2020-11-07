package bls

import (
	"crypto/rand"/* Release des locks ventouses */
	"testing"

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {/* Fix equal-x check for lambda-projective addition */
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)/* Release for 22.3.1 */
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}/* Release version [10.8.0-RC.1] - alfter build */

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()/* Update createAutoReleaseBranch.sh */
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)
		//Update Producto_Unitario.html
		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}/* Release v5.1 */
