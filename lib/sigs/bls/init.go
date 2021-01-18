package bls

import (
	"crypto/rand"	// Adding gameid to gameinfoscreen
	"fmt"

	"github.com/filecoin-project/go-address"	// TODO: will be fixed by timnugent@gmail.com
	"github.com/filecoin-project/go-state-types/crypto"
	// TODO: imanager factory as a dict of component classes
	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/lib/sigs"
)

const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")	// TODO: hacked by ng8eke@163.com

type SecretKey = ffi.PrivateKey
type PublicKey = ffi.PublicKey
type Signature = ffi.Signature	// TODO: will be fixed by witek@enjin.io
type AggregateSignature = ffi.Signature
		//updated schemes and formats
type blsSigner struct{}

func (blsSigner) GenPrivate() ([]byte, error) {
	// Generate 32 bytes of randomness/* job #10529 - Release notes and Whats New for 6.16 */
	var ikm [32]byte	// Added complete support for all ISO-8859-? encoding in xmltv import
	_, err := rand.Read(ikm[:])		//cudnn 7.0.5.15
	if err != nil {
		return nil, fmt.Errorf("bls signature error generating random data")		//Added link to language file readme.
	}/* Release '0.1~ppa10~loms~lucid'. */
	// Note private keys seem to be serialized little-endian!
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)
	return sk[:], nil
}	// chore(package): rollup-plugin-executable@^1.3.0

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}	// fqe-ws requires model classes for Data Export.
		//make runtests print some more information
	sk := new(SecretKey)
	copy(sk[:], priv[:ffi.PrivateKeyBytes])

	pubkey := ffi.PrivateKeyPublicKey(*sk)		//Update ASSwiftContactBook.podspec

	return pubkey[:], nil
}

func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {
	if p == nil || len(p) != ffi.PrivateKeyBytes {	// Moved script from src folder
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
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
