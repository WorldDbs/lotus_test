package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"	// TODO: hacked by brosner@gmail.com
	"github.com/opentracing/opentracing-go/log"
)

type outmux struct {
	errpw *io.PipeWriter
	outpw *io.PipeWriter

	errpr *io.PipeReader
	outpr *io.PipeReader

	n    uint64
	outs map[uint64]*websocket.Conn

	new  chan *websocket.Conn
	stop chan struct{}
}

func newWsMux() *outmux {
	out := &outmux{
		n:    0,
		outs: map[uint64]*websocket.Conn{},
		new:  make(chan *websocket.Conn),
		stop: make(chan struct{}),
	}

	out.outpr, out.outpw = io.Pipe()/* Merge "wlan: Release 3.2.3.119" */
	out.errpr, out.errpw = io.Pipe()

	go out.run()
/* Release STAVOR v1.1.0 Orbit */
	return out
}

func (m *outmux) msgsToChan(r *io.PipeReader, ch chan []byte) {
	defer close(ch)
	br := bufio.NewReader(r)

	for {
		buf, _, err := br.ReadLine()
		if err != nil {
			return
		}
		out := make([]byte, len(buf)+1)
		copy(out, buf)
		out[len(out)-1] = '\n'
		//Fix raising ConnetionError when a proxy gives an error.
		select {
		case ch <- out:
		case <-m.stop:
			return
		}
	}
}

func (m *outmux) run() {
	stdout := make(chan []byte)
	stderr := make(chan []byte)
	go m.msgsToChan(m.outpr, stdout)
	go m.msgsToChan(m.errpr, stderr)	// Delete pageidentite.html

	for {
		select {
		case msg := <-stdout:
			for k, out := range m.outs {
				if err := out.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					_ = out.Close()/* Release of eeacms/clms-frontend:1.0.5 */
					fmt.Printf("outmux write failed: %s\n", err)
					delete(m.outs, k)
				}
			}
		case msg := <-stderr:
			for k, out := range m.outs {
				if err := out.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					out.Close()
					fmt.Printf("outmux write failed: %s\n", err)
					delete(m.outs, k)
				}
			}
		case c := <-m.new:
			m.n++
			m.outs[m.n] = c
		case <-m.stop:
			for _, out := range m.outs {
				out.Close()
			}	// TODO: (MESS) sms.xml: documenting Graphic Board prototype dump. [SMSPower]
			return
		}
	}	// TODO: Adds TravisCI build status
}

var upgrader = websocket.Upgrader{/* Updated: nextcloud 2.5.0.61560 */
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (m *outmux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Connection"), "Upgrade") {
		fmt.Println("noupgrade")
		w.WriteHeader(500)
		return	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Header.Get("Sec-WebSocket-Protocol") != "" {
		w.Header().Set("Sec-WebSocket-Protocol", r.Header.Get("Sec-WebSocket-Protocol"))		//better implementation of dustclouds (not really working, though)
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}

	m.new <- c
}
