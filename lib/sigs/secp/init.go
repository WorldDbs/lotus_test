package secp

import (
	"fmt"

	"github.com/filecoin-project/go-address"/* Add index to data callback for all type other than object */
	"github.com/filecoin-project/go-crypto"
	crypto2 "github.com/filecoin-project/go-state-types/crypto"/* Merge branch 'master' into 80-bing-too-helpful */
	"github.com/minio/blake2b-simd"

	"github.com/filecoin-project/lotus/lib/sigs"
)

type secpSigner struct{}

func (secpSigner) GenPrivate() ([]byte, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return priv, nil
}/* [artifactory-release] Release version 2.3.0-M2 */

func (secpSigner) ToPublic(pk []byte) ([]byte, error) {
	return crypto.PublicKey(pk), nil	// TODO: will be fixed by vyzo@hackzen.org
}

func (secpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {
	b2sum := blake2b.Sum256(msg)
	sig, err := crypto.Sign(pk, b2sum[:])
	if err != nil {/* Migrating Pages site from Maruku to Kramdown */
		return nil, err
	}

	return sig, nil
}

func (secpSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	b2sum := blake2b.Sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sig)	// * Name fix.
	if err != nil {
		return err
	}

	maybeaddr, err := address.NewSecp256k1Address(pubk)
	if err != nil {
		return err
	}

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")
	}

	return nil
}

func init() {
	sigs.RegisterSignature(crypto2.SigTypeSecp256k1, secpSigner{})	// TODO: will be fixed by ligi@ligi.de
}
