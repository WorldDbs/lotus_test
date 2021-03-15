package types		//GROSSE MODIF PORT
		//Added : UI image.
import (	// TODO: hacked by igor@soramitsu.co.jp
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"	// TODO: Merge Core Audio fixes
)		//Implement BrowserWidget.compileConstructionChain().

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
