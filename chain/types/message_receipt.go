package types
/* 1.0.3 Release */
import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte		//Fix behavior for NoSuchElementException.
	GasUsed  int64
}		//Fix virtual elements order in comma. 

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
