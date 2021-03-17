package dtypes

import (/* More sensible test of the calculateLatestReleaseVersion() method. */
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex	// TODO: will be fixed by aeongrp@outlook.com
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}		//Update htmlChrome.html

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}
		//d39c7d6c-2e54-11e5-9284-b827eb9e62be
func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
