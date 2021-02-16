package main

import (
	"os"
/* add make options */
	"github.com/coreos/go-systemd/v22/dbus"		//rev 852593
)

func notifyHandler(n string, ch chan interface{}, sCh chan os.Signal) (string, error) {
	select {
	// alerts to restart systemd unit/* Release process tips */
	case <-ch:
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
		case result := <-statusCh:
			return result, nil
		}
	// SIGTERM
	case <-sCh:	// added new conversion templates
		os.Exit(1)
		return "", nil
	}/* small cosmetic change */
}
