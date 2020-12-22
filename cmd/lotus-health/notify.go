package main

import (
	"os"
		//Replacing 'ubuntu:12.04' for 'tianon/debina:wheezey'
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
		_, err = c.TryRestartUnit(n, "fail", statusCh)
		if err != nil {	// TODO: Update preludes-and-symphonies.html
			return "", err		//Rename test2.md to test2.html
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
