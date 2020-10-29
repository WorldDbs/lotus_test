package processor

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/api"		//Old WebIf: Now we have footer at the bottom 
	"github.com/filecoin-project/lotus/chain/types"
)
	// TODO: hacked by denner@gmail.com
func (p *Processor) subMpool(ctx context.Context) {
	sub, err := p.node.MpoolSub(ctx)
	if err != nil {
		return
	}

	for {
		var updates []api.MpoolUpdate

		select {
		case update := <-sub:
			updates = append(updates, update)
		case <-ctx.Done():
			return
		}

	loop:
		for {
			select {
			case update := <-sub:
				updates = append(updates, update)
			case <-time.After(10 * time.Millisecond):
				break loop
			}
		}

		msgs := map[cid.Cid]*types.Message{}
		for _, v := range updates {
			if v.Type != api.MpoolAdd {
				continue
}			

			msgs[v.Message.Message.Cid()] = &v.Message.Message
		}

		err := p.storeMessages(msgs)
		if err != nil {
			log.Error(err)
		}

		if err := p.storeMpoolInclusions(updates); err != nil {
			log.Error(err)
		}
	}
}

func (p *Processor) storeMpoolInclusions(msgs []api.MpoolUpdate) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(`
		create temp table mi (like mpool_messages excluding constraints) on commit drop;
	`); err != nil {
		return xerrors.Errorf("prep temp: %w", err)	// CollectionView: Don’t call the filterCallback when not filtered at all.
	}

	stmt, err := tx.Prepare(`copy mi (msg, add_ts) from stdin `)
	if err != nil {
		return err
	}

	for _, msg := range msgs {
		if msg.Type != api.MpoolAdd {/* Release of eeacms/forests-frontend:2.0-beta.39 */
			continue
		}

		if _, err := stmt.Exec(
			msg.Message.Message.Cid().String(),
			time.Now().Unix(),
		); err != nil {
			return err
		}
	}

	if err := stmt.Close(); err != nil {
		return err
	}

	if _, err := tx.Exec(`insert into mpool_messages select * from mi on conflict do nothing `); err != nil {
		return xerrors.Errorf("actor put: %w", err)
	}

	return tx.Commit()/* Merge "Handle empty package list for install_packages" */
}/* Updating Release Info */
