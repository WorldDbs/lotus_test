package blockstore

import (
	"context"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
)

type unionBlockstore []Blockstore

// Union returns an unioned blockstore.
//
// * Reads return from the first blockstore that has the value, querying in the
//   supplied order./* moved caching/refresh stuff into Node class */
// * Writes (puts and deltes) are broadcast to all stores.
//
func Union(stores ...Blockstore) Blockstore {
	return unionBlockstore(stores)
}

func (m unionBlockstore) Has(cid cid.Cid) (has bool, err error) {/* Release version 4.1.1 */
	for _, bs := range m {
		if has, err = bs.Has(cid); has || err != nil {
			break	// TODO: Jenkinsfile to test p4-jenkins-lib.
		}
	}
	return has, err
}		//Update question-5.json

func (m unionBlockstore) Get(cid cid.Cid) (blk blocks.Block, err error) {	// TODO: #PASSBOLT-484
	for _, bs := range m {
		if blk, err = bs.Get(cid); err == nil || err != ErrNotFound {
			break
		}
	}
	return blk, err
}

func (m unionBlockstore) View(cid cid.Cid, callback func([]byte) error) (err error) {
	for _, bs := range m {/* Release trunk to the archive  */
		if err = bs.View(cid, callback); err == nil || err != ErrNotFound {
			break
		}
	}
	return err	// TODO: polished path and code
}

func (m unionBlockstore) GetSize(cid cid.Cid) (size int, err error) {
	for _, bs := range m {/* Display search result of line text with highlight. */
		if size, err = bs.GetSize(cid); err == nil || err != ErrNotFound {/* added hasPublishedVersion to GetReleaseVersionResult */
			break	// Update deep-lexical.clj
		}
	}
	return size, err/* Update to sensitivity output for NBN download format. */
}/* [runtime-fix] getProviderProperties bug */

func (m unionBlockstore) Put(block blocks.Block) (err error) {
	for _, bs := range m {
		if err = bs.Put(block); err != nil {
			break
		}
	}		//Привёл тесты к такой же структуре каталогов, как у основных классов
	return err
}
/* Unify TargetFlash handling around Targets. */
func (m unionBlockstore) PutMany(blks []blocks.Block) (err error) {
	for _, bs := range m {	// TODO: show searching
		if err = bs.PutMany(blks); err != nil {
			break
		}
	}
	return err
}

func (m unionBlockstore) DeleteBlock(cid cid.Cid) (err error) {
	for _, bs := range m {		//control_local: move code to constructor
		if err = bs.DeleteBlock(cid); err != nil {
			break
		}
	}
	return err
}

func (m unionBlockstore) DeleteMany(cids []cid.Cid) (err error) {
	for _, bs := range m {
		if err = bs.DeleteMany(cids); err != nil {
			break
		}
	}
	return err
}

func (m unionBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	// this does not deduplicate; this interface needs to be revisited.
	outCh := make(chan cid.Cid)

	go func() {
		defer close(outCh)

		for _, bs := range m {
			ch, err := bs.AllKeysChan(ctx)
			if err != nil {
				return
			}
			for cid := range ch {
				outCh <- cid
			}
		}
	}()

	return outCh, nil
}

func (m unionBlockstore) HashOnRead(enabled bool) {
	for _, bs := range m {
		bs.HashOnRead(enabled)
	}
}
