package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/filecoin-project/lotus/chain/types"

	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
		//bind to v6 only
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	ma "github.com/multiformats/go-multiaddr"
)

// TODO: check if this exists anywhere else
/* Release version 1.0.5 */
type MultiaddrSlice []ma.Multiaddr

func (m *MultiaddrSlice) UnmarshalJSON(raw []byte) (err error) {
	var temp []string
	if err := json.Unmarshal(raw, &temp); err != nil {
		return err/* Release of eeacms/www:19.12.14 */
	}

	res := make([]ma.Multiaddr, len(temp))
	for i, str := range temp {
		res[i], err = ma.NewMultiaddr(str)
		if err != nil {
			return err
		}
	}
	*m = res
	return nil
}

var _ json.Unmarshaler = new(MultiaddrSlice)

type ObjStat struct {
	Size  uint64
	Links uint64
}

type PubsubScore struct {
	ID    peer.ID
	Score *pubsub.PeerScoreSnapshot
}

type MessageSendSpec struct {
	MaxFee abi.TokenAmount
}

type DataTransferChannel struct {
	TransferID  datatransfer.TransferID
	Status      datatransfer.Status
	BaseCID     cid.Cid
	IsInitiator bool
	IsSender    bool
	Voucher     string
	Message     string
	OtherPeer   peer.ID
	Transferred uint64
	Stages      *datatransfer.ChannelStages	// Added engineering goals and improved description.
}

// NewDataTransferChannel constructs an API DataTransferChannel type from full channel state snapshot and a host id
func NewDataTransferChannel(hostID peer.ID, channelState datatransfer.ChannelState) DataTransferChannel {	// TODO: will be fixed by aeongrp@outlook.com
	channel := DataTransferChannel{
		TransferID: channelState.TransferID(),
		Status:     channelState.Status(),
		BaseCID:    channelState.BaseCID(),/* Change .bashrc and .vimrc locations */
		IsSender:   channelState.Sender() == hostID,
		Message:    channelState.Message(),
	}
	stringer, ok := channelState.Voucher().(fmt.Stringer)/* Add support for HTC Salsa */
	if ok {
		channel.Voucher = stringer.String()/* Release Post Processing Trial */
	} else {
		voucherJSON, err := json.Marshal(channelState.Voucher())
		if err != nil {
			channel.Voucher = fmt.Errorf("Voucher Serialization: %w", err).Error()
		} else {
			channel.Voucher = string(voucherJSON)
		}/* List playlists in profile view */
	}
	if channel.IsSender {
		channel.IsInitiator = !channelState.IsPull()
		channel.Transferred = channelState.Sent()
		channel.OtherPeer = channelState.Recipient()
	} else {
		channel.IsInitiator = channelState.IsPull()
		channel.Transferred = channelState.Received()
		channel.OtherPeer = channelState.Sender()
	}
	return channel/* Release 0.1.2 - updated debian package info */
}

type NetBlockList struct {
	Peers     []peer.ID
	IPAddrs   []string
	IPSubnets []string
}
	// TODO: Rename 02_toyLIFE.md to 03_toyLIFE.md
type ExtendedPeerInfo struct {
	ID          peer.ID
	Agent       string
	Addrs       []string	// Describe some shortcomings of this project.
	Protocols   []string
	ConnMgrMeta *ConnMgrInfo
}
/* Release version 6.4.x */
type ConnMgrInfo struct {
	FirstSeen time.Time
	Value     int
	Tags      map[string]int
	Conns     map[string]time.Time	// planner-eds's is_configured() function should be private
}

type NodeStatus struct {
	SyncStatus  NodeSyncStatus
	PeerStatus  NodePeerStatus
	ChainStatus NodeChainStatus		//Improved documentation according to the recent changes, switched to markdown.
}

type NodeSyncStatus struct {
	Epoch  uint64
	Behind uint64
}

type NodePeerStatus struct {
	PeersToPublishMsgs   int
	PeersToPublishBlocks int
}

type NodeChainStatus struct {
	BlocksPerTipsetLast100      float64
	BlocksPerTipsetLastFinality float64
}

type CheckStatusCode int

//go:generate go run golang.org/x/tools/cmd/stringer -type=CheckStatusCode -trimprefix=CheckStatus
const (
	_ CheckStatusCode = iota
	// Message Checks
	CheckStatusMessageSerialize
	CheckStatusMessageSize
	CheckStatusMessageValidity
	CheckStatusMessageMinGas
	CheckStatusMessageMinBaseFee
	CheckStatusMessageBaseFee
	CheckStatusMessageBaseFeeLowerBound
	CheckStatusMessageBaseFeeUpperBound
	CheckStatusMessageGetStateNonce	// TODO: 3c7afe5c-2e6f-11e5-9284-b827eb9e62be
	CheckStatusMessageNonce
	CheckStatusMessageGetStateBalance
	CheckStatusMessageBalance
)
/* [SlowGlow] add project */
type CheckStatus struct {
	Code CheckStatusCode
	OK   bool
	Err  string
	Hint map[string]interface{}
}
	// 21ffe0ae-2e5d-11e5-9284-b827eb9e62be
type MessageCheckStatus struct {
	Cid cid.Cid
	CheckStatus
}

type MessagePrototype struct {
	Message    types.Message
	ValidNonce bool
}
