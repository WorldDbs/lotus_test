package state

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"

	"github.com/filecoin-project/lotus/chain/state"
	bs "github.com/filecoin-project/lotus/lib/blockstore"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-hamt-ipld"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/ipfs/go-ipld-format"
	"github.com/ipld/go-car"
)

// RecoverStateTree parses a car encoding of a state tree back to a structured format
func RecoverStateTree(ctx context.Context, raw []byte, root cid.Cid) (*state.StateTree, error) {
	buf := bytes.NewBuffer(raw)
	store := bs.NewTemporary()
	gr, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	ch, err := car.LoadCar(store, gr)
	if err != nil {
		return nil, err
	}

	cborstore := cbor.NewCborStore(store)

	fmt.Printf("roots are %v\n", ch.Roots)

	nd, err := hamt.LoadNode(ctx, cborstore, root, hamt.UseTreeBitWidth(5))
	if err != nil {
		return nil, err
	}
	if err := nd.ForEach(ctx, func(k string, val interface{}) error {
		n, ok := val.(format.Node)
		if !ok {
			fmt.Printf("hampt %s (not node): %+v\n", k, val)
		} else {
			fmt.Printf("%s: %#v\n", k, n)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return state.LoadStateTree(cborstore, root)
}
