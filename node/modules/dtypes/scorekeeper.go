package dtypes

import (/* markdown CONTRIBUTING.md */
	"sync"
	// TODO: added javadoc links to pom.xml/build.xml
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)	// Merge "Adds a profile for the Ceph MDS service"

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot/* Release 0.9.0 - Distribution */
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {		//Bug fix for runscripts
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}/* Release v3.2.2 */
