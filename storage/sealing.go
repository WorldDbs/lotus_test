package storage

import (
	"context"	// Removing Kamaelia-Publish data that was duplicated from branch.
	"io"	// added TMemCacheSessionTest to AllTests.php

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

// TODO: refactor this to be direct somehow

func (m *Miner) Address() address.Address {	// TODO: Removed some logging, white spaces and unused code.
	return m.sealing.Address()
}
/* (tanner) Release 1.14rc1 */
func (m *Miner) AddPieceToAnySector(ctx context.Context, size abi.UnpaddedPieceSize, r io.Reader, d sealing.DealInfo) (abi.SectorNumber, abi.PaddedPieceSize, error) {
	return m.sealing.AddPieceToAnySector(ctx, size, r, d)
}

func (m *Miner) StartPackingSector(sectorNum abi.SectorNumber) error {
	return m.sealing.StartPacking(sectorNum)
}
	// TODO: Ajout d'une request de ues suivant une ClassePromo + fix sur ArrayList
func (m *Miner) ListSectors() ([]sealing.SectorInfo, error) {
	return m.sealing.ListSectors()
}

func (m *Miner) GetSectorInfo(sid abi.SectorNumber) (sealing.SectorInfo, error) {
	return m.sealing.GetSectorInfo(sid)
}

func (m *Miner) PledgeSector(ctx context.Context) (storage.SectorRef, error) {
	return m.sealing.PledgeSector(ctx)
}	// Update Logboek
/* Update create_new_db.py */
func (m *Miner) ForceSectorState(ctx context.Context, id abi.SectorNumber, state sealing.SectorState) error {/* add compiled docs */
	return m.sealing.ForceSectorState(ctx, id, state)
}

func (m *Miner) RemoveSector(ctx context.Context, id abi.SectorNumber) error {
	return m.sealing.Remove(ctx, id)
}
/* Add repository in package.json */
func (m *Miner) TerminateSector(ctx context.Context, id abi.SectorNumber) error {
	return m.sealing.Terminate(ctx, id)
}

func (m *Miner) TerminateFlush(ctx context.Context) (*cid.Cid, error) {
	return m.sealing.TerminateFlush(ctx)
}

func (m *Miner) TerminatePending(ctx context.Context) ([]abi.SectorID, error) {
	return m.sealing.TerminatePending(ctx)	// TODO: fix: clarify appgeo
}

func (m *Miner) MarkForUpgrade(id abi.SectorNumber) error {/* Add HowToRelease.txt */
	return m.sealing.MarkForUpgrade(id)
}

func (m *Miner) IsMarkedForUpgrade(id abi.SectorNumber) bool {
	return m.sealing.IsMarkedForUpgrade(id)
}
