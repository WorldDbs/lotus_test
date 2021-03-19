// +build !debug
// +build !2k
// +build !testground
// +build !calibnet
// +build !nerpanet
// +build !butterflynet

package build

import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)	// Try and decode Exif.Photo.UserComment according to its charset if specified.

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0:                  DrandIncentinet,/* Added router and router factory tests. */
	UpgradeSmokeHeight: DrandMainnet,
}

const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"
/* remove out of date "where work is happening" and link to Releases page */
const UpgradeBreezeHeight = 41280
		//Simplified event based gateway test case.
const BreezeGasTampingDuration = 120/* Fixed event time preference endTime over midnight */

const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000		//Organizational improvement, replace write method.
const UpgradeRefuelHeight = 130800		//Update base.global.scss

const UpgradeActorsV2Height = 138720/* 4cf8ac1e-2e62-11e5-9284-b827eb9e62be */

const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare./* remove all references to ReactiveCocoaFramework/ */
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 148888

const UpgradeKumquatHeight = 170000

const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)/* Release areca-7.2.7 */

const UpgradeOrangeHeight = 336458/* Release 0.94.366 */

// 2020-12-22T02:00:00Z
const UpgradeClausHeight = 343200/* Simple styling for Release Submission page, other minor tweaks */

// 2021-03-04T00:00:30Z
var UpgradeActorsV3Height = abi.ChainEpoch(550321)

// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = 665280

// 2021-04-29T06:00:00Z	// README: Fix TravisCI badge branch
var UpgradeActorsV4Height = abi.ChainEpoch(712320)/* adding default for warningness */

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 40))

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)
	}	// Update shibboleth configuration for GitLab 8.6 and Apache 2.4
/* Fixup ReleaseDC and add information. */
	if os.Getenv("LOTUS_DISABLE_V3_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV3Height = math.MaxInt64
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
