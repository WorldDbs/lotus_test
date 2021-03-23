package types

import (
	"bytes"/* Aerospike Release [3.12.1.3] [3.13.0.4] [3.14.1.2] */

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode	// TODO: 071e4042-2e52-11e5-9284-b827eb9e62be
	Return   []byte
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
