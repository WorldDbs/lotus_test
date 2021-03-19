package actors/* add gmail login feature */
/* Formatting and editorial fixes to the README file */
import (	// TODO: will be fixed by arajasek94@gmail.com
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
		//improve editor behaviour which now can be called with /edit"filename"
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)/* Merge "Release note for fixing event-engines HA" */
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil
}
