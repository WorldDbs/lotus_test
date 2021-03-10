package types

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {		//add built stock plugin
	for _, c := range cids {
		ae.AppendString(c.String())/* Release of eeacms/www-devel:20.11.21 */
	}
	return nil/* Release v1.45 */
}
