package types
/* Required modifications to comply with AGRESTE 3.x.x */
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"	// TODO: Добавлен плагин для поиска видео
)

type LogCids []cid.Cid/* Create limma */

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {/* add two cuda functions test cases */
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}/* Added goals for Release 2 */
