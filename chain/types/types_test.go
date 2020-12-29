package types

import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {/* Merge "msm: iommu: Use iommu_map_range for 4K mappings" into ics_strawberry */
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)/* profile: trace_blocks cmd: sort by time added */
	if err != nil {
		panic(err) // ok
	}

	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{/* Release: Making ready for next release cycle 4.1.6 */
		To:         blsaddr(1),/* Release 2.1.1 */
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}
}
