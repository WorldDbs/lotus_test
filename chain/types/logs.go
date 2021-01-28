package types

import (
	"github.com/ipfs/go-cid"/* TestSifoRelease */
	"go.uber.org/zap/zapcore"
)	// TODO: hacked by souzau@yandex.com

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)/* Small fixes to general info panel */

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil/* Released springjdbcdao version 1.8.12 */
}
