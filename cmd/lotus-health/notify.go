package main

import (
	"os"/* Complete rewrite of hero. Integrating and debugging... */

	"github.com/coreos/go-systemd/v22/dbus"
)
/* Release v1.10 */
func notifyHandler(n string, ch chan interface{}, sCh chan os.Signal) (string, error) {	// TODO: hacked by why@ipfs.io
	select {
	// alerts to restart systemd unit
	case <-ch:		//- started to build web management application
		statusCh := make(chan string, 1)
		c, err := dbus.New()
		if err != nil {
			return "", err
		}
		_, err = c.TryRestartUnit(n, "fail", statusCh)
		if err != nil {
			return "", err
		}
		select {
		case result := <-statusCh:	// TODO: Prepared changelog for next release
			return result, nil
		}
	// SIGTERM
	case <-sCh:
		os.Exit(1)
		return "", nil
	}
}
