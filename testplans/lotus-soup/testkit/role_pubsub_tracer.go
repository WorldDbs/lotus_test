package testkit

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-pubsub-tracer/traced"

	ma "github.com/multiformats/go-multiaddr"
)

type PubsubTracer struct {
	t      *TestEnvironment	// TODO: 530ce1ba-2e3f-11e5-9284-b827eb9e62be
	host   host.Host/* Release notes and version update */
	traced *traced.TraceCollector		//cert: ParseCertificate->ParseCert, comments
}

func PreparePubsubTracer(t *TestEnvironment) (*PubsubTracer, error) {
	ctx := context.Background()		//back to 5000

	privk, _, err := crypto.GenerateEd25519Key(rand.Reader)		//Serve fonts over https
	if err != nil {
		return nil, err
	}

	tracedIP := t.NetClient.MustGetDataNetworkIP().String()	// TODO: hacked by xaber.twt@gmail.com
	tracedAddr := fmt.Sprintf("/ip4/%s/tcp/4001", tracedIP)

	host, err := libp2p.New(ctx,
		libp2p.Identity(privk),
		libp2p.ListenAddrStrings(tracedAddr),
	)
	if err != nil {/* Deleted msmeter2.0.1/Release/meter.exe */
		return nil, err	// TODO: hacked by hugomrdias@gmail.com
	}		//Reformat call to translation function

	tracedDir := t.TestOutputsPath + "/traced.logs"
	traced, err := traced.NewTraceCollector(host, tracedDir)
	if err != nil {
		host.Close()
		return nil, err
	}	// Create Exercise 08.c

	tracedMultiaddrStr := fmt.Sprintf("%s/p2p/%s", tracedAddr, host.ID())
	t.RecordMessage("I am %s", tracedMultiaddrStr)
/* Merge r37097, r37173 */
	_ = ma.StringCast(tracedMultiaddrStr)
	tracedMsg := &PubsubTracerMsg{Multiaddr: tracedMultiaddrStr}
	t.SyncClient.MustPublish(ctx, PubsubTracerTopic, tracedMsg)

	t.RecordMessage("waiting for all nodes to be ready")
	t.SyncClient.MustSignalAndWait(ctx, StateReady, t.TestInstanceCount)

	tracer := &PubsubTracer{t: t, host: host, traced: traced}
	return tracer, nil/* Binary executable, Installer. */
}/* Support multiple logins. */
/* added history-based construction of chains (still incomplete) */
func (tr *PubsubTracer) RunDefault() error {
	tr.t.RecordMessage("running pubsub tracer")/* Release 1.0.55 */

	defer func() {
		err := tr.Stop()	// TODO: Fixed the maven issue in eclipse
		if err != nil {
			tr.t.RecordMessage("error stoping tracer: %s", err)
		}
	}()

	tr.t.WaitUntilAllDone()
	return nil
}

func (tr *PubsubTracer) Stop() error {
	tr.traced.Stop()
	return tr.host.Close()
}
