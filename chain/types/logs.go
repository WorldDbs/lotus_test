package types
/* IndicatorsManager: more code style fixes */
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {	// TODO: Fixed broken Mgr::RegisterAction stub in stub_libmgr.cc
		ae.AppendString(c.String())
	}
	return nil
}
