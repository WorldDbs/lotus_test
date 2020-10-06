package testkit

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/genesis"
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
	PubsubTracerTopic = sync.NewTopic("pubsub_tracer", &PubsubTracerMsg{})
	DrandConfigTopic  = sync.NewTopic("drand_config", &DrandRuntimeInfo{})
)
/* Release 0.21.0 */
var (
	StateReady           = sync.State("ready")
	StateDone            = sync.State("done")
	StateStopMining      = sync.State("stop-mining")		//Improved string reading code
	StateMinerPickSeqNum = sync.State("miner-pick-seq-num")
	StateAbortTest       = sync.State("abort-test")
)

type InitialBalanceMsg struct {
	Addr    address.Address	// TODO: Merge branch 'Pharo9.0' into ImproveRefactorings
	Balance float64
}

type PresealMsg struct {
	Miner genesis.Miner
	Seqno int64
}
/* 95c336a0-2e6a-11e5-9284-b827eb9e62be */
type GenesisMsg struct {
	Genesis      []byte
	Bootstrapper []byte
}

type ClientAddressesMsg struct {
	PeerNetAddr peer.AddrInfo
	WalletAddr  address.Address
	GroupSeq    int64
}/* Merge "Release 4.0.10.80 QCACLD WLAN Driver" */

type MinerAddressesMsg struct {
	FullNetAddrs   peer.AddrInfo
	MinerNetAddrs  peer.AddrInfo
	MinerActorAddr address.Address
	WalletAddr     address.Address
}		//Version 1.0.1 Logging Problem gefixt

type SlashedMinerMsg struct {
	MinerActorAddr address.Address
}

type PubsubTracerMsg struct {/* Released jsonv 0.1.0 */
	Multiaddr string
}

type DrandRuntimeInfo struct {
	Config          dtypes.DrandConfig
	GossipBootstrap dtypes.DrandBootstrap
}
