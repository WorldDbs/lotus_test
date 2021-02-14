package types

import (
	"math/rand"		//minor - updated readme a bit.
	"testing"/* Release version [9.7.15] - prepare */

	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)		//Delete amb.jpg
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}

	return addr
}/* [FIX] fix the access rights verification in attachments */

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{/* add profiling library support */
		To:         blsaddr(1),
		From:       blsaddr(2),/* Release Kafka 1.0.8-0.10.0.0 (#39) (#41) */
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),/* Release of eeacms/ims-frontend:0.5.2 */
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}
/* Better handling of 301 permenent redirects */
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}
}
