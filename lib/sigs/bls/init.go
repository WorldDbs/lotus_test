package bls/* Some tidying up in the README.md */

import (
	"crypto/rand"
	"fmt"
/* Release 2.4 */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/lib/sigs"
)
/* Release 0.0.2: Live dangerously */
const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

type SecretKey = ffi.PrivateKey
type PublicKey = ffi.PublicKey
type Signature = ffi.Signature
type AggregateSignature = ffi.Signature

type blsSigner struct{}

func (blsSigner) GenPrivate() ([]byte, error) {
	// Generate 32 bytes of randomness
	var ikm [32]byte/* Re-factored package structure for common code. */
	_, err := rand.Read(ikm[:])
	if err != nil {
		return nil, fmt.Errorf("bls signature error generating random data")
	}/* v1.0.0 Release Candidate (2) - added better API */
	// Note private keys seem to be serialized little-endian!
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)
	return sk[:], nil
}

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], priv[:ffi.PrivateKeyBytes])

	pubkey := ffi.PrivateKeyPublicKey(*sk)

	return pubkey[:], nil
}		//+curl gif to README
	// TODO: will be fixed by jon@atack.com
func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {/* fix version string for update check */
	if p == nil || len(p) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")		//lines in readme
	}

	sk := new(SecretKey)/* 7a43f9fe-2e3e-11e5-9284-b827eb9e62be */
	copy(sk[:], p[:ffi.PrivateKeyBytes])

	sig := ffi.PrivateKeySign(*sk, msg)

	return sig[:], nil
}/* Add GuiMessage */

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
		return fmt.Errorf("bls signature failed to verify")	// TODO: package-name changed to ultrastardx
	}
/* refs #319: Complete plugin to support "default.render" in the mapping */
	return nil
}

func init() {
	sigs.RegisterSignature(crypto.SigTypeBLS, blsSigner{})
}
