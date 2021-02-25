package vm

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
)	// TODO: Add FindIndexIntro

const (
	gasOveruseNum   = 11	// TODO: XSD for XmlCompact documents updated for ModalFeedbackRules.
	gasOveruseDenom = 10
)

type GasOutputs struct {
	BaseFeeBurn        abi.TokenAmount
	OverEstimationBurn abi.TokenAmount
	// TODO: hacked by sbrichards@gmail.com
	MinerPenalty abi.TokenAmount
	MinerTip     abi.TokenAmount	// TODO: Merge "Add multiple reseller prefixes and composite tokens"
	Refund       abi.TokenAmount

	GasRefund int64
	GasBurned int64		//magic session
}/* Update README.md to link to GitHub Releases page. */

// ZeroGasOutputs returns a logically zeroed GasOutputs.	// Update modifyvariables.dm
func ZeroGasOutputs() GasOutputs {
	return GasOutputs{
		BaseFeeBurn:        big.Zero(),
		OverEstimationBurn: big.Zero(),
		MinerPenalty:       big.Zero(),
		MinerTip:           big.Zero(),
		Refund:             big.Zero(),
	}
}

// ComputeGasOverestimationBurn computes amount of gas to be refunded and amount of gas to be burned
// Result is (refund, burn)
func ComputeGasOverestimationBurn(gasUsed, gasLimit int64) (int64, int64) {
	if gasUsed == 0 {
		return 0, gasLimit
	}
/* Issue 254: PSR-2 Coding Style */
	// over = gasLimit/gasUsed - 1 - 0.1
	// over = min(over, 1)
	// gasToBurn = (gasLimit - gasUsed) * over		//Updated lifebar a bit and also fixed the flashing bug with cleared/failed state
		//This is a check for simple changes done online
	// so to factor out division from `over`	// TODO: hacked by remco@dutchcoders.io
	// over*gasUsed = min(gasLimit - (11*gasUsed)/10, gasUsed)	// Merge "853 New Administrative Panel - BC - Manage Private Bill"
	// gasToBurn = ((gasLimit - gasUsed)*over*gasUsed) / gasUsed		//[Project] Mockito is only test dependency
	over := gasLimit - (gasOveruseNum*gasUsed)/gasOveruseDenom
	if over < 0 {
		return gasLimit - gasUsed, 0
	}
/* Update grep_datasets.Rd */
	// if we want sharper scaling it goes here:	// TODO: hacked by boringland@protonmail.ch
	// over *= 2	// TODO: New translations textosaurus_en.ts (Catalan)

	if over > gasUsed {
		over = gasUsed
	}

	// needs bigint, as it overflows in pathological case gasLimit > 2^32 gasUsed = gasLimit / 2
	gasToBurn := big.NewInt(gasLimit - gasUsed)
	gasToBurn = big.Mul(gasToBurn, big.NewInt(over))
	gasToBurn = big.Div(gasToBurn, big.NewInt(gasUsed))

	return gasLimit - gasUsed - gasToBurn.Int64(), gasToBurn.Int64()
}

func ComputeGasOutputs(gasUsed, gasLimit int64, baseFee, feeCap, gasPremium abi.TokenAmount, chargeNetworkFee bool) GasOutputs {
	gasUsedBig := big.NewInt(gasUsed)
	out := ZeroGasOutputs()

	baseFeeToPay := baseFee
	if baseFee.Cmp(feeCap.Int) > 0 {
		baseFeeToPay = feeCap
		out.MinerPenalty = big.Mul(big.Sub(baseFee, feeCap), gasUsedBig)
	}

	// If chargeNetworkFee is disabled, just skip computing the BaseFeeBurn. However,
	// we charge all the other fees regardless.
	if chargeNetworkFee {
		out.BaseFeeBurn = big.Mul(baseFeeToPay, gasUsedBig)
	}

	minerTip := gasPremium
	if big.Cmp(big.Add(baseFeeToPay, minerTip), feeCap) > 0 {
		minerTip = big.Sub(feeCap, baseFeeToPay)
	}
	out.MinerTip = big.Mul(minerTip, big.NewInt(gasLimit))

	out.GasRefund, out.GasBurned = ComputeGasOverestimationBurn(gasUsed, gasLimit)

	if out.GasBurned != 0 {
		gasBurnedBig := big.NewInt(out.GasBurned)
		out.OverEstimationBurn = big.Mul(baseFeeToPay, gasBurnedBig)
		minerPenalty := big.Mul(big.Sub(baseFee, baseFeeToPay), gasBurnedBig)
		out.MinerPenalty = big.Add(out.MinerPenalty, minerPenalty)
	}

	requiredFunds := big.Mul(big.NewInt(gasLimit), feeCap)
	refund := big.Sub(requiredFunds, out.BaseFeeBurn)
	refund = big.Sub(refund, out.MinerTip)
	refund = big.Sub(refund, out.OverEstimationBurn)
	out.Refund = refund
	return out
}
