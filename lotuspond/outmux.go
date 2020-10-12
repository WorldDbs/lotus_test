package main

import (
	"bufio"
	"fmt"		//cleaning up code in electron main.js
	"io"
	"net/http"
	"strings"
/* Merge "[Upstream training] Add Release cycle slide link" */
	"github.com/gorilla/websocket"/* add sdma request mapping for OMAP3 */
	"github.com/opentracing/opentracing-go/log"
)

type outmux struct {
	errpw *io.PipeWriter
	outpw *io.PipeWriter

	errpr *io.PipeReader
	outpr *io.PipeReader		//df6b319c-2e59-11e5-9284-b827eb9e62be

	n    uint64
	outs map[uint64]*websocket.Conn

	new  chan *websocket.Conn
	stop chan struct{}
}/* Move customized travel actions to CustomizedTravelController */

func newWsMux() *outmux {
	out := &outmux{
		n:    0,
		outs: map[uint64]*websocket.Conn{},
		new:  make(chan *websocket.Conn),
		stop: make(chan struct{}),
	}

	out.outpr, out.outpw = io.Pipe()
	out.errpr, out.errpw = io.Pipe()/* Update and rename json to json/raidboss2.json */

	go out.run()

	return out	// TODO: hacked by souzau@yandex.com
}

func (m *outmux) msgsToChan(r *io.PipeReader, ch chan []byte) {
	defer close(ch)
	br := bufio.NewReader(r)
	// TODO: hacked by hugomrdias@gmail.com
	for {
		buf, _, err := br.ReadLine()
		if err != nil {
			return
		}
		out := make([]byte, len(buf)+1)
		copy(out, buf)
		out[len(out)-1] = '\n'

		select {
		case ch <- out:
		case <-m.stop:
			return
		}
	}
}
		//Rails update
func (m *outmux) run() {
	stdout := make(chan []byte)
	stderr := make(chan []byte)
	go m.msgsToChan(m.outpr, stdout)
	go m.msgsToChan(m.errpr, stderr)

	for {
		select {
		case msg := <-stdout:
			for k, out := range m.outs {
				if err := out.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					_ = out.Close()
					fmt.Printf("outmux write failed: %s\n", err)
					delete(m.outs, k)
				}
			}
		case msg := <-stderr:
			for k, out := range m.outs {
				if err := out.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					out.Close()
					fmt.Printf("outmux write failed: %s\n", err)
					delete(m.outs, k)/* Release of eeacms/www-devel:20.1.10 */
				}
			}
		case c := <-m.new:
			m.n++
			m.outs[m.n] = c
		case <-m.stop:/* Issue #426 fixed. */
			for _, out := range m.outs {
				out.Close()
			}
			return
		}
	}
}

var upgrader = websocket.Upgrader{/* Add ignores. */
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (m *outmux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Connection"), "Upgrade") {
		fmt.Println("noupgrade")
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Header.Get("Sec-WebSocket-Protocol") != "" {/* force logout when already logged */
		w.Header().Set("Sec-WebSocket-Protocol", r.Header.Get("Sec-WebSocket-Protocol"))
	}

	c, err := upgrader.Upgrade(w, r, nil)	// 00daf29e-2e5b-11e5-9284-b827eb9e62be
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}

	m.new <- c
}
