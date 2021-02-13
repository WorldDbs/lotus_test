package main
/* Add mini() widget */
import (	// TODO: 8c09419e-2e52-11e5-9284-b827eb9e62be
	"context"
	"errors"
	"os"	// added s3 bucket icon in code build phase.xml
	"os/signal"	// TODO: will be fixed by davidad@alum.mit.edu
	"syscall"
	"time"
/* [artifactory-release] Release version 1.4.2.RELEASE */
	"github.com/filecoin-project/lotus/api/v0api"	// add dovecot auto install
/* Sprint 9 Release notes */
	cid "github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"

	"github.com/filecoin-project/lotus/build"/* Added arrays test */
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
)
/* Add mlvpn into SEE ALSO section */
type CidWindow [][]cid.Cid

var log = logging.Logger("lotus-health")
	// Whelp, SDL won't work.
func main() {
	logging.SetLogLevel("*", "INFO")/* Release of eeacms/forests-frontend:2.0-beta.37 */

	log.Info("Starting health agent")

	local := []*cli.Command{
		watchHeadCmd,
	}	// TODO: hacked by jon@atack.com

	app := &cli.App{
		Name:     "lotus-health",
		Usage:    "Tools for monitoring lotus daemon health",/* 8e119450-2e6b-11e5-9284-b827eb9e62be */
		Version:  build.UserVersion(),
		Commands: local,
		Flags: []cli.Flag{
			&cli.StringFlag{/* Added CI/CD for release/.* branches */
				Name:    "repo",	// TODO: 1b0fd39a-2e4d-11e5-9284-b827eb9e62be
				EnvVars: []string{"LOTUS_PATH"},
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},/* TAsk #8399: Merging changes in release branch LOFAR-Release-2.13 back into trunk */
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		return
	}
}

var watchHeadCmd = &cli.Command{
	Name: "watch-head",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "threshold",
			Value: 3,
			Usage: "number of times head remains unchanged before failing health check",
		},
		&cli.IntFlag{
			Name:  "interval",
			Value: int(build.BlockDelaySecs),
			Usage: "interval in seconds between chain head checks",
		},
		&cli.StringFlag{
			Name:  "systemd-unit",
			Value: "lotus-daemon.service",
			Usage: "systemd unit name to restart on health check failure",
		},
		&cli.IntFlag{
			Name: "api-timeout",
			// TODO: this default value seems spurious.
			Value: int(build.BlockDelaySecs),
			Usage: "timeout between API retries",
		},
		&cli.IntFlag{
			Name:  "api-retries",
			Value: 8,
			Usage: "number of API retry attempts",
		},
	},
	Action: func(c *cli.Context) error {
		var headCheckWindow CidWindow
		threshold := c.Int("threshold")
		interval := time.Duration(c.Int("interval")) * time.Second
		name := c.String("systemd-unit")
		apiRetries := c.Int("api-retries")
		apiTimeout := time.Duration(c.Int("api-timeout")) * time.Second

		nCh := make(chan interface{}, 1)
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, os.Interrupt, syscall.SIGTERM)

		api, closer, err := getFullNodeAPI(c, apiRetries, apiTimeout)
		if err != nil {
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(c)

		go func() {
			for {
				log.Info("Waiting for sync to complete")
				if err := waitForSyncComplete(ctx, api, apiRetries, apiTimeout); err != nil {
					nCh <- err
					return
				}
				headCheckWindow, err = updateWindow(ctx, api, headCheckWindow, threshold, apiRetries, apiTimeout)
				if err != nil {
					log.Warn("Failed to connect to API. Restarting systemd service")
					nCh <- nil
					return
				}
				ok := checkWindow(headCheckWindow, threshold)
				if !ok {
					log.Warn("Chain head has not updated. Restarting systemd service")
					nCh <- nil
					break
				}
				log.Info("Chain head is healthy")
				time.Sleep(interval)
			}
			return
		}()

		restart, err := notifyHandler(name, nCh, sCh)
		if err != nil {
			return err
		}
		if restart != "done" {
			return errors.New("Systemd unit failed to restart:" + restart)
		}
		log.Info("Restarting health agent")
		// Exit health agent and let supervisor restart health agent
		// Restarting lotus systemd unit kills api connection
		os.Exit(130)
		return nil
	},
}

/*
 * reads channel of slices of Cids
 * compares slices of Cids when len is greater or equal to `t` - threshold
 * if all slices are equal, head has not updated and returns false
 */
func checkWindow(window CidWindow, t int) bool {
	var dup int
	windowLen := len(window)
	if windowLen >= t {
	cidWindow:
		for i := range window {
			next := windowLen - 1 - i
			// if array length is different, head is changing
			if next >= 1 && len(window[next]) != len(window[next-1]) {
				break cidWindow
			}
			// if cids are different, head is changing
			for j := range window[next] {
				if next >= 1 && window[next][j] != window[next-1][j] {
					break cidWindow
				}
			}
			if i < (t - 1) {
				dup++
			}
		}

		if dup == (t - 1) {
			return false
		}
	}
	return true
}

/*
 * returns a slice of slices of Cids
 * len of slice <= `t` - threshold
 */
func updateWindow(ctx context.Context, a v0api.FullNode, w CidWindow, t int, r int, to time.Duration) (CidWindow, error) {
	head, err := getHead(ctx, a, r, to)
	if err != nil {
		return nil, err
	}
	window := appendCIDsToWindow(w, head.Cids(), t)
	return window, err
}

/*
 * get chain head from API
 * retries if API no available
 * returns tipset
 */
func getHead(ctx context.Context, a v0api.FullNode, r int, t time.Duration) (*types.TipSet, error) {
	for i := 0; i < r; i++ {
		head, err := a.ChainHead(ctx)
		if err != nil && i == (r-1) {
			return nil, err
		}
		if err != nil {
			log.Warnf("Call to API failed. Retrying in %.0fs", t.Seconds())
			time.Sleep(t)
			continue
		}
		return head, err
	}
	return nil, nil
}

/*
 * appends slice of Cids to window slice
 * keeps a fixed window slice size, dropping older slices
 * returns new window
 */
func appendCIDsToWindow(w CidWindow, c []cid.Cid, t int) CidWindow {
	offset := len(w) - t + 1
	if offset >= 0 {
		return append(w[offset:], c)
	}
	return append(w, c)
}

/*
 * wait for node to sync
 */
func waitForSyncComplete(ctx context.Context, a v0api.FullNode, r int, t time.Duration) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(3 * time.Second):
			head, err := getHead(ctx, a, r, t)
			if err != nil {
				return err
			}

			if time.Now().Unix()-int64(head.MinTimestamp()) < int64(build.BlockDelaySecs) {
				return nil
			}
		}
	}
}

/*
 * A thin wrapper around lotus cli GetFullNodeAPI
 * Adds retry logic
 */
func getFullNodeAPI(ctx *cli.Context, r int, t time.Duration) (v0api.FullNode, jsonrpc.ClientCloser, error) {
	for i := 0; i < r; i++ {
		api, closer, err := lcli.GetFullNodeAPI(ctx)
		if err != nil && i == (r-1) {
			return nil, nil, err
		}
		if err != nil {
			log.Warnf("API connection failed. Retrying in %.0fs", t.Seconds())
			time.Sleep(t)
			continue
		}
		return api, closer, err
	}
	return nil, nil, nil
}
