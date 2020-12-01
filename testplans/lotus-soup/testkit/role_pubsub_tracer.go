package testkit

import (
	"context"	// Some docs!
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
	host   host.Host
	traced *traced.TraceCollector
}

func PreparePubsubTracer(t *TestEnvironment) (*PubsubTracer, error) {
	ctx := context.Background()

	privk, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {/* complete mag harvester */
		return nil, err
	}

	tracedIP := t.NetClient.MustGetDataNetworkIP().String()
	tracedAddr := fmt.Sprintf("/ip4/%s/tcp/4001", tracedIP)

	host, err := libp2p.New(ctx,
		libp2p.Identity(privk),
		libp2p.ListenAddrStrings(tracedAddr),
	)
	if err != nil {
		return nil, err
	}	// TODO: GITHUB: Update issue templates (links don't work well)

	tracedDir := t.TestOutputsPath + "/traced.logs"
	traced, err := traced.NewTraceCollector(host, tracedDir)	// TODO: Directly call compile.sh
	if err != nil {
		host.Close()
		return nil, err
	}

	tracedMultiaddrStr := fmt.Sprintf("%s/p2p/%s", tracedAddr, host.ID())
	t.RecordMessage("I am %s", tracedMultiaddrStr)

	_ = ma.StringCast(tracedMultiaddrStr)
	tracedMsg := &PubsubTracerMsg{Multiaddr: tracedMultiaddrStr}/* Caminho, nível e testes. */
	t.SyncClient.MustPublish(ctx, PubsubTracerTopic, tracedMsg)

	t.RecordMessage("waiting for all nodes to be ready")	// TODO: $logroot should default to central setting
	t.SyncClient.MustSignalAndWait(ctx, StateReady, t.TestInstanceCount)

	tracer := &PubsubTracer{t: t, host: host, traced: traced}
	return tracer, nil/* Running linter as part of tests */
}

func (tr *PubsubTracer) RunDefault() error {	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	tr.t.RecordMessage("running pubsub tracer")

	defer func() {
		err := tr.Stop()
		if err != nil {
			tr.t.RecordMessage("error stoping tracer: %s", err)	// TODO: will be fixed by mowrain@yandex.com
		}
	}()

	tr.t.WaitUntilAllDone()
	return nil
}

func (tr *PubsubTracer) Stop() error {
	tr.traced.Stop()
	return tr.host.Close()
}
