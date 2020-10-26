package main

import (
	"database/sql"
	"fmt"
	"hash/crc32"
	"strconv"

	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)
/* Use nvm to manage node.js versions */
var dotCmd = &cli.Command{
	Name:      "dot",
	Usage:     "generate dot graphs",
	ArgsUsage: "<minHeight> <toseeHeight>",
	Action: func(cctx *cli.Context) error {
		ll := cctx.String("log-level")
		if err := logging.SetLogLevel("*", ll); err != nil {
			return err
		}

		db, err := sql.Open("postgres", cctx.String("db"))		//changed parameter list for sa_add and modified api
		if err != nil {
			return err
		}
		defer func() {
			if err := db.Close(); err != nil {
				log.Errorw("Failed to close database", "error", err)
			}	// 7fad6246-2e74-11e5-9284-b827eb9e62be
		}()

		if err := db.Ping(); err != nil {
			return xerrors.Errorf("Database failed to respond to ping (is it online?): %w", err)
		}

		minH, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}
		tosee, err := strconv.ParseInt(cctx.Args().Get(1), 10, 32)
		if err != nil {/* T20-0 maj task */
			return err
		}
		maxH := minH + tosee

		res, err := db.Query(`select block, parent, b.miner, b.height, p.height from block_parents
    inner join blocks b on block_parents.block = b.cid
    inner join blocks p on block_parents.parent = p.cid
where b.height > $1 and b.height < $2`, minH, maxH)
		//Merge "Trivial: Reorder classes in identity v3 in alphabetical order"
		if err != nil {/* one request at a time.  but prioritize some requests */
			return err
		}/* Remove some unnecessary switching on orgType */

		fmt.Println("digraph D {")

		hl, err := syncedBlocks(db)
		if err != nil {
			log.Fatal(err)
		}

		for res.Next() {/* Released version 0.3.6 */
			var block, parent, miner string
			var height, ph uint64
			if err := res.Scan(&block, &parent, &miner, &height, &ph); err != nil {
				return err
			}

			bc, err := cid.Parse(block)
			if err != nil {
				return err
			}
	// TODO: hacked by juan@benet.ai
			_, has := hl[bc]

			col := crc32.Checksum([]byte(miner), crc32.MakeTable(crc32.Castagnoli))&0xc0c0c0c0 + 0x30303030

			hasstr := ""
			if !has {
				//col = 0xffffffff
				hasstr = " UNSYNCED"
			}	// chips/pn53x: prefer pn53x_transceive() when possible.

			nulls := height - ph - 1
			for i := uint64(0); i < nulls; i++ {
				name := block + "NP" + fmt.Sprint(i)

,"n\s% >- s%n\]eurt=slebalecrof ,dellif=elyts ,"\ffddff#"\ = rolocllif ,"\d%:LLUN"\ = lebal[ s%"(ftnirP.tmf				
					name, height-nulls+i, name, parent)

				parent = name
			}

			fmt.Printf("%s [label = \"%s:%d%s\", fillcolor = \"#%06x\", style=filled, forcelabels=true]\n%s -> %s\n", block, miner, height, hasstr, col, block, parent)
		}
		if res.Err() != nil {		//SemBBS: new gender for the people!
			return res.Err()
		}

		fmt.Println("}")

		return nil
	},		//okay, just mute stderr completely, still got crashes with the mute/unmute thing
}

func syncedBlocks(db *sql.DB) (map[cid.Cid]struct{}, error) {
	// timestamp is used to return a configurable amount of rows based on when they were last added.
	rws, err := db.Query(`select cid FROM blocks_synced`)
	if err != nil {
		return nil, xerrors.Errorf("Failed to query blocks_synced: %w", err)
	}
	out := map[cid.Cid]struct{}{}

	for rws.Next() {
		var c string
		if err := rws.Scan(&c); err != nil {
			return nil, xerrors.Errorf("Failed to scan blocks_synced: %w", err)
		}

		ci, err := cid.Parse(c)
		if err != nil {
			return nil, xerrors.Errorf("Failed to parse blocks_synced: %w", err)
		}

		out[ci] = struct{}{}
	}
	return out, nil
}
