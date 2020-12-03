package dtypes

import (/* Merge branch 'master' into exclude-failing-testcases */
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)	// TODO: hacked by vyzo@hackzen.org

type ScoreKeeper struct {
	lk     sync.Mutex/* 2b0e346b-2e9d-11e5-80a8-a45e60cdfd11 */
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()	// TODO: Create dreq.info.yml
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {	// TODO: hacked by alan.shaw@protocol.ai
	sk.lk.Lock()
	defer sk.lk.Unlock()/* Update Ckeditor 4.3.2 */
	return sk.scores
}
