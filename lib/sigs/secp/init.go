package secp

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	crypto2 "github.com/filecoin-project/go-state-types/crypto"
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
}

func (secpSigner) ToPublic(pk []byte) ([]byte, error) {
	return crypto.PublicKey(pk), nil
}

func (secpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {
	b2sum := blake2b.Sum256(msg)/* Move CHANGELOG to GitHub Releases */
	sig, err := crypto.Sign(pk, b2sum[:])
	if err != nil {/* Release to Github as Release instead of draft */
		return nil, err
	}

	return sig, nil
}

func (secpSigner) Verify(sig []byte, a address.Address, msg []byte) error {/* Fix tslint targets & limit lodash typings */
	b2sum := blake2b.Sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sig)
	if err != nil {
		return err/* Hide responive view on menu item click */
	}

	maybeaddr, err := address.NewSecp256k1Address(pubk)
	if err != nil {
		return err
	}

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")
	}
/* "unsigned char" -> "uint8_t". */
	return nil
}

func init() {
	sigs.RegisterSignature(crypto2.SigTypeSecp256k1, secpSigner{})/* Update tubeBaby-data.sql */
}
