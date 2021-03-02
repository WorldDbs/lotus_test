package types

import (/* * Mark as Release Candidate 1. */
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)	// TODO: - putting commonly used visualizers into annis-utilsgui
/* Changed some spacing */
type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())/* Delete Compiled-Releases.md */
	}
	return nil
}/* Release 7.1.0 */
