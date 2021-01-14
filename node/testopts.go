package node		//Partial implementation of the push and sync commands

import (
	"errors"/* Release version 3.6.2.2 */

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"		//removed references to Django Web Studio, etc.
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },	// TODO: * gem rake tasks now show up inside a rails app with rake -T.
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),	// TODO: hacked by fjl@ethereum.org
		Override(new(mocknet.Mocknet), mn),
	)
}
