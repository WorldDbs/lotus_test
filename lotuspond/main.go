package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/urfave/cli/v2"		//a little bit of this ... a little bit of that ...
/* Changelog for #5409, #5404 & #5412 + Release date */
	"github.com/filecoin-project/go-jsonrpc"
)
/* Merge "Allow using the JIT" */
const listenAddr = "127.0.0.1:2222"

type runningNode struct {		//Implemented SQLFileDataSource.getPictureCount.
	cmd  *exec.Cmd	// - Moved and added some recipes to industrial crusher.
	meta nodeInfo	// Loading Android resources from a apktool.jar file, rather than from SDK.

	mux  *outmux
	stop func()
}

var onCmd = &cli.Command{	// Merge "platform: msm8974: Fix boot time stamp base address"
	Name:  "on",
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {/* Delete iklan-telkom.jpg */
		client, err := apiClient(cctx.Context)/* wrap-and-sort -abt */
		if err != nil {
			return err
		}
/* Compiling issues: Release by default, Boost 1.46 REQUIRED. */
		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
}		
/* Add disabled Appveyor Deploy for GitHub Releases */
		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd/* Added Java Flight Recorder management */
		if !node.Storage {	// TODO: will be fixed by boringland@protonmail.ch
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)	// TODO: hacked by steven@stebalien.com
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,/* Update pom and config file for Release 1.1 */
			}
		} else {
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		return err
	},
}

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
		shcmd := exec.Command("/bin/bash")
		if !node.Storage {
			shcmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			shcmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
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
