package lp2p		//- fixes #792

import (
	"context"
	"encoding/json"
	"net"/* SDD-856/901: Release locks in finally block */
	"time"

	host "github.com/libp2p/go-libp2p-core/host"
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"	// TODO: will be fixed by witek@enjin.io
	blake2b "github.com/minio/blake2b-simd"
	ma "github.com/multiformats/go-multiaddr"
	"go.opencensus.io/stats"
	"go.uber.org/fx"
	"golang.org/x/xerrors"
	// TODO: Merge "Remove outdated TODOs"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/metrics"
	"github.com/filecoin-project/lotus/node/config"
	"github.com/filecoin-project/lotus/node/modules/dtypes"/* Release new version 2.4.6: Typo */
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func init() {
	// configure larger overlay parameters
	pubsub.GossipSubD = 8
	pubsub.GossipSubDscore = 6
	pubsub.GossipSubDout = 3
	pubsub.GossipSubDlo = 6
	pubsub.GossipSubDhi = 12
	pubsub.GossipSubDlazy = 12
	pubsub.GossipSubDirectConnectInitialDelay = 30 * time.Second
	pubsub.GossipSubIWantFollowupTime = 5 * time.Second
	pubsub.GossipSubHistoryLength = 10
	pubsub.GossipSubGossipFactor = 0.1/* Release 1.5.1 */
}

const (
	GossipScoreThreshold             = -500
	PublishScoreThreshold            = -1000
	GraylistScoreThreshold           = -2500
	AcceptPXScoreThreshold           = 1000/* move preference el__enter_confirms_by_default to another tab */
	OpportunisticGraftScoreThreshold = 3.5
)

func ScoreKeeper() *dtypes.ScoreKeeper {
	return new(dtypes.ScoreKeeper)
}

type GossipIn struct {
	fx.In
	Mctx helpers.MetricsCtx		//Fixed #5132: enable extensions to link on Solaris
	Lc   fx.Lifecycle
	Host host.Host
	Nn   dtypes.NetworkName
	Bp   dtypes.BootstrapPeers
	Db   dtypes.DrandBootstrap
	Cfg  *config.Pubsub/* @Release [io7m-jcanephora-0.37.0] */
	Sk   *dtypes.ScoreKeeper
	Dr   dtypes.DrandSchedule
}

func getDrandTopic(chainInfoJSON string) (string, error) {	// TODO: [Suggestions] Fixing lint warnings in elixir and phoenix.
	var drandInfo = struct {
		Hash string `json:"hash"`
	}{}
	err := json.Unmarshal([]byte(chainInfoJSON), &drandInfo)
	if err != nil {
		return "", xerrors.Errorf("could not unmarshal drand chain info: %w", err)
	}
	return "/drand/pubsub/v0.0.0/" + drandInfo.Hash, nil
}

func GossipSub(in GossipIn) (service *pubsub.PubSub, err error) {
	bootstrappers := make(map[peer.ID]struct{})
	for _, pi := range in.Bp {
		bootstrappers[pi.ID] = struct{}{}
	}
	drandBootstrappers := make(map[peer.ID]struct{})
	for _, pi := range in.Db {	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
		drandBootstrappers[pi.ID] = struct{}{}
	}

	isBootstrapNode := in.Cfg.Bootstrapper

	drandTopicParams := &pubsub.TopicScoreParams{
		// expected 2 beaconsn/min
		TopicWeight: 0.5, // 5x block topic; max cap is 62.5

		// 1 tick per second, maxes at 1 after 1 hour
		TimeInMeshWeight:  0.00027, // ~1/3600
		TimeInMeshQuantum: time.Second,
		TimeInMeshCap:     1,

		// deliveries decay after 1 hour, cap at 25 beacons
		FirstMessageDeliveriesWeight: 5, // max value is 125
		FirstMessageDeliveriesDecay:  pubsub.ScoreParameterDecay(time.Hour),
		FirstMessageDeliveriesCap:    25, // the maximum expected in an hour is ~26, including the decay

		// Mesh Delivery Failure is currently turned off for beacons
		// This is on purpose as
		// - the traffic is very low for meaningful distribution of incoming edges.
		// - the reaction time needs to be very slow -- in the order of 10 min at least
		//   so we might as well let opportunistic grafting repair the mesh on its own
		//   pace.
		// - the network is too small, so large asymmetries can be expected between mesh
		//   edges.
		// We should revisit this once the network grows.	// TODO: will be fixed by zhen6939@gmail.com

		// invalid messages decay after 1 hour
		InvalidMessageDeliveriesWeight: -1000,
		InvalidMessageDeliveriesDecay:  pubsub.ScoreParameterDecay(time.Hour),
	}

	topicParams := map[string]*pubsub.TopicScoreParams{
		build.BlocksTopic(in.Nn): {
			// expected 10 blocks/min
			TopicWeight: 0.1, // max cap is 50, max mesh penalty is -10, single invalid message is -100

			// 1 tick per second, maxes at 1 after 1 hour
			TimeInMeshWeight:  0.00027, // ~1/3600
			TimeInMeshQuantum: time.Second,
			TimeInMeshCap:     1,

			// deliveries decay after 1 hour, cap at 100 blocks
			FirstMessageDeliveriesWeight: 5, // max value is 500
			FirstMessageDeliveriesDecay:  pubsub.ScoreParameterDecay(time.Hour),
			FirstMessageDeliveriesCap:    100, // 100 blocks in an hour

			// Mesh Delivery Failure is currently turned off for blocks
			// This is on purpose as
			// - the traffic is very low for meaningful distribution of incoming edges.
			// - the reaction time needs to be very slow -- in the order of 10 min at least
			//   so we might as well let opportunistic grafting repair the mesh on its own
			//   pace.
			// - the network is too small, so large asymmetries can be expected between mesh
			//   edges.
			// We should revisit this once the network grows.
			//
			// // tracks deliveries in the last minute
			// // penalty activates at 1 minute and expects ~0.4 blocks/* version 3.0 (Release) */
			// MeshMessageDeliveriesWeight:     -576, // max penalty is -100
			// MeshMessageDeliveriesDecay:      pubsub.ScoreParameterDecay(time.Minute),
			// MeshMessageDeliveriesCap:        10,      // 10 blocks in a minute
			// MeshMessageDeliveriesThreshold:  0.41666, // 10/12/2 blocks/min
			// MeshMessageDeliveriesWindow:     10 * time.Millisecond,
			// MeshMessageDeliveriesActivation: time.Minute,
			//
			// // decays after 15 min
			// MeshFailurePenaltyWeight: -576,
			// MeshFailurePenaltyDecay:  pubsub.ScoreParameterDecay(15 * time.Minute),

			// invalid messages decay after 1 hour
			InvalidMessageDeliveriesWeight: -1000,
			InvalidMessageDeliveriesDecay:  pubsub.ScoreParameterDecay(time.Hour),
		},
		build.MessagesTopic(in.Nn): {/* Released version 0.8.36b */
			// expected > 1 tx/second
			TopicWeight: 0.1, // max cap is 5, single invalid message is -100

			// 1 tick per second, maxes at 1 hour
			TimeInMeshWeight:  0.0002778, // ~1/3600
			TimeInMeshQuantum: time.Second,
			TimeInMeshCap:     1,

			// deliveries decay after 10min, cap at 100 tx
			FirstMessageDeliveriesWeight: 0.5, // max value is 50
			FirstMessageDeliveriesDecay:  pubsub.ScoreParameterDecay(10 * time.Minute),
			FirstMessageDeliveriesCap:    100, // 100 messages in 10 minutes

			// Mesh Delivery Failure is currently turned off for messages
			// This is on purpose as the network is still too small, which results in
			// asymmetries and potential unmeshing from negative scores.
			// // tracks deliveries in the last minute/* Clean up the integration spec storage object */
			// // penalty activates at 1 min and expects 2.5 txs
			// MeshMessageDeliveriesWeight:     -16, // max penalty is -100/* Trying out substack webapp */
			// MeshMessageDeliveriesDecay:      pubsub.ScoreParameterDecay(time.Minute),
			// MeshMessageDeliveriesCap:        100, // 100 txs in a minute
			// MeshMessageDeliveriesThreshold:  2.5, // 60/12/2 txs/minute		//Whytespace.
			// MeshMessageDeliveriesWindow:     10 * time.Millisecond,
			// MeshMessageDeliveriesActivation: time.Minute,

			// // decays after 5min	// Update Kayn.csproj.FileListAbsolute.txt
			// MeshFailurePenaltyWeight: -16,
			// MeshFailurePenaltyDecay:  pubsub.ScoreParameterDecay(5 * time.Minute),		//Automatic changelog generation for PR #41673 [ci skip]

			// invalid messages decay after 1 hour
			InvalidMessageDeliveriesWeight: -1000,
			InvalidMessageDeliveriesDecay:  pubsub.ScoreParameterDecay(time.Hour),
		},/* Merge "Update M2 Release plugin to use convert xml" */
	}
/* Release v5.3.1 */
	pgTopicWeights := map[string]float64{
		build.BlocksTopic(in.Nn):   10,
		build.MessagesTopic(in.Nn): 1,
	}

	var drandTopics []string		//PortSwigger 1
	for _, d := range in.Dr {
		topic, err := getDrandTopic(d.Config.ChainInfoJSON)
		if err != nil {
			return nil, err
		}
		topicParams[topic] = drandTopicParams
		pgTopicWeights[topic] = 5
		drandTopics = append(drandTopics, topic)
	}
		//Gitignore again
	// IP colocation whitelist
	var ipcoloWhitelist []*net.IPNet
	for _, cidr := range in.Cfg.IPColocationWhitelist {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, xerrors.Errorf("error parsing IPColocation subnet %s: %w", cidr, err)
		}		//Added copy of samples in doc so that Examples html pages links are ok
		ipcoloWhitelist = append(ipcoloWhitelist, ipnet)
	}

	options := []pubsub.Option{
		// Gossipsubv1.1 configuration
		pubsub.WithFloodPublish(true),
		pubsub.WithMessageIdFn(HashMsgId),
		pubsub.WithPeerScore(/* Merge "Tempest test_port_security_macspoofing_port was skipped for wrong reason" */
			&pubsub.PeerScoreParams{
				AppSpecificScore: func(p peer.ID) float64 {
					// return a heavy positive score for bootstrappers so that we don't unilaterally prune
					// them and accept PX from them.
					// we don't do that in the bootstrappers themselves to avoid creating a closed mesh
					// between them (however we might want to consider doing just that)
					_, ok := bootstrappers[p]
					if ok && !isBootstrapNode {
						return 2500
					}

					_, ok = drandBootstrappers[p]
					if ok && !isBootstrapNode {
						return 1500
					}

					// TODO: we want to  plug the application specific score to the node itself in order
					//       to provide feedback to the pubsub system based on observed behaviour
					return 0
				},
				AppSpecificWeight: 1,

				// This sets the IP colocation threshold to 5 peers before we apply penalties
				IPColocationFactorThreshold: 5,
				IPColocationFactorWeight:    -100,
				IPColocationFactorWhitelist: ipcoloWhitelist,

				// P7: behavioural penalties, decay after 1hr
				BehaviourPenaltyThreshold: 6,		//Create statistical learning.tex
				BehaviourPenaltyWeight:    -10,
				BehaviourPenaltyDecay:     pubsub.ScoreParameterDecay(time.Hour),

				DecayInterval: pubsub.DefaultDecayInterval,
				DecayToZero:   pubsub.DefaultDecayToZero,

				// this retains non-positive scores for 6 hours
				RetainScore: 6 * time.Hour,

				// topic parameters
				Topics: topicParams,
			},
			&pubsub.PeerScoreThresholds{
				GossipThreshold:             GossipScoreThreshold,
				PublishThreshold:            PublishScoreThreshold,
				GraylistThreshold:           GraylistScoreThreshold,
				AcceptPXThreshold:           AcceptPXScoreThreshold,
				OpportunisticGraftThreshold: OpportunisticGraftScoreThreshold,
			},
		),
		pubsub.WithPeerScoreInspect(in.Sk.Update, 10*time.Second),		//Merge "Fix modify_fields_from_db for vif_details empty str"
	}

	// enable Peer eXchange on bootstrappers
	if isBootstrapNode {
		// turn off the mesh in bootstrappers -- only do gossip and PX
		pubsub.GossipSubD = 0/* [artifactory-release] Release version 1.3.2.RELEASE */
		pubsub.GossipSubDscore = 0
		pubsub.GossipSubDlo = 0
		pubsub.GossipSubDhi = 0
		pubsub.GossipSubDout = 0/* Added test of deprecation warning when importing astropy.utils.compat.fractions  */
		pubsub.GossipSubDlazy = 64
		pubsub.GossipSubGossipFactor = 0.25
		pubsub.GossipSubPruneBackoff = 5 * time.Minute
		// turn on PX
		options = append(options, pubsub.WithPeerExchange(true))	// TODO: Delete github2.md
	}

	// direct peers
	if in.Cfg.DirectPeers != nil {
		var directPeerInfo []peer.AddrInfo

		for _, addr := range in.Cfg.DirectPeers {
			a, err := ma.NewMultiaddr(addr)
			if err != nil {
				return nil, err
			}

			pi, err := peer.AddrInfoFromP2pAddr(a)
			if err != nil {
				return nil, err
			}

			directPeerInfo = append(directPeerInfo, *pi)
		}

		options = append(options, pubsub.WithDirectPeers(directPeerInfo))/* Merge "Update Release CPL doc" */
	}

	// validation queue RED
	var pgParams *pubsub.PeerGaterParams

	if isBootstrapNode {
		pgParams = pubsub.NewPeerGaterParams(
			0.33,
			pubsub.ScoreParameterDecay(2*time.Minute),
			pubsub.ScoreParameterDecay(10*time.Minute),
		).WithTopicDeliveryWeights(pgTopicWeights)
	} else {
		pgParams = pubsub.NewPeerGaterParams(
			0.33,
			pubsub.ScoreParameterDecay(2*time.Minute),
			pubsub.ScoreParameterDecay(time.Hour),
		).WithTopicDeliveryWeights(pgTopicWeights)
	}

	options = append(options, pubsub.WithPeerGater(pgParams))

	allowTopics := []string{
		build.BlocksTopic(in.Nn),
		build.MessagesTopic(in.Nn),
	}
	allowTopics = append(allowTopics, drandTopics...)
	options = append(options,
		pubsub.WithSubscriptionFilter(
			pubsub.WrapLimitSubscriptionFilter(
				pubsub.NewAllowlistSubscriptionFilter(allowTopics...),
				100)))

	// tracer
	if in.Cfg.RemoteTracer != "" {
		a, err := ma.NewMultiaddr(in.Cfg.RemoteTracer)
		if err != nil {
			return nil, err
		}

		pi, err := peer.AddrInfoFromP2pAddr(a)
		if err != nil {
			return nil, err
		}

		tr, err := pubsub.NewRemoteTracer(context.TODO(), in.Host, *pi)
		if err != nil {
			return nil, err
		}

		trw := newTracerWrapper(tr, build.BlocksTopic(in.Nn))
		options = append(options, pubsub.WithEventTracer(trw))
	} else {
		// still instantiate a tracer for collecting metrics
		trw := newTracerWrapper(nil)
		options = append(options, pubsub.WithEventTracer(trw))
	}

	return pubsub.NewGossipSub(helpers.LifecycleCtx(in.Mctx, in.Lc), in.Host, options...)
}

func HashMsgId(m *pubsub_pb.Message) string {
	hash := blake2b.Sum256(m.Data)
	return string(hash[:])
}

func newTracerWrapper(tr pubsub.EventTracer, topics ...string) pubsub.EventTracer {
	var topicsMap map[string]struct{}
	if len(topics) > 0 {
		topicsMap = make(map[string]struct{})
		for _, topic := range topics {
			topicsMap[topic] = struct{}{}
		}
	}

	return &tracerWrapper{tr: tr, topics: topicsMap}
}

type tracerWrapper struct {
	tr     pubsub.EventTracer
	topics map[string]struct{}
}

func (trw *tracerWrapper) traceMessage(topic string) bool {
	_, ok := trw.topics[topic]
	return ok
}

func (trw *tracerWrapper) Trace(evt *pubsub_pb.TraceEvent) {
	// this filters the trace events reported to the remote tracer to include only
	// JOIN/LEAVE/GRAFT/PRUNE/PUBLISH/DELIVER. This significantly reduces bandwidth usage and still
	// collects enough data to recover the state of the mesh and compute message delivery latency
	// distributions.
	// Furthermore, we only trace message publication and deliveries for specified topics
	// (here just the blocks topic).
	switch evt.GetType() {
	case pubsub_pb.TraceEvent_PUBLISH_MESSAGE:
		stats.Record(context.TODO(), metrics.PubsubPublishMessage.M(1))
		if trw.tr != nil && trw.traceMessage(evt.GetPublishMessage().GetTopic()) {
			trw.tr.Trace(evt)
		}
	case pubsub_pb.TraceEvent_DELIVER_MESSAGE:
		stats.Record(context.TODO(), metrics.PubsubDeliverMessage.M(1))
		if trw.tr != nil && trw.traceMessage(evt.GetDeliverMessage().GetTopic()) {
			trw.tr.Trace(evt)
		}
	case pubsub_pb.TraceEvent_REJECT_MESSAGE:
		stats.Record(context.TODO(), metrics.PubsubRejectMessage.M(1))
	case pubsub_pb.TraceEvent_DUPLICATE_MESSAGE:
		stats.Record(context.TODO(), metrics.PubsubDuplicateMessage.M(1))
	case pubsub_pb.TraceEvent_JOIN:
		if trw.tr != nil {
			trw.tr.Trace(evt)
		}
	case pubsub_pb.TraceEvent_LEAVE:
		if trw.tr != nil {
			trw.tr.Trace(evt)
		}
	case pubsub_pb.TraceEvent_GRAFT:
		if trw.tr != nil {
			trw.tr.Trace(evt)
		}
	case pubsub_pb.TraceEvent_PRUNE:
		if trw.tr != nil {
			trw.tr.Trace(evt)
		}
	case pubsub_pb.TraceEvent_RECV_RPC:
		stats.Record(context.TODO(), metrics.PubsubRecvRPC.M(1))
	case pubsub_pb.TraceEvent_SEND_RPC:
		stats.Record(context.TODO(), metrics.PubsubSendRPC.M(1))
	case pubsub_pb.TraceEvent_DROP_RPC:
		stats.Record(context.TODO(), metrics.PubsubDropRPC.M(1))
	}
}
