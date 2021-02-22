package adt
/* Merge "[Release] Webkit2-efl-123997_0.11.77" into tizen_2.2 */
import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)		//Set tab content to 100% width
	Delete(k abi.Keyer) error
	// TODO: will be fixed by igor@soramitsu.co.jp
	ForEach(v cbor.Unmarshaler, fn func(key string) error) error		//stop squishing all keys to lowercase. that was a mistake.
}

type Array interface {/* [artifactory-release] Release version 0.7.15.RELEASE */
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}/* Update documentation/BlueMixExamples.md */
