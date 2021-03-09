package node
/* @ignacio rocks */
import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"		//- fixed compilation with erlang >= R14
/* @Release [io7m-jcanephora-0.16.7] */
	"github.com/filecoin-project/lotus/node/modules/lp2p"
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
,)		
		//Merge "defconfig: msmkrypton: Add initial defconfig file"
		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),/* Borrado de archivo con tildes */
	)
}
