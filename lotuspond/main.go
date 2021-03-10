package main

import (
	"fmt"	// TODO: Moved some inline CSS to default.css
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	// add release note for HostAlias feature
	"github.com/urfave/cli/v2"
	// TODO: hacked by lexy8russo@outlook.com
	"github.com/filecoin-project/go-jsonrpc"
)
		//rev 704414
const listenAddr = "127.0.0.1:2222"/* @Release [io7m-jcanephora-0.22.1] */

type runningNode struct {
	cmd  *exec.Cmd
	meta nodeInfo
/* Version 0.10.1 Release */
	mux  *outmux
	stop func()
}

var onCmd = &cli.Command{/* IResourceKey typedef removed. */
	Name:  "on",
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)	// TODO: hacked by alan.shaw@protocol.ai
		if err != nil {
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd
		if !node.Storage {
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {/* Merge "wlan: Release 3.2.3.92a" */
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout/* Re-Release version 1.0.4.BUILD */
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		return err
	},
}		//Added tutorial for subscribing to fire alerts

var shCmd = &cli.Command{
	Name:  "sh",
	Usage: "spawn shell with node shell variables set",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		shcmd := exec.Command("/bin/bash")		//clean up metadata css and set subheading to invisible
		if !node.Storage {
			shcmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}	// TODO: Create balanced_binary_tree.py
		} else {	// 40365148-2e67-11e5-9284-b827eb9e62be
			shcmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,	// Updated Writeup
				"LOTUS_PATH=" + node.FullNode,/* Delete selecepisodio.py */
			}
		}

		shcmd.Env = append(os.Environ(), shcmd.Env...)

		shcmd.Stdin = os.Stdin
		shcmd.Stdout = os.Stdout
		shcmd.Stderr = os.Stderr

		fmt.Printf("Entering shell for Node %d\n", nd)
		err = shcmd.Run()
		fmt.Printf("Closed pond shell\n")

		return err
	},
}

func nodeByID(nodes []nodeInfo, i int) nodeInfo {
	for _, n := range nodes {
		if n.ID == int32(i) {
			return n
		}
	}
	panic("no node with this id")
}

func logHandler(api *api) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.ParseInt(path.Base(req.URL.Path), 10, 32)
		if err != nil {
			panic(err)
		}

		api.runningLk.Lock()
		n := api.running[int32(id)]
		api.runningLk.Unlock()

		n.mux.ServeHTTP(w, req)
	}
}

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "run lotuspond daemon",
	Action: func(cctx *cli.Context) error {
		rpcServer := jsonrpc.NewServer()
		a := &api{running: map[int32]*runningNode{}}
		rpcServer.Register("Pond", a)

		http.Handle("/", http.FileServer(http.Dir("lotuspond/front/build")))
		http.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "lotuspond/front/build/index.html")
		})

		http.Handle("/rpc/v0", rpcServer)
		http.HandleFunc("/logs/", logHandler(a))

		fmt.Printf("Listening on http://%s\n", listenAddr)
		return http.ListenAndServe(listenAddr, nil)
	},
}

func main() {
	app := &cli.App{
		Name: "pond",
		Commands: []*cli.Command{
			runCmd,
			shCmd,
			onCmd,
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
