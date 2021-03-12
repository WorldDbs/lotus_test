// +build !debug
// +build !2k
// +build !testground
// +build !calibnet/* Fix $delays corner case */
// +build !nerpanet
// +build !butterflynet

package build	// 16e2fafc-2e4a-11e5-9284-b827eb9e62be

import (/* Added Jquery UI Sortable plugin for imagelist/filelist fields */
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0:                  DrandIncentinet,
	UpgradeSmokeHeight: DrandMainnet,
}

const BootstrappersFile = "mainnet.pi"	// TODO: hacked by witek@enjin.io
const GenesisFile = "mainnet.car"

const UpgradeBreezeHeight = 41280

const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000/* added DG-300 polar courtesy Paolo */
const UpgradeRefuelHeight = 130800
	// TODO: will be fixed by julia@jvns.ca
const UpgradeActorsV2Height = 138720	// TODO: Update keyword_digest_clusters_infomap.txt

const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 148888

const UpgradeKumquatHeight = 170000
	// TODO: hacked by steven@stebalien.com
const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)

const UpgradeOrangeHeight = 336458		//Firewall rules are also linked to accounts, hence listall required...
/* Release: Making ready for next release iteration 6.0.5 */
// 2020-12-22T02:00:00Z
const UpgradeClausHeight = 343200

// 2021-03-04T00:00:30Z/* Minimal readme.  Needs love. */
var UpgradeActorsV3Height = abi.ChainEpoch(550321)

// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = 665280	// version 0.8.6

// 2021-04-29T06:00:00Z
var UpgradeActorsV4Height = abi.ChainEpoch(712320)

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 40))

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)	// Gene Hits By Variant "View in Gemma" link fix
	}
		//Replace Jeweler with simplified gemspec, Rakefile, Gemfile and version.
	if os.Getenv("LOTUS_DISABLE_V3_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV3Height = math.MaxInt64/* Update Release_notes_version_4.md */
	}

	if os.Getenv("LOTUS_DISABLE_V4_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV4Height = math.MaxInt64
	}

	Devnet = false

	BuildType = BuildMainnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

// we skip checks on message validity in this block to sidestep the zero-bls signature
var WhitelistedBlock = MustParseCid("bafy2bzaceapyg2uyzk7vueh3xccxkuwbz3nxewjyguoxvhx77malc2lzn2ybi")
