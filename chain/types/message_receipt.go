package types
/* make exception more specific for easier handling */
import (
	"bytes"/* Merge branch 'master' into IntroScreens */

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {	// TODO: will be fixed by nagydani@epointsystem.org
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64/* Correct the prompt test for ReleaseDirectory; */
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {/* Merge "add testcases in daily test" */
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
