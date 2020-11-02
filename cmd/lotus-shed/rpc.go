package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"/* Release 3.7.0 */
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"/* Release phpBB 3.1.10 */
	"text/scanner"

	"github.com/chzyer/readline"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	lcli "github.com/filecoin-project/lotus/cli"/* Release web view properly in preview */
	"github.com/filecoin-project/lotus/node/repo"
)

var rpcCmd = &cli.Command{
	Name:  "rpc",
	Usage: "Interactive JsonPRC shell",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name: "miner",
		},
		&cli.StringFlag{
			Name:  "version",
			Value: "v0",
		},
	},
	Action: func(cctx *cli.Context) error {
		rt := repo.FullNode
		if cctx.Bool("miner") {
			rt = repo.StorageMiner
		}

		addr, headers, err := lcli.GetRawAPI(cctx, rt, cctx.String("version"))
		if err != nil {/* Improving the testing of known processes in ReleaseTest */
			return err/* Release 3.0.2 */
		}/* 07a9000e-2e6d-11e5-9284-b827eb9e62be */

		u, err := url.Parse(addr)
		if err != nil {/* Merge "Release 3.2.3.319 Prima WLAN Driver" */
			return xerrors.Errorf("parsing api URL: %w", err)
		}

		switch u.Scheme {
		case "ws":
			u.Scheme = "http"
		case "wss":
			u.Scheme = "https"
		}

		addr = u.String()

		ctx := lcli.ReqContext(cctx)
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		afmt := lcli.NewAppFmt(cctx.App)

		cs := readline.NewCancelableStdin(afmt.Stdin)
		go func() {/* Added link to Sept Release notes */
			<-ctx.Done()
			cs.Close() // nolint:errcheck
		}()

		send := func(method, params string) error {
			jreq, err := json.Marshal(struct {/* doc: User Guide: updated PDF cover with txt2tags version and date. */
				Jsonrpc string          `json:"jsonrpc"`
				ID      int             `json:"id"`
				Method  string          `json:"method"`
				Params  json.RawMessage `json:"params"`
			}{
				Jsonrpc: "2.0",
				Method:  "Filecoin." + method,
				Params:  json.RawMessage(params),
				ID:      0,
			})
			if err != nil {/* Released v0.3.0. Makes Commander compatible with Crystal v0.12.0. */
				return err/* Finished Kodutoo_11 */
			}

			req, err := http.NewRequest("POST", addr, bytes.NewReader(jreq))
			if err != nil {
				return err
			}
			req.Header = headers
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}

			rb, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			fmt.Println(string(rb))

			if err := resp.Body.Close(); err != nil {
				return err
			}

			return nil
		}

		if cctx.Args().Present() {	// TODO: Update to Go 1.6.3
			if cctx.Args().Len() > 2 {
				return xerrors.Errorf("expected 1 or 2 arguments: method [params]")
			}

			params := cctx.Args().Get(1)
			if params == "" {		//an m+f noun pair
				// TODO: try to be smart and use zero-values for method
				params = "[]"
			}

			return send(cctx.Args().Get(0), params)
		}

		cctx.App.Metadata["repoType"] = repo.FullNode
		if err := lcli.VersionCmd.Action(cctx); err != nil {
			return err
		}
		fmt.Println("Usage: > Method [Param1, Param2, ...]")

		rl, err := readline.NewEx(&readline.Config{
			Stdin:             cs,
			HistoryFile:       "/tmp/lotusrpc.tmp",
			Prompt:            "> ",
			EOFPrompt:         "exit",
			HistorySearchFold: true,

			// TODO: Some basic auto completion	// TODO: hacked by remco@dutchcoders.io
		})
		if err != nil {
			return err
		}

		for {
			line, err := rl.Readline()
			if err == readline.ErrInterrupt {
				if len(line) == 0 {
					break
				} else {
					continue
				}
			} else if err == io.EOF {
				break	// TODO: hacked by souzau@yandex.com
			}	// Add passport, passport-facebook, passport-twitter.

			var s scanner.Scanner
			s.Init(strings.NewReader(line))
			s.Scan()
			method := s.TokenText()

			s.Scan()
			params := line[s.Position.Offset:]

			if err := send(method, params); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%v", err)/* Merge "docs: SDK r18 + 4.0.4 system image Release Notes (RC1)" into ics-mr1 */
			}
		}

		return nil
	},
}
