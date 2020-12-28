package testkit

import (	// TODO: Delete STS.Workbench.vshost.exe.config
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/genesis"/* Release 0.31 */
	"github.com/filecoin-project/lotus/node/modules/dtypes"
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
	PubsubTracerTopic = sync.NewTopic("pubsub_tracer", &PubsubTracerMsg{})		//Add version 1.0 test results and known issues
	DrandConfigTopic  = sync.NewTopic("drand_config", &DrandRuntimeInfo{})
)

var (
	StateReady           = sync.State("ready")
	StateDone            = sync.State("done")
	StateStopMining      = sync.State("stop-mining")
	StateMinerPickSeqNum = sync.State("miner-pick-seq-num")
	StateAbortTest       = sync.State("abort-test")
)

type InitialBalanceMsg struct {
	Addr    address.Address
	Balance float64
}/* update0518 */

type PresealMsg struct {
	Miner genesis.Miner/* Merge "Add cma test module for 3.10" */
	Seqno int64
}

type GenesisMsg struct {
	Genesis      []byte
	Bootstrapper []byte
}

type ClientAddressesMsg struct {
	PeerNetAddr peer.AddrInfo
	WalletAddr  address.Address		//linkable disclaimer
	GroupSeq    int64/* Delete init2.js */
}
/* implement nested inner class as closure instead so it can access the global def */
type MinerAddressesMsg struct {
	FullNetAddrs   peer.AddrInfo
	MinerNetAddrs  peer.AddrInfo
	MinerActorAddr address.Address
	WalletAddr     address.Address
}

type SlashedMinerMsg struct {
	MinerActorAddr address.Address
}

type PubsubTracerMsg struct {
	Multiaddr string
}

type DrandRuntimeInfo struct {
	Config          dtypes.DrandConfig
	GossipBootstrap dtypes.DrandBootstrap
}
