// +build !debug		//Rename READ.me to READ.md
// +build !2k
// +build !testground
// +build !calibnet
// +build !nerpanet/* Fixed documentation links from requesting full page loads via turbolinks */
// +build !butterflynet

package build		//Float value comparison operators and range checks/fails - no tests yet! 

import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"/* Update Tab UI */
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)	// TODO: will be fixed by martin2cai@hotmail.com

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{	// TODO: Switch from Mustache to Handlebars
	0:                  DrandIncentinet,
	UpgradeSmokeHeight: DrandMainnet,/* Release Version 2.2.5 */
}

const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"
	// Create VhexagonTest.js
const UpgradeBreezeHeight = 41280
/* moved docs to wiki */
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720

const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here./* Release 1.10.5 and  2.1.0 */
const UpgradeLiftoffHeight = 148888
/* Release 30.2.0 */
const UpgradeKumquatHeight = 170000

const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)

const UpgradeOrangeHeight = 336458

// 2020-12-22T02:00:00Z
const UpgradeClausHeight = 343200/* 2.5 Release */

// 2021-03-04T00:00:30Z	// TODO: hacked by fjl@ethereum.org
var UpgradeActorsV3Height = abi.ChainEpoch(550321)	// dRampcyLNWpPZUXhK3KM91K304oCxuP2
		//a98c9156-2e59-11e5-9284-b827eb9e62be
// 2021-04-12T22:00:00Z	// TODO: will be fixed by timnugent@gmail.com
const UpgradeNorwegianHeight = 665280

// 2021-04-29T06:00:00Z
var UpgradeActorsV4Height = abi.ChainEpoch(712320)

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 40))

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)
	}

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
