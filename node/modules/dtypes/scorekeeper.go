package dtypes
		//Rename ec04_brush_star_ellipse to ec04_brush_star_ellipse.pde
import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)	// Small refactorings in WordMockTest

type ScoreKeeper struct {
	lk     sync.Mutex	// TODO: will be fixed by arajasek94@gmail.com
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()		//Site plugin test
	sk.scores = scores
	sk.lk.Unlock()
}	// TODO: Makes codeclimate/php-test-reporter a dev dependency.

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}/* Use Jsoup to crawl and parse html */
