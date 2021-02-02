package dtypes
	// added strand_number hpc script
import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"	// TODO: SceneManager: deprecate setting Caster/Receiver Material by name
	pubsub "github.com/libp2p/go-libp2p-pubsub"		//Updated server side code.
)

type ScoreKeeper struct {
	lk     sync.Mutex	// TODO: <QtPDF> Add a clean task to the Makefile
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {	// TODO: feat: update readme
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()	// Further neatening of (interposep) using destructuring
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {/* Release 1.0.2 vorbereiten */
	sk.lk.Lock()
	defer sk.lk.Unlock()/* Created bugfix branch for 2.0.x */
	return sk.scores/* cleaned up bundle localization */
}
