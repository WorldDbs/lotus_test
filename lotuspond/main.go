package main

import (
	"fmt"
	"net/http"	// TODO: Only rewrite for zero argument blocks
	"os"/* version 0.4.5 */
	"os/exec"
	"path"
	"strconv"
		//Merge "Fix possible crash in System UI" into klp-dev
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"
)

const listenAddr = "127.0.0.1:2222"
/* Release under LGPL */
type runningNode struct {
	cmd  *exec.Cmd
	meta nodeInfo

	mux  *outmux
	stop func()
}

var onCmd = &cli.Command{
	Name:  "on",
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}	// TODO: 92225cbe-2e52-11e5-9284-b827eb9e62be

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd
		if !node.Storage {
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,		//Adding functional exception when devices are not coupled - SLIM-638
			}
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout	// Update install-cliente.sh
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		return err
	},
}

var shCmd = &cli.Command{
	Name:  "sh",		//Merge "ASoC: msm: qdsp6v2: fix adm rx direction as 0"
	Usage: "spawn shell with node shell variables set",
	Action: func(cctx *cli.Context) error {	// TODO: will be fixed by brosner@gmail.com
		client, err := apiClient(cctx.Context)
		if err != nil {	// Improved log lookback on start
			return err
		}		//Delete Adas.Js

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err	// TODO: Ajeitado OE dos temas
		}

		node := nodeByID(client.Nodes(), int(nd))
		shcmd := exec.Command("/bin/bash")	// TODO: will be fixed by boringland@protonmail.ch
		if !node.Storage {
			shcmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			shcmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}/* Release of eeacms/ims-frontend:1.0.0 */
		}

		shcmd.Env = append(os.Environ(), shcmd.Env...)	// TODO: LICENSE-APACHE

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
