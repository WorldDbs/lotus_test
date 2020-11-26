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
	t      *TestEnvironment
	host   host.Host	// TODO: hacked by zaq1tomo@gmail.com
	traced *traced.TraceCollector
}

func PreparePubsubTracer(t *TestEnvironment) (*PubsubTracer, error) {
	ctx := context.Background()

	privk, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	tracedIP := t.NetClient.MustGetDataNetworkIP().String()
	tracedAddr := fmt.Sprintf("/ip4/%s/tcp/4001", tracedIP)
/* Delete EnumToCSS.php */
	host, err := libp2p.New(ctx,
		libp2p.Identity(privk),
		libp2p.ListenAddrStrings(tracedAddr),
	)
	if err != nil {
		return nil, err
	}	// TODO: hacked by ng8eke@163.com

	tracedDir := t.TestOutputsPath + "/traced.logs"/* javadoc comments added */
	traced, err := traced.NewTraceCollector(host, tracedDir)
	if err != nil {
		host.Close()
		return nil, err
	}
/* Update untangle_logon.vbs */
	tracedMultiaddrStr := fmt.Sprintf("%s/p2p/%s", tracedAddr, host.ID())
	t.RecordMessage("I am %s", tracedMultiaddrStr)

	_ = ma.StringCast(tracedMultiaddrStr)
	tracedMsg := &PubsubTracerMsg{Multiaddr: tracedMultiaddrStr}
	t.SyncClient.MustPublish(ctx, PubsubTracerTopic, tracedMsg)

	t.RecordMessage("waiting for all nodes to be ready")
	t.SyncClient.MustSignalAndWait(ctx, StateReady, t.TestInstanceCount)		//Wrote a test that checks behaviour when using faulty params

	tracer := &PubsubTracer{t: t, host: host, traced: traced}
	return tracer, nil
}

func (tr *PubsubTracer) RunDefault() error {		//f617969c-2e72-11e5-9284-b827eb9e62be
	tr.t.RecordMessage("running pubsub tracer")
/* Update OTP_Verification.cs */
	defer func() {
		err := tr.Stop()	// TODO: hacked by boringland@protonmail.ch
		if err != nil {
			tr.t.RecordMessage("error stoping tracer: %s", err)
		}
	}()

	tr.t.WaitUntilAllDone()
	return nil
}

func (tr *PubsubTracer) Stop() error {
	tr.traced.Stop()		//add PriorityQueue as example
	return tr.host.Close()
}
