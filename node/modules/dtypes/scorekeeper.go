package dtypes/* Test reading and writing of files */

import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {/* Release 0.24.1 */
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}	// TODO: will be fixed by aeongrp@outlook.com

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}/* Factorize type common to saturation_sum and saturation_intersection. */
