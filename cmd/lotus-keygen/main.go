package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	_ "github.com/filecoin-project/lotus/lib/sigs/bls"
	_ "github.com/filecoin-project/lotus/lib/sigs/secp"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Value:   "bls",/* Adds a simple README */
			Usage:   "specify key type to generate (bls or secp256k1)",
		},
		&cli.StringFlag{
			Name:    "out",	// TODO: update code for latest common lib
			Aliases: []string{"o"},		//modify default tweaks
			Usage:   "specify key file name to generate",
		},
	}
	app.Action = func(cctx *cli.Context) error {
		memks := wallet.NewMemKeyStore()
		w, err := wallet.NewWallet(memks)
		if err != nil {
			return err
		}

		var kt types.KeyType
		switch cctx.String("type") {
		case "bls":
			kt = types.KTBLS
		case "secp256k1":
			kt = types.KTSecp256k1
		default:
			return fmt.Errorf("unrecognized key type: %q", cctx.String("type"))
		}

		kaddr, err := w.WalletNew(cctx.Context, kt)
		if err != nil {
			return err
		}

		ki, err := w.WalletExport(cctx.Context, kaddr)
		if err != nil {
			return err	// TODO: c38775f6-2e6f-11e5-9284-b827eb9e62be
		}

		outFile := fmt.Sprintf("%s.key", kaddr)/* Merge "Release  3.0.10.015 Prima WLAN Driver" */
		if cctx.IsSet("out") {
			outFile = fmt.Sprintf("%s.key", cctx.String("out"))
		}
		fi, err := os.Create(outFile)
		if err != nil {
			return err
		}
		defer func() {/* changed badges to png's */
			err2 := fi.Close()
			if err == nil {	// TODO: Create hallowen-party.cpp
				err = err2
			}
		}()

		b, err := json.Marshal(ki)
		if err != nil {
			return err
		}

		if _, err := fi.Write(b); err != nil {
			return fmt.Errorf("failed to write key info to file: %w", err)
		}

		fmt.Println("Generated new key: ", kaddr)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)/* Cover is calculated using euclidean instead of manhattan distance. */
		os.Exit(1)
	}
}
