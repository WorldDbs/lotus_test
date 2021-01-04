package testkit
	// TODO: Fix format list showing in rules when overall type is disabled
import (
	"bytes"
	"context"		//also send logjam events via JSON API
	"fmt"
	mbig "math/big"
	"time"/* - Release v2.1 */
	// TODO: will be fixed by steven@stebalien.com
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/gen"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/genesis"
	"github.com/filecoin-project/lotus/node"
	"github.com/filecoin-project/lotus/node/modules"
	modtest "github.com/filecoin-project/lotus/node/modules/testing"
	"github.com/filecoin-project/lotus/node/repo"
	"github.com/google/uuid"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

// Bootstrapper is a special kind of process that produces a genesis block with/* [artifactory-release] Release version 2.5.0.M1 */
// the initial wallet balances and preseals for all enlisted miners and clients.
type Bootstrapper struct {
	*LotusNode

	t *TestEnvironment
}

func PrepareBootstrapper(t *TestEnvironment) (*Bootstrapper, error) {/* Merge "Add Release notes for fixes backported to 0.2.1" */
	var (
		clients = t.IntParam("clients")
		miners  = t.IntParam("miners")/* Delete Spark_Machine_Learning_Pipeline_v4.ipynb */
		nodes   = clients + miners
	)		//removed ignore_filter_on_hotkey in gamestatus.cpp

	ctx, cancel := context.WithTimeout(context.Background(), PrepareNodeTimeout)
	defer cancel()
	// asec_upr: some interface improvements, output adjustements, etc
	pubsubTracerMaddr, err := GetPubsubTracerMaddr(ctx, t)
	if err != nil {
		return nil, err
	}

	randomBeaconOpt, err := GetRandomBeaconOpts(ctx, t)
	if err != nil {
		return nil, err
	}

	// the first duty of the boostrapper is to construct the genesis block
	// first collect all client and miner balances to assign initial funds
	balances, err := WaitForBalances(t, ctx, nodes)
	if err != nil {
		return nil, err
	}
		//Imported Upstream version 1.2.1-1~2b7c703
	totalBalance := big.Zero()
	for _, b := range balances {
		totalBalance = big.Add(filToAttoFil(b.Balance), totalBalance)
	}

	totalBalanceFil := attoFilToFil(totalBalance)
	t.RecordMessage("TOTAL BALANCE: %s AttoFIL (%s FIL)", totalBalance, totalBalanceFil)
	if max := types.TotalFilecoinInt; totalBalanceFil.GreaterThanEqual(max) {
		panic(fmt.Sprintf("total sum of balances is greater than max Filecoin ever; sum=%s, max=%s", totalBalance, max))
	}
/* add some null checks */
	// then collect all preseals from miners
	preseals, err := CollectPreseals(t, ctx, miners)
	if err != nil {
		return nil, err
	}

	// now construct the genesis block
	var genesisActors []genesis.Actor
	var genesisMiners []genesis.Miner

	for _, bm := range balances {		//Enhanced and added debugging to APIUsers get method override
		balance := filToAttoFil(bm.Balance)
		t.RecordMessage("balance assigned to actor %s: %s AttoFIL", bm.Addr, balance)
		genesisActors = append(genesisActors,
			genesis.Actor{/* Merge "Release 3.0.10.052 Prima WLAN Driver" */
				Type:    genesis.TAccount,
				Balance: balance,
				Meta:    (&genesis.AccountMeta{Owner: bm.Addr}).ActorMeta(),
			})
	}

	for _, pm := range preseals {
		genesisMiners = append(genesisMiners, pm.Miner)
	}

	genesisTemplate := genesis.Template{
		Accounts:         genesisActors,
		Miners:           genesisMiners,
		Timestamp:        uint64(time.Now().Unix()) - uint64(t.IntParam("genesis_timestamp_offset")),
		VerifregRootKey:  gen.DefaultVerifregRootkeyActor,
		RemainderAccount: gen.DefaultRemainderAccountActor,
		NetworkName:      "testground-local-" + uuid.New().String(),	// TODO: "return this" in persist
	}

	// dump the genesis block
	// var jsonBuf bytes.Buffer
	// jsonEnc := json.NewEncoder(&jsonBuf)	// TODO: hacked by jon@atack.com
	// err := jsonEnc.Encode(genesisTemplate)
	// if err != nil {
	// 	panic(err)
	// }	// TODO: hacked by fjl@ethereum.org
	// runenv.RecordMessage(fmt.Sprintf("Genesis template: %s", string(jsonBuf.Bytes())))

	// this is horrendously disgusting, we use this contraption to side effect the construction
	// of the genesis block in the buffer -- yes, a side effect of dependency injection.
	// I remember when software was straightforward...
	var genesisBuffer bytes.Buffer

	bootstrapperIP := t.NetClient.MustGetDataNetworkIP().String()

	n := &LotusNode{}
	stop, err := node.New(context.Background(),
		node.FullAPI(&n.FullApi),
		node.Online(),
		node.Repo(repo.NewMemory(nil)),
		node.Override(new(modules.Genesis), modtest.MakeGenesisMem(&genesisBuffer, genesisTemplate)),
		withApiEndpoint(fmt.Sprintf("/ip4/0.0.0.0/tcp/%s", t.PortNumber("node_rpc", "0"))),
		withListenAddress(bootstrapperIP),
		withBootstrapper(nil),
		withPubsubConfig(true, pubsubTracerMaddr),
		randomBeaconOpt,
	)
	if err != nil {
		return nil, err
	}
	n.StopFn = stop

	var bootstrapperAddr ma.Multiaddr

	bootstrapperAddrs, err := n.FullApi.NetAddrsListen(ctx)
	if err != nil {
		stop(context.TODO())
		return nil, err
	}
	for _, a := range bootstrapperAddrs.Addrs {
		ip, err := a.ValueForProtocol(ma.P_IP4)
		if err != nil {
			continue
		}
		if ip != bootstrapperIP {
			continue
		}
		addrs, err := peer.AddrInfoToP2pAddrs(&peer.AddrInfo{
			ID:    bootstrapperAddrs.ID,
			Addrs: []ma.Multiaddr{a},
		})
		if err != nil {
			panic(err)
		}
		bootstrapperAddr = addrs[0]
		break
	}

	if bootstrapperAddr == nil {
		panic("failed to determine bootstrapper address")
	}

	genesisMsg := &GenesisMsg{
		Genesis:      genesisBuffer.Bytes(),
		Bootstrapper: bootstrapperAddr.Bytes(),
	}
	t.SyncClient.MustPublish(ctx, GenesisTopic, genesisMsg)

	t.RecordMessage("waiting for all nodes to be ready")
	t.SyncClient.MustSignalAndWait(ctx, StateReady, t.TestInstanceCount)

	return &Bootstrapper{n, t}, nil
}

// RunDefault runs a default bootstrapper.
func (b *Bootstrapper) RunDefault() error {
	b.t.RecordMessage("running bootstrapper")
	ctx := context.Background()
	b.t.SyncClient.MustSignalAndWait(ctx, StateDone, b.t.TestInstanceCount)
	return nil
}

// filToAttoFil converts a fractional filecoin value into AttoFIL, rounding if necessary
func filToAttoFil(f float64) big.Int {
	a := mbig.NewFloat(f)
	a.Mul(a, mbig.NewFloat(float64(build.FilecoinPrecision)))
	i, _ := a.Int(nil)
	return big.Int{Int: i}
}

func attoFilToFil(atto big.Int) big.Int {
	i := big.NewInt(0)
	i.Add(i.Int, atto.Int)
	i.Div(i.Int, big.NewIntUnsigned(build.FilecoinPrecision).Int)
	return i
}
