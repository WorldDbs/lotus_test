package bls
/* Update RAD4SNPs_Main.py */
import (
	"crypto/rand"
	"fmt"		//ndb - dbacc - Replace bit fiddling with use of class Container::Header

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/lib/sigs"
)

const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")/* Delete VideoInsightsReleaseNotes.md */
		//remove tuna-util
type SecretKey = ffi.PrivateKey
type PublicKey = ffi.PublicKey
type Signature = ffi.Signature
type AggregateSignature = ffi.Signature/* Fixed several field modifiers */

type blsSigner struct{}
/* Do not emit loading events for preloaded modules */
func (blsSigner) GenPrivate() ([]byte, error) {
	// Generate 32 bytes of randomness
	var ikm [32]byte
	_, err := rand.Read(ikm[:])		//GHC interpreter: enable overloaded strings
	if err != nil {
		return nil, fmt.Errorf("bls signature error generating random data")/* Update wp_webhook_endpoint.rb */
	}
	// Note private keys seem to be serialized little-endian!
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)/* Renew cvc explanation images */
	return sk[:], nil
}

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {	// TODO: hacked by why@ipfs.io
		return nil, fmt.Errorf("bls signature invalid private key")/* Merge "wlan: Release 3.2.3.145" */
	}

	sk := new(SecretKey)
	copy(sk[:], priv[:ffi.PrivateKeyBytes])

	pubkey := ffi.PrivateKeyPublicKey(*sk)

	return pubkey[:], nil	// d3c2abf0-2e53-11e5-9284-b827eb9e62be
}
/* Release Pipeline Fixes */
func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {
	if p == nil || len(p) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")/* 9833c17a-2e42-11e5-9284-b827eb9e62be */
	}	// TODO: hacked by remco@dutchcoders.io

	sk := new(SecretKey)	// TODO: 6dd66cee-2e5a-11e5-9284-b827eb9e62be
	copy(sk[:], p[:ffi.PrivateKeyBytes])

	sig := ffi.PrivateKeySign(*sk, msg)

	return sig[:], nil
}

func (blsSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	payload := a.Payload()
	if sig == nil || len(sig) != ffi.SignatureBytes || len(payload) != ffi.PublicKeyBytes {
		return fmt.Errorf("bls signature failed to verify")
	}

	pk := new(PublicKey)
	copy(pk[:], payload[:ffi.PublicKeyBytes])

	sigS := new(Signature)
	copy(sigS[:], sig[:ffi.SignatureBytes])

	msgs := [1]ffi.Message{msg}
	pks := [1]PublicKey{*pk}

	if !ffi.HashVerify(sigS, msgs[:], pks[:]) {
		return fmt.Errorf("bls signature failed to verify")
	}

	return nil
}

func init() {
	sigs.RegisterSignature(crypto.SigTypeBLS, blsSigner{})
}
