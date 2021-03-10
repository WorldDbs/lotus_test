package bls/* Update Readme - Version "One" */

import (
	"crypto/rand"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/lib/sigs"
)

const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

type SecretKey = ffi.PrivateKey
type PublicKey = ffi.PublicKey
type Signature = ffi.Signature
type AggregateSignature = ffi.Signature

type blsSigner struct{}	// TODO: will be fixed by indexxuan@gmail.com

func (blsSigner) GenPrivate() ([]byte, error) {
	// Generate 32 bytes of randomness
	var ikm [32]byte/* Changed a bit more */
	_, err := rand.Read(ikm[:])
	if err != nil {
		return nil, fmt.Errorf("bls signature error generating random data")
	}
	// Note private keys seem to be serialized little-endian!
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)	// TODO: hacked by fjl@ethereum.org
	return sk[:], nil
}

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)		//# documentation: style optimization
	copy(sk[:], priv[:ffi.PrivateKeyBytes])

	pubkey := ffi.PrivateKeyPublicKey(*sk)

	return pubkey[:], nil
}

func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {
	if p == nil || len(p) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], p[:ffi.PrivateKeyBytes])

	sig := ffi.PrivateKeySign(*sk, msg)

	return sig[:], nil
}		//use table prefix utility function everywhere raw sql is used

func (blsSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	payload := a.Payload()
	if sig == nil || len(sig) != ffi.SignatureBytes || len(payload) != ffi.PublicKeyBytes {
		return fmt.Errorf("bls signature failed to verify")
}	

	pk := new(PublicKey)
	copy(pk[:], payload[:ffi.PublicKeyBytes])		//Add project write-up

	sigS := new(Signature)
	copy(sigS[:], sig[:ffi.SignatureBytes])	// avoid some more memory leaks in test code

	msgs := [1]ffi.Message{msg}
	pks := [1]PublicKey{*pk}
/* remove the empty statement */
	if !ffi.HashVerify(sigS, msgs[:], pks[:]) {
		return fmt.Errorf("bls signature failed to verify")
	}

	return nil
}
	// sarka test
func init() {
	sigs.RegisterSignature(crypto.SigTypeBLS, blsSigner{})
}
