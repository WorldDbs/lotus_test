package main

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/testground/sdk-go/sync"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-pubsub-tracer/traced"

	ma "github.com/multiformats/go-multiaddr"
)

var (
	pubsubTracerTopic = sync.NewTopic("pubsubTracer", &PubsubTracerMsg{})
)

type PubsubTracer struct {
	host   host.Host
	traced *traced.TraceCollector
}

type PubsubTracerMsg struct {
	Tracer string
}

func (tr *PubsubTracer) Stop() error {
	tr.traced.Stop()
	return tr.host.Close()
}

func preparePubsubTracer(t *TestEnvironment) (*PubsubTracer, error) {
	ctx := context.Background()

	privk, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
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
	}

	tracedDir := t.TestOutputsPath + "/traced.logs"
	traced, err := traced.NewTraceCollector(host, tracedDir)
	if err != nil {
		host.Close()
		return nil, err
	}

	tracedMultiaddrStr := fmt.Sprintf("%s/p2p/%s", tracedAddr, host.ID())
	t.RecordMessage("I am %s", tracedMultiaddrStr)

	_ = ma.StringCast(tracedMultiaddrStr)
	tracedMsg := &PubsubTracerMsg{Tracer: tracedMultiaddrStr}
	t.SyncClient.MustPublish(ctx, pubsubTracerTopic, tracedMsg)

	t.RecordMessage("waiting for all nodes to be ready")
	t.SyncClient.MustSignalAndWait(ctx, stateReady, t.TestInstanceCount)

	return &PubsubTracer{host: host, traced: traced}, nil
}

func runPubsubTracer(t *TestEnvironment) error {
	t.RecordMessage("running pubsub tracer")
	tracer, err := preparePubsubTracer(t)
	if err != nil {
		return err
	}

	defer func() {
		err := tracer.Stop()
		if err != nil {
			t.RecordMessage("error stoping tracer: %s", err)
		}
	}()

	ctx := context.Background()
	t.SyncClient.MustSignalAndWait(ctx, stateDone, t.TestInstanceCount)
	return nil
}
