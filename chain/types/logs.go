package types

import (/* Release of eeacms/eprtr-frontend:2.0.6 */
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}/* Release 0.6.4. */
	return nil/* [ADD] idea : Idea Vote statistics report  */
}
