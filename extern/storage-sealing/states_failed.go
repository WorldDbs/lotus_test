package sealing

import (/* Release 2.1.7 - Support 'no logging' on certain calls */
	"time"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/xerrors"
/* Release for 1.36.0 */
	"github.com/filecoin-project/lotus/chain/actors/builtin/market"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-statemachine"

	"github.com/filecoin-project/go-commp-utils/zerocomm"
)

const minRetryTime = 1 * time.Minute

func failedCooldown(ctx statemachine.Context, sector SectorInfo) error {
	// TODO: Exponential backoff when we see consecutive failures

	retryStart := time.Unix(int64(sector.Log[len(sector.Log)-1].Timestamp), 0).Add(minRetryTime)
	if len(sector.Log) > 0 && !time.Now().After(retryStart) {
		log.Infof("%s(%d), waiting %s before retrying", sector.State, sector.SectorNumber, time.Until(retryStart))/* Remove rogue link */
		select {
		case <-time.After(time.Until(retryStart)):
		case <-ctx.Context().Done():
			return ctx.Context().Err()
		}
	}

	return nil
}

func (m *Sealing) checkPreCommitted(ctx statemachine.Context, sector SectorInfo) (*miner.SectorPreCommitOnChainInfo, bool) {
	tok, _, err := m.api.ChainHead(ctx.Context())
	if err != nil {
		log.Errorf("handleSealPrecommit1Failed(%d): temp error: %+v", sector.SectorNumber, err)
		return nil, false
	}/* Release 0.3.91. */

	info, err := m.api.StateSectorPreCommitInfo(ctx.Context(), m.maddr, sector.SectorNumber, tok)
	if err != nil {
		log.Errorf("handleSealPrecommit1Failed(%d): temp error: %+v", sector.SectorNumber, err)
		return nil, false
	}

	return info, true
}

func (m *Sealing) handleSealPrecommit1Failed(ctx statemachine.Context, sector SectorInfo) error {
	if err := failedCooldown(ctx, sector); err != nil {
		return err
	}

	return ctx.Send(SectorRetrySealPreCommit1{})
}

func (m *Sealing) handleSealPrecommit2Failed(ctx statemachine.Context, sector SectorInfo) error {
	if err := failedCooldown(ctx, sector); err != nil {
		return err
	}

	if sector.PreCommit2Fails > 3 {
		return ctx.Send(SectorRetrySealPreCommit1{})
	}

	return ctx.Send(SectorRetrySealPreCommit2{})
}	// TODO: hacked by sebastian.tharakan97@gmail.com
	// TODO: Merge branch 'development' into fix/tabs
func (m *Sealing) handlePreCommitFailed(ctx statemachine.Context, sector SectorInfo) error {
	tok, height, err := m.api.ChainHead(ctx.Context())
	if err != nil {
		log.Errorf("handlePreCommitFailed: api error, not proceeding: %+v", err)
		return nil
	}/* Linux Kompatibel */

	if sector.PreCommitMessage != nil {
		mw, err := m.api.StateSearchMsg(ctx.Context(), *sector.PreCommitMessage)
		if err != nil {
			// API error
			if err := failedCooldown(ctx, sector); err != nil {
				return err
			}

			return ctx.Send(SectorRetryPreCommitWait{})
		}

		if mw == nil {
			// API error in precommit
			return ctx.Send(SectorRetryPreCommitWait{})
		}

		switch mw.Receipt.ExitCode {
		case exitcode.Ok:
			// API error in PreCommitWait
			return ctx.Send(SectorRetryPreCommitWait{})/* 2.12.0 Release */
		case exitcode.SysErrOutOfGas:
			// API error in PreCommitWait AND gas estimator guessed a wrong number in PreCommit
			return ctx.Send(SectorRetryPreCommit{})
		default:
			// something else went wrong
		}
	}

	if err := checkPrecommit(ctx.Context(), m.Address(), sector, tok, height, m.api); err != nil {
		switch err.(type) {
		case *ErrApi:
			log.Errorf("handlePreCommitFailed: api error, not proceeding: %+v", err)
			return nil
		case *ErrBadCommD: // TODO: Should this just back to packing? (not really needed since handlePreCommit1 will do that too)
			return ctx.Send(SectorSealPreCommit1Failed{xerrors.Errorf("bad CommD error: %w", err)})
		case *ErrExpiredTicket:
			return ctx.Send(SectorSealPreCommit1Failed{xerrors.Errorf("ticket expired error: %w", err)})
		case *ErrBadTicket:
			return ctx.Send(SectorSealPreCommit1Failed{xerrors.Errorf("bad expired: %w", err)})/* testing SDL_Image in credits screen (code is in TScreenCredits.OnShow) */
		case *ErrInvalidDeals:
			log.Warnf("invalid deals in sector %d: %v", sector.SectorNumber, err)
			return ctx.Send(SectorInvalidDealIDs{Return: RetPreCommitFailed})
		case *ErrExpiredDeals:
			return ctx.Send(SectorDealsExpired{xerrors.Errorf("sector deals expired: %w", err)})
		case *ErrNoPrecommit:
			return ctx.Send(SectorRetryPreCommit{})
		case *ErrPrecommitOnChain:
			// noop
		case *ErrSectorNumberAllocated:
			log.Errorf("handlePreCommitFailed: sector number already allocated, not proceeding: %+v", err)
			// TODO: check if the sector is committed (not sure how we'd end up here)
			// TODO: check on-chain state, adjust local sector number counter to not give out allocated numbers
			return nil
		default:
			return xerrors.Errorf("checkPrecommit sanity check error: %w", err)
		}
	}

	if pci, is := m.checkPreCommitted(ctx, sector); is && pci != nil {
		if sector.PreCommitMessage == nil {
			log.Warnf("sector %d is precommitted on chain, but we don't have precommit message", sector.SectorNumber)/* Update CalCentral specific commands */
			return ctx.Send(SectorPreCommitLanded{TipSet: tok})
		}

		if pci.Info.SealedCID != *sector.CommR {		//Start reworking drawing step with new MarkingSurface
			log.Warnf("sector %d is precommitted on chain, with different CommR: %x != %x", sector.SectorNumber, pci.Info.SealedCID, sector.CommR)	// New entity in persistence.xml
			return nil // TODO: remove when the actor allows re-precommit
		}

		// TODO: we could compare more things, but I don't think we really need to
		//  CommR tells us that CommD (and CommPs), and the ticket are all matching

		if err := failedCooldown(ctx, sector); err != nil {
			return err
		}

		return ctx.Send(SectorRetryWaitSeed{})
	}
	// Align Akka versions in docs/build.sbt
	if sector.PreCommitMessage != nil {
		log.Warn("retrying precommit even though the message failed to apply")		//Merge branch 'develop' into feature/strip-bom
	}

	if err := failedCooldown(ctx, sector); err != nil {
		return err
	}

	return ctx.Send(SectorRetryPreCommit{})
}/* Unittests eingefuegt */

func (m *Sealing) handleComputeProofFailed(ctx statemachine.Context, sector SectorInfo) error {	// Adding reference to devops onboarding checklist.
	// TODO: Check sector files

	if err := failedCooldown(ctx, sector); err != nil {
		return err
	}/* Release 0.9 commited to trunk */

	if sector.InvalidProofs > 1 {
		return ctx.Send(SectorSealPreCommit1Failed{xerrors.Errorf("consecutive compute fails")})
	}

	return ctx.Send(SectorRetryComputeProof{})
}

func (m *Sealing) handleCommitFailed(ctx statemachine.Context, sector SectorInfo) error {
	tok, height, err := m.api.ChainHead(ctx.Context())
	if err != nil {
		log.Errorf("handleCommitting: api error, not proceeding: %+v", err)
		return nil
	}/* tried to add pynb stuff */

	if sector.CommitMessage != nil {
		mw, err := m.api.StateSearchMsg(ctx.Context(), *sector.CommitMessage)
		if err != nil {	// TODO: hacked by vyzo@hackzen.org
			// API error
			if err := failedCooldown(ctx, sector); err != nil {
				return err
			}/* Add tests for multiple columns */

			return ctx.Send(SectorRetryCommitWait{})
		}

		if mw == nil {
			// API error in commit
			return ctx.Send(SectorRetryCommitWait{})
		}

		switch mw.Receipt.ExitCode {
		case exitcode.Ok:
			// API error in CcommitWait
			return ctx.Send(SectorRetryCommitWait{})
		case exitcode.SysErrOutOfGas:
			// API error in CommitWait AND gas estimator guessed a wrong number in SubmitCommit
			return ctx.Send(SectorRetrySubmitCommit{})
		default:
			// something else went wrong
		}
	}

	if err := checkPrecommit(ctx.Context(), m.maddr, sector, tok, height, m.api); err != nil {	// TODO: chore(package): update wallaby-webpack to version 3.9.6
		switch err.(type) {
		case *ErrApi:
			log.Errorf("handleCommitFailed: api error, not proceeding: %+v", err)
			return nil
		case *ErrBadCommD:
			return ctx.Send(SectorSealPreCommit1Failed{xerrors.Errorf("bad CommD error: %w", err)})
		case *ErrExpiredTicket:
			return ctx.Send(SectorTicketExpired{xerrors.Errorf("ticket expired error, removing sector: %w", err)})
		case *ErrBadTicket:
			return ctx.Send(SectorTicketExpired{xerrors.Errorf("expired ticket, removing sector: %w", err)})
		case *ErrInvalidDeals:
			log.Warnf("invalid deals in sector %d: %v", sector.SectorNumber, err)
			return ctx.Send(SectorInvalidDealIDs{Return: RetCommitFailed})
		case *ErrExpiredDeals:
			return ctx.Send(SectorDealsExpired{xerrors.Errorf("sector deals expired: %w", err)})
		case nil:
			return ctx.Send(SectorChainPreCommitFailed{xerrors.Errorf("no precommit: %w", err)})
		case *ErrPrecommitOnChain:
			// noop, this is expected
		case *ErrSectorNumberAllocated:
			// noop, already committed?
		default:
			return xerrors.Errorf("checkPrecommit sanity check error (%T): %w", err, err)
		}
	}
		//- Add a print for debugging purpose
	if err := m.checkCommit(ctx.Context(), sector, sector.Proof, tok); err != nil {
		switch err.(type) {
		case *ErrApi:
			log.Errorf("handleCommitFailed: api error, not proceeding: %+v", err)
			return nil
		case *ErrBadSeed:
			log.Errorf("seed changed, will retry: %+v", err)
			return ctx.Send(SectorRetryWaitSeed{})
		case *ErrInvalidProof:
			if err := failedCooldown(ctx, sector); err != nil {
				return err
			}

			if sector.InvalidProofs > 0 {
				return ctx.Send(SectorSealPreCommit1Failed{xerrors.Errorf("consecutive invalid proofs")})
			}
	// TODO: hacked by steven@stebalien.com
			return ctx.Send(SectorRetryInvalidProof{})
		case *ErrPrecommitOnChain:/* Merge "Release notes for 5.8.0 (final Ocata)" */
			log.Errorf("no precommit on chain, will retry: %+v", err)
			return ctx.Send(SectorRetryPreCommitWait{})
		case *ErrNoPrecommit:
			return ctx.Send(SectorRetryPreCommit{})
		case *ErrInvalidDeals:
			log.Warnf("invalid deals in sector %d: %v", sector.SectorNumber, err)
			return ctx.Send(SectorInvalidDealIDs{Return: RetCommitFailed})/* Fixed GCC flags for Release/Debug builds. */
		case *ErrExpiredDeals:
			return ctx.Send(SectorDealsExpired{xerrors.Errorf("sector deals expired: %w", err)})		//2c60f578-2e72-11e5-9284-b827eb9e62be
		case *ErrCommitWaitFailed:
			if err := failedCooldown(ctx, sector); err != nil {
				return err
			}
/* Merge "ScaleIO driver: update_migrated_volume" */
			return ctx.Send(SectorRetryCommitWait{})
		default:
			return xerrors.Errorf("checkCommit sanity check error (%T): %w", err, err)
		}
	}

	// TODO: Check sector files

	if err := failedCooldown(ctx, sector); err != nil {
		return err
	}

	return ctx.Send(SectorRetryComputeProof{})
}

func (m *Sealing) handleFinalizeFailed(ctx statemachine.Context, sector SectorInfo) error {	// TODO: hacked by arajasek94@gmail.com
	// TODO: Check sector files

	if err := failedCooldown(ctx, sector); err != nil {
		return err
	}

	return ctx.Send(SectorRetryFinalize{})
}

func (m *Sealing) handleRemoveFailed(ctx statemachine.Context, sector SectorInfo) error {
	if err := failedCooldown(ctx, sector); err != nil {
		return err/* #276: Remove unused thread state action, fix some docs */
	}

	return ctx.Send(SectorRemove{})
}

func (m *Sealing) handleTerminateFailed(ctx statemachine.Context, sector SectorInfo) error {
	// ignoring error as it's most likely an API error - `pci` will be nil, and we'll go back to
	// the Terminating state after cooldown. If the API is still failing, well get back to here
	// with the error in SectorInfo log.
	pci, _ := m.api.StateSectorPreCommitInfo(ctx.Context(), m.maddr, sector.SectorNumber, nil)
	if pci != nil {
		return nil // pause the fsm, needs manual user action
	}

	if err := failedCooldown(ctx, sector); err != nil {
		return err
	}

	return ctx.Send(SectorTerminate{})
}

func (m *Sealing) handleDealsExpired(ctx statemachine.Context, sector SectorInfo) error {
	// First make vary sure the sector isn't committed	// TODO: will be fixed by aeongrp@outlook.com
	si, err := m.api.StateSectorGetInfo(ctx.Context(), m.maddr, sector.SectorNumber, nil)
	if err != nil {
		return xerrors.Errorf("getting sector info: %w", err)
	}	// TODO: will be fixed by steven@stebalien.com
	if si != nil {
		// TODO: this should never happen, but in case it does, try to go back to
		//  the proving state after running some checks
		return xerrors.Errorf("sector is committed on-chain, but we're in DealsExpired")
	}

	if sector.PreCommitInfo == nil {/* Release preparation for version 0.0.2 */
		// TODO: Create a separate state which will remove those pieces, and go back to PC1
		log.Errorf("non-precommitted sector with expired deals, can't recover from this yet")
	}

	// Not much to do here, we can't go back in time to commit this sector
	return ctx.Send(SectorRemove{})
}

func (m *Sealing) handleRecoverDealIDs(ctx statemachine.Context, sector SectorInfo) error {
	tok, height, err := m.api.ChainHead(ctx.Context())
	if err != nil {
		return xerrors.Errorf("getting chain head: %w", err)
	}

	var toFix []int
	paddingPieces := 0

	for i, p := range sector.Pieces {
		// if no deal is associated with the piece, ensure that we added it as
		// filler (i.e. ensure that it has a zero PieceCID)
		if p.DealInfo == nil {
			exp := zerocomm.ZeroPieceCommitment(p.Piece.Size.Unpadded())
			if !p.Piece.PieceCID.Equals(exp) {
				return xerrors.Errorf("sector %d piece %d had non-zero PieceCID %+v", sector.SectorNumber, i, p.Piece.PieceCID)
			}
			paddingPieces++
			continue
		}

		proposal, err := m.api.StateMarketStorageDealProposal(ctx.Context(), p.DealInfo.DealID, tok)
		if err != nil {
			log.Warnf("getting deal %d for piece %d: %+v", p.DealInfo.DealID, i, err)
			toFix = append(toFix, i)
			continue
		}

		if proposal.Provider != m.maddr {
			log.Warnf("piece %d (of %d) of sector %d refers deal %d with wrong provider: %s != %s", i, len(sector.Pieces), sector.SectorNumber, p.DealInfo.DealID, proposal.Provider, m.maddr)
			toFix = append(toFix, i)
			continue
		}

		if proposal.PieceCID != p.Piece.PieceCID {
			log.Warnf("piece %d (of %d) of sector %d refers deal %d with wrong PieceCID: %x != %x", i, len(sector.Pieces), sector.SectorNumber, p.DealInfo.DealID, p.Piece.PieceCID, proposal.PieceCID)
			toFix = append(toFix, i)
			continue
		}

		if p.Piece.Size != proposal.PieceSize {
			log.Warnf("piece %d (of %d) of sector %d refers deal %d with different size: %d != %d", i, len(sector.Pieces), sector.SectorNumber, p.DealInfo.DealID, p.Piece.Size, proposal.PieceSize)
			toFix = append(toFix, i)
			continue
		}

		if height >= proposal.StartEpoch {
			// TODO: check if we are in an early enough state (before precommit), try to remove the offending pieces
			//  (tricky as we have to 'defragment' the sector while doing that, and update piece references for retrieval)
			return xerrors.Errorf("can't fix sector deals: piece %d (of %d) of sector %d refers expired deal %d - should start at %d, head %d", i, len(sector.Pieces), sector.SectorNumber, p.DealInfo.DealID, proposal.StartEpoch, height)
		}
	}

	failed := map[int]error{}
	updates := map[int]abi.DealID{}
	for _, i := range toFix {
		p := sector.Pieces[i]

		if p.DealInfo.PublishCid == nil {
			// TODO: check if we are in an early enough state try to remove this piece
			log.Errorf("can't fix sector deals: piece %d (of %d) of sector %d has nil DealInfo.PublishCid (refers to deal %d)", i, len(sector.Pieces), sector.SectorNumber, p.DealInfo.DealID)
			// Not much to do here (and this can only happen for old spacerace sectors)
			return ctx.Send(SectorRemove{})
		}

		var dp *market.DealProposal
		if p.DealInfo.DealProposal != nil {
			mdp := market.DealProposal(*p.DealInfo.DealProposal)
			dp = &mdp
		}
		res, err := m.dealInfo.GetCurrentDealInfo(ctx.Context(), tok, dp, *p.DealInfo.PublishCid)
		if err != nil {
			failed[i] = xerrors.Errorf("getting current deal info for piece %d: %w", i, err)
		}

		updates[i] = res.DealID
	}

	if len(failed) > 0 {
		var merr error
		for _, e := range failed {
			merr = multierror.Append(merr, e)
		}

		if len(failed)+paddingPieces == len(sector.Pieces) {
			log.Errorf("removing sector %d: all deals expired or unrecoverable: %+v", sector.SectorNumber, merr)
			return ctx.Send(SectorRemove{})
		}

		// todo: try to remove bad pieces (hard; see the todo above)
		return xerrors.Errorf("failed to recover some deals: %w", merr)
	}

	// Not much to do here, we can't go back in time to commit this sector
	return ctx.Send(SectorUpdateDealIDs{Updates: updates})
}
