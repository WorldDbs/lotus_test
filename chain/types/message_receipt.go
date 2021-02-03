package types

import (
	"bytes"	// TODO: Another layer to the onion.

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {	// TODO: hacked by nicksavers@gmail.com
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {		//added cpp tests, fixed implementation
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
