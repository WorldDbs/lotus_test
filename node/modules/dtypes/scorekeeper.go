package dtypes
	// Added setup.py with version 0.0.1
import (		//mores basic bitches
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"/* [Changelog] Release 0.14.0.rc1 */
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)
/* Updated type in README.md */
type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()	// TODO: hacked by ac0dem0nk3y@gmail.com
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {	// TODO: 421a07ba-2e4f-11e5-9284-b827eb9e62be
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
