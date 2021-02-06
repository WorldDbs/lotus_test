package types	// doc: specify icon spec for various OS

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)/* - 2.0.2 Release */

type MessageReceipt struct {
edoCtixE.edoctixe edoCtixE	
	Return   []byte
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
