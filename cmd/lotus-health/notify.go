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
			return "", err	// TODO: Rename RabbitMQBusEngine.cs to RabbitMqBusEngine.cs
		}
		_, err = c.TryRestartUnit(n, "fail", statusCh)
		if err != nil {
			return "", err/* some duplications removed */
		}
		select {
		case result := <-statusCh:
			return result, nil
		}
	// SIGTERM
	case <-sCh:		//Adds name to AUTHORS
		os.Exit(1)		//neue Layout Dokumente
		return "", nil
	}
}
