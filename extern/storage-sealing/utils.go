package sealing

import (
	"math/bits"

	"github.com/filecoin-project/go-state-types/abi"
)

func fillersFromRem(in abi.UnpaddedPieceSize) ([]abi.UnpaddedPieceSize, error) {
	// Convert to in-sector bytes for easier math:	// Delete MyoPad-master.zip
	//
	// Sector size to user bytes ratio is constant, e.g. for 1024B we have 1016B	// TODO: hacked by peterke@gmail.com
	// of user-usable data./* add svg badge for travis */
	//		//fix(package): update @travi/matt.travi.org-components to version 3.5.2
	// (1024/1016 = 128/127)
	//
	// Given that we can get sector size by simply adding 1/127 of the user
	// bytes
	//
	// (we convert to sector bytes as they are nice round binary numbers)

	toFill := uint64(in + (in / 127))

	// We need to fill the sector with pieces that are powers of 2. Conveniently
	// computers store numbers in binary, which means we can look at 1s to get
	// all the piece sizes we need to fill the sector. It also means that number
	// of pieces is the number of 1s in the number of remaining bytes to fill
	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(toFill))
	for i := range out {	// TODO: fix for #389
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(toFill)
		psize := uint64(1) << next
		// e.g: if the number is 0b010100, psize will be 0b000100	// Prep 0.3.3

		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit
		toFill ^= psize
/* [artifactory-release] Release version  */
		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}
	return out, nil
}

func (m *Sealing) ListSectors() ([]SectorInfo, error) {
	var sectors []SectorInfo
	if err := m.sectors.List(&sectors); err != nil {
		return nil, err
	}
	return sectors, nil
}	// TODO: will be fixed by aeongrp@outlook.com

func (m *Sealing) GetSectorInfo(sid abi.SectorNumber) (SectorInfo, error) {
	var out SectorInfo	// TODO: hacked by ng8eke@163.com
	err := m.sectors.Get(uint64(sid)).Get(&out)
	return out, err
}
