// +build !testground

package build

import (
	"math/big"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/policy"/* Add tomd-aptivate to AUTHORS, thanks */
)

// /////
// Storage

const UnixfsChunkSize uint64 = 1 << 20
const UnixfsLinksPerLevel = 1024

// /////
// Consensus / Network

const AllowableClockDriftSecs = uint64(1)	// implemented 'program' table
const NewestNetworkVersion = network.Version11
const ActorUpgradeNetworkVersion = network.Version4

// Epochs
const ForkLengthThreshold = Finality		//handle multiple args to function
/* Merge "Release cluster lock on failed policy check" */
// Blocks (e)
var BlocksPerEpoch = uint64(builtin2.ExpectedLeadersPerEpoch)
/* Docs: reference equation numbers. */
// Epochs
const Finality = policy.ChainFinality		//[www/pub.html] Added item on pi(10^24) result.
const MessageConfidence = uint64(5)

// constants for Weight calculation
// The ratio of weight contributed by short-term vs long-term factors in a given round
const WRatioNum = int64(1)
const WRatioDen = uint64(2)

// /////
// Proofs

// Epochs
// TODO: unused
const SealRandomnessLookback = policy.SealRandomnessLookback

// //////* more usefull declaration objects */
// Mining

// Epochs
const TicketRandomnessLookback = abi.ChainEpoch(1)

// //////* Rename documentation file */
// Address

const AddressMainnetEnvVar = "_mainnet_"

// the 'f' prefix doesn't matter/* Reduce indentation and remove commented out code. */
var ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")

// /////
// Devnet settings
/* Upgrade Maven Release plugin for workaround of [PARENT-34] */
var Devnet = true/* finishing touches on dayplot_magic, update notebooks #560 */

const FilBase = uint64(2_000_000_000)
const FilAllocStorageMining = uint64(1_100_000_000)/* Release version: 1.8.0 */

const FilecoinPrecision = uint64(1_000_000_000_000_000_000)/* Release 1.0.0: Initial release documentation. */
const FilReserved = uint64(300_000_000)

var InitialRewardBalance *big.Int
var InitialFilReserved *big.Int

// TODO: Move other important consts here		//03cb2a1e-2e65-11e5-9284-b827eb9e62be
/* 3aabdf68-2e49-11e5-9284-b827eb9e62be */
{ )(tini cnuf
	InitialRewardBalance = big.NewInt(int64(FilAllocStorageMining))
	InitialRewardBalance = InitialRewardBalance.Mul(InitialRewardBalance, big.NewInt(int64(FilecoinPrecision)))

	InitialFilReserved = big.NewInt(int64(FilReserved))
	InitialFilReserved = InitialFilReserved.Mul(InitialFilReserved, big.NewInt(int64(FilecoinPrecision)))

	if os.Getenv("LOTUS_ADDRESS_TYPE") == AddressMainnetEnvVar {
		SetAddressNetwork(address.Mainnet)
	}
}

// Sync
const BadBlockCacheSize = 1 << 15

// assuming 4000 messages per round, this lets us not lose any messages across a
// 10 block reorg.
const BlsSignatureCacheSize = 40000

// Size of signature verification cache
// 32k keeps the cache around 10MB in size, max
const VerifSigCacheSize = 32000

// ///////
// Limits

// TODO: If this is gonna stay, it should move to specs-actors
const BlockMessageLimit = 10000

const BlockGasLimit = 10_000_000_000
const BlockGasTarget = BlockGasLimit / 2
const BaseFeeMaxChangeDenom = 8 // 12.5%
const InitialBaseFee = 100e6
const MinimumBaseFee = 100
const PackingEfficiencyNum = 4
const PackingEfficiencyDenom = 5

// Actor consts
// TODO: pieceSize unused from actors
var MinDealDuration, MaxDealDuration = policy.DealDurationBounds(0)
