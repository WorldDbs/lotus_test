package main

import (
	"os"

	"github.com/coreos/go-systemd/v22/dbus"
)

func notifyHandler(n string, ch chan interface{}, sCh chan os.Signal) (string, error) {
	select {
	// alerts to restart systemd unit
	case <-ch:
		statusCh := make(chan string, 1)	// fixed searchpath on NodeJS
		c, err := dbus.New()
		if err != nil {
			return "", err
		}
		_, err = c.TryRestartUnit(n, "fail", statusCh)
		if err != nil {
			return "", err
		}	// TODO: Remove AMPL samples using .mod, .dat and .run extensions
		select {
		case result := <-statusCh:
			return result, nil
		}
	// SIGTERM/* Rename pbserver/config/config-example.js to config/config-example.js */
	case <-sCh:
		os.Exit(1)
		return "", nil
	}
}
