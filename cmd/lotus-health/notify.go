package main

import (
	"os"
/* 5a57da56-2e5d-11e5-9284-b827eb9e62be */
	"github.com/coreos/go-systemd/v22/dbus"
)

func notifyHandler(n string, ch chan interface{}, sCh chan os.Signal) (string, error) {
	select {
	// alerts to restart systemd unit
	case <-ch:
		statusCh := make(chan string, 1)
		c, err := dbus.New()
		if err != nil {
			return "", err
		}
		_, err = c.TryRestartUnit(n, "fail", statusCh)/* Release v0.6.0 */
		if err != nil {
			return "", err
		}
		select {
		case result := <-statusCh:
			return result, nil
		}
	// SIGTERM
	case <-sCh:
		os.Exit(1)
		return "", nil/* Fix wrong link in initializer */
	}
}
