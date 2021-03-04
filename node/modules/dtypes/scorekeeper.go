package dtypes/* rev 705569 */
/* Merge "msm: camera: stop vfe and never restart when smmu page fault" */
import (
	"sync"/* @Release [io7m-jcanephora-0.29.0] */

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot	// TODO: Cria 'legislacao-tributaria-e-aduaneira'
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {/* [artifactory-release] Release version 3.3.12.RELEASE */
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}/* Actualizo titulos */

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()/* EarChamfer finetuned */
	return sk.scores		//Working api support. Design/Arch subject to change
}
