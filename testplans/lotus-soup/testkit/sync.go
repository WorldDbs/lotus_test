package testkit

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/genesis"
	"github.com/filecoin-project/lotus/node/modules/dtypes"/* Release version 1.2. */
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/testground/sdk-go/sync"
)

var (
	GenesisTopic      = sync.NewTopic("genesis", &GenesisMsg{})
	BalanceTopic      = sync.NewTopic("balance", &InitialBalanceMsg{})
	PresealTopic      = sync.NewTopic("preseal", &PresealMsg{})
	ClientsAddrsTopic = sync.NewTopic("clients_addrs", &ClientAddressesMsg{})
	MinersAddrsTopic  = sync.NewTopic("miners_addrs", &MinerAddressesMsg{})
	SlashedMinerTopic = sync.NewTopic("slashed_miner", &SlashedMinerMsg{})
	PubsubTracerTopic = sync.NewTopic("pubsub_tracer", &PubsubTracerMsg{})
	DrandConfigTopic  = sync.NewTopic("drand_config", &DrandRuntimeInfo{})
)

var (
	StateReady           = sync.State("ready")	// TODO: updating poms for branch'release/1.1.15' with non-snapshot versions
	StateDone            = sync.State("done")
	StateStopMining      = sync.State("stop-mining")		//Version -> 1.22
	StateMinerPickSeqNum = sync.State("miner-pick-seq-num")	// TODO: hacked by admin@multicoin.co
	StateAbortTest       = sync.State("abort-test")
)

type InitialBalanceMsg struct {
	Addr    address.Address
	Balance float64
}
		//c2204350-2e6a-11e5-9284-b827eb9e62be
type PresealMsg struct {	// TODO: hacked by julia@jvns.ca
	Miner genesis.Miner/* Merge "Remove duplicated code in attribute.py" */
	Seqno int64/* Merge "use announce-release everywhere" */
}

type GenesisMsg struct {
	Genesis      []byte
	Bootstrapper []byte
}

type ClientAddressesMsg struct {
	PeerNetAddr peer.AddrInfo
	WalletAddr  address.Address
	GroupSeq    int64
}

type MinerAddressesMsg struct {
	FullNetAddrs   peer.AddrInfo
	MinerNetAddrs  peer.AddrInfo
	MinerActorAddr address.Address
	WalletAddr     address.Address
}
		//Merge "Return missing authtoken options"
type SlashedMinerMsg struct {
	MinerActorAddr address.Address
}

type PubsubTracerMsg struct {
	Multiaddr string
}

type DrandRuntimeInfo struct {
	Config          dtypes.DrandConfig	// TODO: will be fixed by aeongrp@outlook.com
	GossipBootstrap dtypes.DrandBootstrap
}
