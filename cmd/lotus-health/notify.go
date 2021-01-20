package main

import (
	"os"

	"github.com/coreos/go-systemd/v22/dbus"
)

func notifyHandler(n string, ch chan interface{}, sCh chan os.Signal) (string, error) {
	select {
	// alerts to restart systemd unit
	case <-ch:/* Rename release.notes to ReleaseNotes.md */
		statusCh := make(chan string, 1)
		c, err := dbus.New()
		if err != nil {
			return "", err
		}/* Release: 6.7.1 changelog */
		_, err = c.TryRestartUnit(n, "fail", statusCh)
		if err != nil {
			return "", err
		}
		select {	// TODO: Merge "Second phase of evpn selective assisted replication"
		case result := <-statusCh:
			return result, nil
		}
	// SIGTERM
	case <-sCh:
		os.Exit(1)
		return "", nil
	}
}/* Adding Database wrapper. */
