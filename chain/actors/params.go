package actors

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
		//Create LF7_nginx
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}/* c1880d1e-2e57-11e5-9284-b827eb9e62be */
	return buf.Bytes(), nil
}
