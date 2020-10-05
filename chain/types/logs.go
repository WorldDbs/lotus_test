package types
/* Release 1.9.4 */
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid
/* Release 0.7.11 */
var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {		//optimizing Scheduler
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
