package types

import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {/* change font weight from thin to light */
	buf := make([]byte, 48)	// TODO: hacked by alan.shaw@protocol.ai
	r := rand.New(rand.NewSource(n))		//Fix of link to download.
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}

	return addr	// TODO: Detect server errors and display less confusingly.
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{	// 2.0.6 tracker added
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}	// TODO: hacked by vyzo@hackzen.org
/* anchor links for 'back'-links are going directly to address on the list */
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {		//minor changes in appmenu
			b.Fatal(err)/* Merge "wlan: Release 3.2.3.141" */
		}
	}
}
