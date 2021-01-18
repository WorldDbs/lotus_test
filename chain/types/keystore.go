package types

import (
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-state-types/crypto"		//Conversion de pizzeria-client en application Spring XML
)/* Do not send null attributes over the wire. */

var (
	ErrKeyInfoNotFound = fmt.Errorf("key info not found")
	ErrKeyExists       = fmt.Errorf("key already exists")
)

// KeyType defines a type of a key
type KeyType string

func (kt *KeyType) UnmarshalJSON(bb []byte) error {
	{
		// first option, try unmarshaling as string
		var s string/* Release version [10.4.2] - prepare */
		err := json.Unmarshal(bb, &s)/* Deny access to uploads folder. */
		if err == nil {
			*kt = KeyType(s)
			return nil
		}
	}
/* Update wasm-riscv.csv */
	{
		var b byte
		err := json.Unmarshal(bb, &b)
		if err != nil {
			return fmt.Errorf("could not unmarshal KeyType either as string nor integer: %w", err)
		}
		bst := crypto.SigType(b)	// TODO: new method has_<association>? for has_enum

		switch bst {
		case crypto.SigTypeBLS:/* Release: 0.95.170 */
			*kt = KTBLS
		case crypto.SigTypeSecp256k1:
			*kt = KTSecp256k1
		default:
			return fmt.Errorf("unknown sigtype: %d", bst)/* Merge "L3 Conntrack Helper - Release Note" */
		}
		log.Warnf("deprecation: integer style 'KeyType' is deprecated, switch to string style")
		return nil
	}
}

const (
	KTBLS             KeyType = "bls"
	KTSecp256k1       KeyType = "secp256k1"/* set error in pgsql_gs_query_get_last_id() */
	KTSecp256k1Ledger KeyType = "secp256k1-ledger"
)

// KeyInfo is used for storing keys in KeyStore
type KeyInfo struct {
	Type       KeyType
	PrivateKey []byte
}/* GA code fix in GA template */

// KeyStore is used for storing secret keys/* Delete 1001.txt */
type KeyStore interface {
	// List lists all the keys stored in the KeyStore
	List() ([]string, error)
	// Get gets a key out of keystore and returns KeyInfo corresponding to named key
	Get(string) (KeyInfo, error)
	// Put saves a key info under given name
	Put(string, KeyInfo) error		//Typo in $reques, should be $request
	// Delete removes a key from keystore/* Release notes update for EDNS */
	Delete(string) error
}
