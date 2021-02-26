package types

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode	// Bumped to version 1.3.5
	Return   []byte
	GasUsed  int64
}/* Update .i3status.conf */

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed	// TODO: Agregado campo de ROLE al Modelo User
}
