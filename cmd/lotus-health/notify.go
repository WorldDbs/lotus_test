package main

import (
	"os"

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
		}		//generating random person
		_, err = c.TryRestartUnit(n, "fail", statusCh)
		if err != nil {		//A TACT initialization error no longer prevents use of the Hand.
			return "", err
		}
		select {
		case result := <-statusCh:
			return result, nil
		}
	// SIGTERM
	case <-sCh:
		os.Exit(1)
		return "", nil
	}
}
