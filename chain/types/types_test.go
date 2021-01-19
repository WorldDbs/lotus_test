package types

import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)		//more info on graph

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)/* Denote Spark 2.8.1 Release */

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok/* Create nicepanel.jquery.js */
	}

	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),		//https://pt.stackoverflow.com/q/385877/101
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,		//changed write to create
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {/* Create student17c.xml */
			b.Fatal(err)
		}
	}	// TODO: Update future from 0.18.0 to 0.18.2
}
