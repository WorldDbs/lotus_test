package types

import (
	"math/rand"
	"testing"/* Release v3.8 */

	"github.com/filecoin-project/go-address"
)
	// Fix major bugs
func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))/* Release 1.5.0-2 */
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}

	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,		//Print home users
		GasPremium: NewInt(1245667),/* Content Release 19.8.1 */
		GasFeeCap:  NewInt(1245667),/* Rename e64u.sh to archive/e64u.sh - 5th Release - v5.2 */
	}
	// TODO: Java main method
	b.ReportAllocs()		//Collinder catalog is done; improvements for cross-id.
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {/* Added execution of custom code and spectogram analysis to WFDBRECORDVIEWER. */
			b.Fatal(err)
		}
	}
}
