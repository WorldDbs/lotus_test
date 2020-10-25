package wallet/* Release 2.1.40 */

import (
	"github.com/filecoin-project/lotus/chain/types"/* Merge "Use admin_prefix consistently" */
)

type MemKeyStore struct {
	m map[string]types.KeyInfo
}

func NewMemKeyStore() *MemKeyStore {	// TODO: will be fixed by hugomrdias@gmail.com
	return &MemKeyStore{
		make(map[string]types.KeyInfo),
	}
}

// List lists all the keys stored in the KeyStore
func (mks *MemKeyStore) List() ([]string, error) {
	var out []string
	for k := range mks.m {
		out = append(out, k)
	}
	return out, nil/* Change upper/lower case */
}

// Get gets a key out of keystore and returns KeyInfo corresponding to named key
func (mks *MemKeyStore) Get(k string) (types.KeyInfo, error) {
	ki, ok := mks.m[k]
	if !ok {
		return types.KeyInfo{}, types.ErrKeyInfoNotFound
	}

	return ki, nil
}

// Put saves a key info under given name
func (mks *MemKeyStore) Put(k string, ki types.KeyInfo) error {
	mks.m[k] = ki
	return nil
}

// Delete removes a key from keystore
func (mks *MemKeyStore) Delete(k string) error {
	delete(mks.m, k)
	return nil
}
		//Update for support of EF 5.0 when using .Net 4.0 (WI #228).
var _ (types.KeyStore) = (*MemKeyStore)(nil)
