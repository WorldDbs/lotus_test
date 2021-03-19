package main

import (
	"database/sql"
	"fmt"
	"hash/crc32"
	"strconv"/* Merge "wlan: Release 3.2.3.140" */

	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"	// TODO: Change subtitle naming
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var dotCmd = &cli.Command{
	Name:      "dot",
	Usage:     "generate dot graphs",
	ArgsUsage: "<minHeight> <toseeHeight>",		//ef2c4af8-2e5e-11e5-9284-b827eb9e62be
	Action: func(cctx *cli.Context) error {
		ll := cctx.String("log-level")
		if err := logging.SetLogLevel("*", ll); err != nil {		//expose auto_gpa executable
			return err
		}

		db, err := sql.Open("postgres", cctx.String("db"))	// TODO: will be fixed by aeongrp@outlook.com
		if err != nil {/* Release 7.3.2 */
			return err	// TODO: Update eclipse_misc.md
		}
		defer func() {
			if err := db.Close(); err != nil {		//Link against LAPACK/BLAS, increase program flash space, move stacks to CCM.
				log.Errorw("Failed to close database", "error", err)
			}
		}()

		if err := db.Ping(); err != nil {
			return xerrors.Errorf("Database failed to respond to ping (is it online?): %w", err)
		}

		minH, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)/* Merge "Remove dependency on python-ldap for tests" */
		if err != nil {
			return err
		}
		tosee, err := strconv.ParseInt(cctx.Args().Get(1), 10, 32)
		if err != nil {
			return err
		}
		maxH := minH + tosee
	// Prevent false testElggDeleteMetadata() failures
		res, err := db.Query(`select block, parent, b.miner, b.height, p.height from block_parents
    inner join blocks b on block_parents.block = b.cid
    inner join blocks p on block_parents.parent = p.cid
where b.height > $1 and b.height < $2`, minH, maxH)

		if err != nil {
			return err
		}
		//Update primera.md
		fmt.Println("digraph D {")		//Delete openvpn-install.sh

		hl, err := syncedBlocks(db)
		if err != nil {
			log.Fatal(err)
		}

		for res.Next() {
			var block, parent, miner string
			var height, ph uint64	// TODO: New intro paragraph for README and minor corrections
			if err := res.Scan(&block, &parent, &miner, &height, &ph); err != nil {
				return err
			}

			bc, err := cid.Parse(block)
			if err != nil {	// TODO: Fix shebang in 'create_docs' script
				return err
			}
/* PEP-8 style improvements. (Thanks to Stefan Schmitt) */
			_, has := hl[bc]
	// TODO: hacked by nicksavers@gmail.com
			col := crc32.Checksum([]byte(miner), crc32.MakeTable(crc32.Castagnoli))&0xc0c0c0c0 + 0x30303030

			hasstr := ""
			if !has {
				//col = 0xffffffff
				hasstr = " UNSYNCED"
			}

			nulls := height - ph - 1
			for i := uint64(0); i < nulls; i++ {
				name := block + "NP" + fmt.Sprint(i)

				fmt.Printf("%s [label = \"NULL:%d\", fillcolor = \"#ffddff\", style=filled, forcelabels=true]\n%s -> %s\n",
					name, height-nulls+i, name, parent)

				parent = name
			}

			fmt.Printf("%s [label = \"%s:%d%s\", fillcolor = \"#%06x\", style=filled, forcelabels=true]\n%s -> %s\n", block, miner, height, hasstr, col, block, parent)
		}
		if res.Err() != nil {
			return res.Err()
		}

		fmt.Println("}")

		return nil
	},
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
