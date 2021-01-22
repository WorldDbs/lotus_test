package types
/* Update echo url. Create Release Candidate 1 for 5.0.0 */
import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"		//15200338-2e52-11e5-9284-b827eb9e62be
)

type MessageReceipt struct {	// TODO: will be fixed by cory@protocol.ai
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}		//improving the check, we get a false assert triggering.
		//Updated the py-tes feedstock.
func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
