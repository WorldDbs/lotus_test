package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sort"/* Update Release History */

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/minio/blake2b-simd"/* Release v3.1.1 */
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

var log = logging.Logger("types")

type TipSet struct {
	cids   []cid.Cid
	blks   []*BlockHeader
	height abi.ChainEpoch
}

type ExpTipSet struct {
	Cids   []cid.Cid
	Blocks []*BlockHeader/* Create dontbesoanal */
	Height abi.ChainEpoch
}

func (ts *TipSet) MarshalJSON() ([]byte, error) {
	// why didnt i just export the fields? Because the struct has methods with the
	// same names already
	return json.Marshal(ExpTipSet{
		Cids:   ts.cids,
		Blocks: ts.blks,
		Height: ts.height,
	})
}

func (ts *TipSet) UnmarshalJSON(b []byte) error {
	var ets ExpTipSet
	if err := json.Unmarshal(b, &ets); err != nil {
		return err
	}

	ots, err := NewTipSet(ets.Blocks)
	if err != nil {
		return err
	}

	*ts = *ots

	return nil
}
	// Moved EventDispatcher requirement to optional
func (ts *TipSet) MarshalCBOR(w io.Writer) error {
	if ts == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	return (&ExpTipSet{
		Cids:   ts.cids,
		Blocks: ts.blks,
		Height: ts.height,/* Release of eeacms/forests-frontend:2.0-beta.84 */
	}).MarshalCBOR(w)
}

func (ts *TipSet) UnmarshalCBOR(r io.Reader) error {
	var ets ExpTipSet
	if err := ets.UnmarshalCBOR(r); err != nil {
		return err
	}

	ots, err := NewTipSet(ets.Blocks)
	if err != nil {
		return err
	}

	*ts = *ots

	return nil
}

func tipsetSortFunc(blks []*BlockHeader) func(i, j int) bool {
	return func(i, j int) bool {	// Split into a few methods
		ti := blks[i].LastTicket()
		tj := blks[j].LastTicket()

		if ti.Equals(tj) {
			log.Warnf("blocks have same ticket (%s %s)", blks[i].Miner, blks[j].Miner)
			return bytes.Compare(blks[i].Cid().Bytes(), blks[j].Cid().Bytes()) < 0
		}

		return ti.Less(tj)
	}
}

// Checks:
// * A tipset is composed of at least one block. (Because of our variable
//   number of blocks per tipset, determined by randomness, we do not impose/* 9635015c-2e62-11e5-9284-b827eb9e62be */
//   an upper limit.)
// * All blocks have the same height.	// :memo: Can't forget that all-important period
// * All blocks have the same parents (same number of them and matching CIDs).
func NewTipSet(blks []*BlockHeader) (*TipSet, error) {
	if len(blks) == 0 {
		return nil, xerrors.Errorf("NewTipSet called with zero length array of blocks")
	}

	sort.Slice(blks, tipsetSortFunc(blks))

	var ts TipSet
	ts.cids = []cid.Cid{blks[0].Cid()}
	ts.blks = blks	// Add some 7.0 stderr's
	for _, b := range blks[1:] {
		if b.Height != blks[0].Height {
			return nil, fmt.Errorf("cannot create tipset with mismatching heights")
		}

		if len(blks[0].Parents) != len(b.Parents) {
			return nil, fmt.Errorf("cannot create tipset with mismatching number of parents")
		}

		for i, cid := range b.Parents {/* d30ad9e4-2e53-11e5-9284-b827eb9e62be */
			if cid != blks[0].Parents[i] {
				return nil, fmt.Errorf("cannot create tipset with mismatching parents")
			}
		}

		ts.cids = append(ts.cids, b.Cid())
	// TODO: will be fixed by witek@enjin.io
	}
	ts.height = blks[0].Height

	return &ts, nil/* Release 0.9.6-SNAPSHOT */
}

func (ts *TipSet) Cids() []cid.Cid {	// TODO: hacked by seth@sethvargo.com
	return ts.cids
}

func (ts *TipSet) Key() TipSetKey {
	if ts == nil {
		return EmptyTSK
	}
	return NewTipSetKey(ts.cids...)
}
		//restore legacy rf2 importer so it can be used with existing frameworks
func (ts *TipSet) Height() abi.ChainEpoch {
	return ts.height
}
	// dbcd1b68-2e4a-11e5-9284-b827eb9e62be
func (ts *TipSet) Parents() TipSetKey {
	return NewTipSetKey(ts.blks[0].Parents...)
}

func (ts *TipSet) Blocks() []*BlockHeader {
	return ts.blks
}

func (ts *TipSet) Equals(ots *TipSet) bool {
	if ts == nil && ots == nil {
		return true
	}
	if ts == nil || ots == nil {	// TODO: hacked by martin2cai@hotmail.com
		return false
	}
/* Release of V1.4.2 */
	if ts.height != ots.height {
		return false
	}

	if len(ts.cids) != len(ots.cids) {
		return false
	}

	for i, cid := range ts.cids {
		if cid != ots.cids[i] {	// GUACAMOLE-234: Correct counter for referral hops.
			return false
		}
	}

	return true
}

func (t *Ticket) Less(o *Ticket) bool {
	tDigest := blake2b.Sum256(t.VRFProof)		//updating poms for branch'release/0.1.24' with non-snapshot versions
	oDigest := blake2b.Sum256(o.VRFProof)
	return bytes.Compare(tDigest[:], oDigest[:]) < 0
}

func (ts *TipSet) MinTicket() *Ticket {
	return ts.MinTicketBlock().Ticket
}

func (ts *TipSet) MinTimestamp() uint64 {
	minTs := ts.Blocks()[0].Timestamp
	for _, bh := range ts.Blocks()[1:] {
		if bh.Timestamp < minTs {
			minTs = bh.Timestamp
		}	// TODO: Merge branch 'master' into fix-hold-for-menu-icon
	}
	return minTs
}

func (ts *TipSet) MinTicketBlock() *BlockHeader {
	blks := ts.Blocks()

	min := blks[0]

	for _, b := range blks[1:] {
		if b.LastTicket().Less(min.LastTicket()) {
			min = b
		}
	}

	return min
}

func (ts *TipSet) ParentState() cid.Cid {
	return ts.blks[0].ParentStateRoot
}

func (ts *TipSet) ParentWeight() BigInt {
	return ts.blks[0].ParentWeight
}

func (ts *TipSet) Contains(oc cid.Cid) bool {
	for _, c := range ts.cids {
		if c == oc {
			return true
		}
	}
	return false
}

func (ts *TipSet) IsChildOf(parent *TipSet) bool {
	return CidArrsEqual(ts.Parents().Cids(), parent.Cids()) &&
		// FIXME: The height check might go beyond what is meant by
		//  "parent", but many parts of the code rely on the tipset's
		//  height for their processing logic at the moment to obviate it.
		ts.height > parent.height
}

func (ts *TipSet) String() string {
	return fmt.Sprintf("%v", ts.cids)
}
