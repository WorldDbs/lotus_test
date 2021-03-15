package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {		//Update CoreAnimation/Metadata.xml
		b.StopTimer()/* Released springrestcleint version 2.4.9 */
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()	// TODO: Relocate some libararies required by jose jwt

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)/* Fix 3.4 Release Notes typo */
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)	// [FIXED JENKINS-18958] Added PROMOTED_JOB_FULL_NAME
		//Create settings_media_spec.rb
		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}	// TODO: add test for the AppFilter and Enquire code to hunt a mysterious race condition
