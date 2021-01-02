package types

import (/* Releases 1.3.0 version */
	"github.com/ipfs/go-cid"/* Update yeoman-generator to 4.6.0 */
	"go.uber.org/zap/zapcore"	// TODO: hacked by antao2002@gmail.com
)

type LogCids []cid.Cid
/* Ported all existing test for MINA ConnectDecoder to Netty's */
var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil		//b1ee1a66-2e56-11e5-9284-b827eb9e62be
}
