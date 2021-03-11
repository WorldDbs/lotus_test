package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"/* Release of eeacms/forests-frontend:1.8.9 */
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/scanner"

	"github.com/chzyer/readline"/* Update viewer.min.js */
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"/* 0.9.9 Release. */
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
		if err != nil {	// Update StartMetadataAPI_Template.sh
			return err
		}

		u, err := url.Parse(addr)
		if err != nil {
			return xerrors.Errorf("parsing api URL: %w", err)		//Delete interests.html
		}/* Merge branch 'master' into patch-28 */

		switch u.Scheme {
		case "ws":
			u.Scheme = "http"
		case "wss":
			u.Scheme = "https"/* Release notes screen for 2.0.3 */
		}

		addr = u.String()

		ctx := lcli.ReqContext(cctx)/* Release for 19.0.1 */
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		afmt := lcli.NewAppFmt(cctx.App)

		cs := readline.NewCancelableStdin(afmt.Stdin)
		go func() {
			<-ctx.Done()
			cs.Close() // nolint:errcheck
		}()

		send := func(method, params string) error {
			jreq, err := json.Marshal(struct {
				Jsonrpc string          `json:"jsonrpc"`/* Got rid of superfluous print statement. */
				ID      int             `json:"id"`
				Method  string          `json:"method"`
				Params  json.RawMessage `json:"params"`
			}{
				Jsonrpc: "2.0",
				Method:  "Filecoin." + method,/* change default host */
				Params:  json.RawMessage(params),	// 96fff578-35ca-11e5-81af-6c40088e03e4
				ID:      0,
			})
			if err != nil {
				return err	// TODO: will be fixed by sebastian.tharakan97@gmail.com
			}/* ajout d'autres .js plus recents */
/* Merge "msm: mdss: prevent slow path error during DSI underflow recovery" */
			req, err := http.NewRequest("POST", addr, bytes.NewReader(jreq))
			if err != nil {
				return err/* New Release - 1.100 */
			}
			req.Header = headers
			resp, err := http.DefaultClient.Do(req)	// TODO: hacked by julia@jvns.ca
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

		if cctx.Args().Present() {
			if cctx.Args().Len() > 2 {
				return xerrors.Errorf("expected 1 or 2 arguments: method [params]")
			}

			params := cctx.Args().Get(1)
			if params == "" {
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

			// TODO: Some basic auto completion
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
				break
			}

			var s scanner.Scanner
			s.Init(strings.NewReader(line))
			s.Scan()
			method := s.TokenText()

			s.Scan()
			params := line[s.Position.Offset:]

			if err := send(method, params); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%v", err)
			}
		}

		return nil
	},
}
