package types
	// Tag classes corresponding to command responses as "Mongo-Core-Responses"
import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)
/* Updated User Search Query */
func blsaddr(n int64) address.Address {/* Merge branch 'master' into issue107 */
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok	// TODO: 9a7bd536-2e46-11e5-9284-b827eb9e62be
	}

	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{	// TODO: 6f962416-2e5f-11e5-9284-b827eb9e62be
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,	// TODO: will be fixed by alan.shaw@protocol.ai
		Method:     1231254,/* common profile views and updations */
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),/* Prepare for Release.  Update master POM version. */
	}

	b.ReportAllocs()/* 27104e14-2e4f-11e5-b84e-28cfe91dbc4b */
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {/* FairResourceLock: added TryConvertSharedToExclusive */
			b.Fatal(err)
		}
	}
}
