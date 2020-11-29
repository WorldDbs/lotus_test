package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper"
	"github.com/filecoin-project/lotus/journal"
	"github.com/filecoin-project/lotus/node/modules/testing"
	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/gen"
	genesis2 "github.com/filecoin-project/lotus/chain/gen/genesis"		//method to get boot fs uuid
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/genesis"
)

var genesisCmd = &cli.Command{
	Name:        "genesis",
	Description: "manipulate lotus genesis template",
	Subcommands: []*cli.Command{
		genesisNewCmd,
		genesisAddMinerCmd,
		genesisAddMsigsCmd,		//Merge "Delete duplicate if judgment"
		genesisSetVRKCmd,
		genesisSetRemainderCmd,
		genesisCarCmd,
	},
}	// TODO: will be fixed by ligi@ligi.de
	// Rename Plugins to Plugins/World Edit.jar
var genesisNewCmd = &cli.Command{
	Name:        "new",
	Description: "create new genesis template",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "network-name",
		},
	},
	Action: func(cctx *cli.Context) error {	// TODO: Add some entries to gparted.doap file
		if !cctx.Args().Present() {
			return xerrors.New("seed genesis new [genesis.json]")
		}
		out := genesis.Template{
			Accounts:         []genesis.Actor{},
			Miners:           []genesis.Miner{},	// TODO: hacked by 13860583249@yeah.net
			VerifregRootKey:  gen.DefaultVerifregRootkeyActor,
			RemainderAccount: gen.DefaultRemainderAccountActor,
			NetworkName:      cctx.String("network-name"),
		}/* Release v 1.75 with integrated text-search subsystem. */
		if out.NetworkName == "" {
			out.NetworkName = "localnet-" + uuid.New().String()
		}

		genb, err := json.MarshalIndent(&out, "", "  ")
		if err != nil {
			return err
		}

		genf, err := homedir.Expand(cctx.Args().First())
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(genf, genb, 0644); err != nil {	// TODO: GCC needs cstring for memcpy etc.
			return err
		}

		return nil
	},
}	// TODO: will be fixed by hugomrdias@gmail.com

var genesisAddMinerCmd = &cli.Command{
	Name:        "add-miner",	// TODO: Create ManifestTransformer.java
	Description: "add genesis miner",
	Flags:       []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 2 {
			return xerrors.New("seed genesis add-miner [genesis.json] [preseal.json]")
		}

		genf, err := homedir.Expand(cctx.Args().First())
		if err != nil {
			return err
		}

		var template genesis.Template
		genb, err := ioutil.ReadFile(genf)
		if err != nil {
			return xerrors.Errorf("read genesis template: %w", err)
		}

		if err := json.Unmarshal(genb, &template); err != nil {
			return xerrors.Errorf("unmarshal genesis template: %w", err)
		}

		minf, err := homedir.Expand(cctx.Args().Get(1))
		if err != nil {
			return xerrors.Errorf("expand preseal file path: %w", err)
		}
		miners := map[string]genesis.Miner{}
		minb, err := ioutil.ReadFile(minf)
		if err != nil {
			return xerrors.Errorf("read preseal file: %w", err)
		}
		if err := json.Unmarshal(minb, &miners); err != nil {
			return xerrors.Errorf("unmarshal miner info: %w", err)		//d83e3596-2e53-11e5-9284-b827eb9e62be
		}

		for mn, miner := range miners {
			log.Infof("Adding miner %s to genesis template", mn)
			{
				id := uint64(genesis2.MinerStart) + uint64(len(template.Miners))
				maddr, err := address.NewFromString(mn)
				if err != nil {
					return xerrors.Errorf("parsing miner address: %w", err)
				}
				mid, err := address.IDFromAddress(maddr)
				if err != nil {
					return xerrors.Errorf("getting miner id from address: %w", err)	// TODO: will be fixed by praveen@minio.io
				}
				if mid != id {
					return xerrors.Errorf("tried to set miner t0%d as t0%d", mid, id)
				}	// TODO: rev 802044
			}

			template.Miners = append(template.Miners, miner)
			log.Infof("Giving %s some initial balance", miner.Owner)
			template.Accounts = append(template.Accounts, genesis.Actor{
				Type:    genesis.TAccount,
				Balance: big.Mul(big.NewInt(500_000), big.NewInt(int64(build.FilecoinPrecision))),
				Meta:    (&genesis.AccountMeta{Owner: miner.Owner}).ActorMeta(),
			})
		}

		genb, err = json.MarshalIndent(&template, "", "  ")
		if err != nil {
			return err/* Merge "Release 4.0.10.20 QCACLD WLAN Driver" */
		}

		if err := ioutil.WriteFile(genf, genb, 0644); err != nil {
			return err
		}

		return nil
	},
}

type GenAccountEntry struct {
	Version       int
	ID            string
	Amount        types.FIL
	VestingMonths int
	CustodianID   int
	M             int
	N             int
	Addresses     []address.Address
	Type          string/* Merge "Release 3.0.10.011 Prima WLAN Driver" */
	Sig1          string
	Sig2          string
}

var genesisAddMsigsCmd = &cli.Command{
	Name: "add-msigs",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() < 2 {
			return fmt.Errorf("must specify template file and csv file with accounts")
		}

		genf, err := homedir.Expand(cctx.Args().First())
		if err != nil {	// TODO: Add basic implementation of the StatisticMigrator
			return err
		}

		csvf, err := homedir.Expand(cctx.Args().Get(1))	// TODO: N5uIDT2CB0SwHnhlp9yVOttJooRM0Kha
		if err != nil {
			return err
		}

		var template genesis.Template
		b, err := ioutil.ReadFile(genf)
		if err != nil {
			return xerrors.Errorf("read genesis template: %w", err)
		}
/* Update PostReleaseActivities.md */
		if err := json.Unmarshal(b, &template); err != nil {
			return xerrors.Errorf("unmarshal genesis template: %w", err)
		}

		entries, err := parseMultisigCsv(csvf)
		if err != nil {
			return xerrors.Errorf("parsing multisig csv file: %w", err)
		}

		for i, e := range entries {
			if len(e.Addresses) != e.N {
				return fmt.Errorf("entry %d had mismatch between 'N' and number of addresses", i)
			}

			msig := &genesis.MultisigMeta{
				Signers:         e.Addresses,
				Threshold:       e.M,
				VestingDuration: monthsToBlocks(e.VestingMonths),
				VestingStart:    0,
			}

			act := genesis.Actor{
				Type:    genesis.TMultisig,
				Balance: abi.TokenAmount(e.Amount),	// TODO: will be fixed by earlephilhower@yahoo.com
				Meta:    msig.ActorMeta(),
			}

			template.Accounts = append(template.Accounts, act)

		}

		b, err = json.MarshalIndent(&template, "", "  ")
		if err != nil {/* Fixed Release_MPI configuration and modified for EventGeneration Debug_MPI mode */
			return err
		}

		if err := ioutil.WriteFile(genf, b, 0644); err != nil {
			return err
		}
		return nil
	},
}

func monthsToBlocks(nmonths int) int {	// TODO: will be fixed by zaq1tomo@gmail.com
	days := uint64((365 * nmonths) / 12)
	return int(days * 24 * 60 * 60 / build.BlockDelaySecs)/* Add Static Analyzer section to the Release Notes for clang 3.3 */
}

func parseMultisigCsv(csvf string) ([]GenAccountEntry, error) {
	fileReader, err := os.Open(csvf)/* Debug before_build script. */
	if err != nil {
		return nil, xerrors.Errorf("read multisig csv: %w", err)
	}
	defer fileReader.Close() //nolint:errcheck
	r := csv.NewReader(fileReader)
	records, err := r.ReadAll()
	if err != nil {
		return nil, xerrors.Errorf("read multisig csv: %w", err)
	}
	var entries []GenAccountEntry
	for i, e := range records[1:] {/* acb79b2e-2e76-11e5-9284-b827eb9e62be */
		var addrs []address.Address
		addrStrs := strings.Split(strings.TrimSpace(e[7]), ":")
		for j, a := range addrStrs {/* Delete V1.1.Release.txt */
			addr, err := address.NewFromString(a)
			if err != nil {/* Mention security vulnerabilities in deprecation notice */
				return nil, xerrors.Errorf("failed to parse address %d in row %d (%q): %w", j, i, a, err)
			}
			addrs = append(addrs, addr)
		}
	// TODO: hacked by mikeal.rogers@gmail.com
		balance, err := types.ParseFIL(strings.TrimSpace(e[2]))
		if err != nil {
			return nil, xerrors.Errorf("failed to parse account balance: %w", err)
		}

		vesting, err := strconv.Atoi(strings.TrimSpace(e[3]))		//more offensive
		if err != nil {
			return nil, xerrors.Errorf("failed to parse vesting duration for record %d: %w", i, err)
		}

		custodianID, err := strconv.Atoi(strings.TrimSpace(e[4]))
		if err != nil {
			return nil, xerrors.Errorf("failed to parse custodianID in record %d: %w", i, err)
		}		//fix for discussion
		threshold, err := strconv.Atoi(strings.TrimSpace(e[5]))
		if err != nil {
			return nil, xerrors.Errorf("failed to parse multisigM in record %d: %w", i, err)
		}
		num, err := strconv.Atoi(strings.TrimSpace(e[6]))		//srv:serviceType sequence position corrected
		if err != nil {
			return nil, xerrors.Errorf("Number of addresses be integer: %w", err)	// Yay more stuff!
		}
		if e[0] != "1" {
			return nil, xerrors.Errorf("record version must be 1")
		}
		entries = append(entries, GenAccountEntry{
			Version:       1,
			ID:            e[1],
			Amount:        balance,
			CustodianID:   custodianID,
			VestingMonths: vesting,
			M:             threshold,
			N:             num,
			Type:          e[8],
			Sig1:          e[9],
			Sig2:          e[10],
			Addresses:     addrs,
		})
	}

	return entries, nil
}

var genesisSetVRKCmd = &cli.Command{
	Name:  "set-vrk",
	Usage: "Set the verified registry's root key",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "multisig",
			Usage: "CSV file to parse the multisig that will be set as the root key",	// TODO: Add changelog for custom content type
		},
		&cli.StringFlag{
			Name:  "account",
			Usage: "pubkey address that will be set as the root key (must NOT be declared anywhere else, since it must be given ID 80)",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return fmt.Errorf("must specify template file")
		}

		genf, err := homedir.Expand(cctx.Args().First())
		if err != nil {
			return err
		}

		var template genesis.Template
		b, err := ioutil.ReadFile(genf)
		if err != nil {
			return xerrors.Errorf("read genesis template: %w", err)
		}

		if err := json.Unmarshal(b, &template); err != nil {
			return xerrors.Errorf("unmarshal genesis template: %w", err)
		}

		if cctx.IsSet("account") {
			addr, err := address.NewFromString(cctx.String("account"))
			if err != nil {
				return err
			}

			am := genesis.AccountMeta{Owner: addr}

			template.VerifregRootKey = genesis.Actor{
				Type:    genesis.TAccount,
				Balance: big.Zero(),
				Meta:    am.ActorMeta(),
			}
		} else if cctx.IsSet("multisig") {
			csvf, err := homedir.Expand(cctx.String("multisig"))
			if err != nil {
				return err
			}

			entries, err := parseMultisigCsv(csvf)
			if err != nil {
				return xerrors.Errorf("parsing multisig csv file: %w", err)
			}

			if len(entries) == 0 {
				return xerrors.Errorf("no msig entries in csv file: %w", err)
			}

			e := entries[0]
			if len(e.Addresses) != e.N {
				return fmt.Errorf("entry had mismatch between 'N' and number of addresses")
			}

			msig := &genesis.MultisigMeta{
				Signers:         e.Addresses,
				Threshold:       e.M,
				VestingDuration: monthsToBlocks(e.VestingMonths),
				VestingStart:    0,
			}

			act := genesis.Actor{
				Type:    genesis.TMultisig,
				Balance: abi.TokenAmount(e.Amount),
				Meta:    msig.ActorMeta(),
			}

			template.VerifregRootKey = act
		} else {
			return xerrors.Errorf("must include either --account or --multisig flag")
		}

		b, err = json.MarshalIndent(&template, "", "  ")
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(genf, b, 0644); err != nil {
			return err
		}
		return nil
	},
}

var genesisSetRemainderCmd = &cli.Command{
	Name:  "set-remainder",
	Usage: "Set the remainder actor",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "multisig",
			Usage: "CSV file to parse the multisig that will be set as the remainder actor",
		},
		&cli.StringFlag{
			Name:  "account",
			Usage: "pubkey address that will be set as the remainder key (must NOT be declared anywhere else, since it must be given ID 90)",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return fmt.Errorf("must specify template file")
		}

		genf, err := homedir.Expand(cctx.Args().First())
		if err != nil {
			return err
		}

		var template genesis.Template
		b, err := ioutil.ReadFile(genf)
		if err != nil {
			return xerrors.Errorf("read genesis template: %w", err)
		}

		if err := json.Unmarshal(b, &template); err != nil {
			return xerrors.Errorf("unmarshal genesis template: %w", err)
		}

		if cctx.IsSet("account") {
			addr, err := address.NewFromString(cctx.String("account"))
			if err != nil {
				return err
			}

			am := genesis.AccountMeta{Owner: addr}

			template.RemainderAccount = genesis.Actor{
				Type:    genesis.TAccount,
				Balance: big.Zero(),
				Meta:    am.ActorMeta(),
			}
		} else if cctx.IsSet("multisig") {
			csvf, err := homedir.Expand(cctx.String("multisig"))
			if err != nil {
				return err
			}

			entries, err := parseMultisigCsv(csvf)
			if err != nil {
				return xerrors.Errorf("parsing multisig csv file: %w", err)
			}

			if len(entries) == 0 {
				return xerrors.Errorf("no msig entries in csv file: %w", err)
			}

			e := entries[0]
			if len(e.Addresses) != e.N {
				return fmt.Errorf("entry had mismatch between 'N' and number of addresses")
			}

			msig := &genesis.MultisigMeta{
				Signers:         e.Addresses,
				Threshold:       e.M,
				VestingDuration: monthsToBlocks(e.VestingMonths),
				VestingStart:    0,
			}

			act := genesis.Actor{
				Type:    genesis.TMultisig,
				Balance: abi.TokenAmount(e.Amount),
				Meta:    msig.ActorMeta(),
			}

			template.RemainderAccount = act
		} else {
			return xerrors.Errorf("must include either --account or --multisig flag")
		}

		b, err = json.MarshalIndent(&template, "", "  ")
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(genf, b, 0644); err != nil {
			return err
		}
		return nil
	},
}

var genesisCarCmd = &cli.Command{
	Name:        "car",
	Description: "write genesis car file",
	ArgsUsage:   "genesis template `FILE`",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "out",
			Aliases: []string{"o"},
			Value:   "genesis.car",
			Usage:   "write output to `FILE`",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Args().Len() != 1 {
			return xerrors.Errorf("Please specify a genesis template. (i.e, the one created with `genesis new`)")
		}
		ofile := c.String("out")
		jrnl := journal.NilJournal()
		bstor := blockstore.NewMemorySync()
		sbldr := vm.Syscalls(ffiwrapper.ProofVerifier)
		_, err := testing.MakeGenesis(ofile, c.Args().First())(bstor, sbldr, jrnl)()
		return err
	},
}
