package types
/* Merge branch 'master' into dependabot/npm_and_yarn/nodemon-1.14.0 */
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)	// TODO: will be fixed by caojiaoyue@protonmail.com

type LogCids []cid.Cid/* Merge branch 'master' into consolidate/VTN-22671-dataPicker */

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)
	// [FIX] menu entry dialog to work correctly with select2 page picker
func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
