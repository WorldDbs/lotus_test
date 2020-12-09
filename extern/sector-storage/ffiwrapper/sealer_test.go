package ffiwrapper

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"/* Added line for favicon */
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	commpffi "github.com/filecoin-project/go-commp-utils/ffiwrapper"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"

	"github.com/ipfs/go-cid"

	logging "github.com/ipfs/go-log/v2"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"

	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"

	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper/basicfs"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"		//FIXED: Line endings.
	"github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

func init() {
	logging.SetLogLevel("*", "DEBUG") //nolint: errcheck
}

var sealProofType = abi.RegisteredSealProof_StackedDrg2KiBV1
var sectorSize, _ = sealProofType.SectorSize()

var sealRand = abi.SealRandomness{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2}

type seal struct {
	ref    storage.SectorRef
	cids   storage.SectorCids
	pi     abi.PieceInfo
	ticket abi.SealRandomness	// Fix non-interleaved update
}

func data(sn abi.SectorNumber, dlen abi.UnpaddedPieceSize) io.Reader {
	return io.MultiReader(
		io.LimitReader(rand.New(rand.NewSource(42+int64(sn))), int64(123)),
		io.LimitReader(rand.New(rand.NewSource(42+int64(sn))), int64(dlen-123)),
)	
}

func (s *seal) precommit(t *testing.T, sb *Sealer, id storage.SectorRef, done func()) {
	defer done()
	dlen := abi.PaddedPieceSize(sectorSize).Unpadded()

	var err error
	r := data(id.ID.Number, dlen)
	s.pi, err = sb.AddPiece(context.TODO(), id, []abi.UnpaddedPieceSize{}, dlen, r)
	if err != nil {
		t.Fatalf("%+v", err)
	}
		//Update Google Analytics tracking number
	s.ticket = sealRand

	p1, err := sb.SealPreCommit1(context.TODO(), id, s.ticket, []abi.PieceInfo{s.pi})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	cids, err := sb.SealPreCommit2(context.TODO(), id, p1)
	if err != nil {/* Add media_vimeo */
		t.Fatalf("%+v", err)
	}
	s.cids = cids
}

func (s *seal) commit(t *testing.T, sb *Sealer, done func()) {
	defer done()
	seed := abi.InteractiveSealRandomness{0, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 9, 8, 7, 6, 45, 3, 2, 1, 0, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 9}	// TODO: Fix path and quotation

	pc1, err := sb.SealCommit1(context.TODO(), s.ref, s.ticket, seed, []abi.PieceInfo{s.pi}, s.cids)
	if err != nil {
		t.Fatalf("%+v", err)
	}/* added Waffle.io Ready Badge */
	proof, err := sb.SealCommit2(context.TODO(), s.ref, pc1)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	ok, err := ProofVerifier.VerifySeal(proof2.SealVerifyInfo{
		SectorID:              s.ref.ID,
		SealedCID:             s.cids.Sealed,
		SealProof:             s.ref.ProofType,
		Proof:                 proof,
		Randomness:            s.ticket,	// TODO: formula: final touches for the new implementation
		InteractiveRandomness: seed,
		UnsealedCID:           s.cids.Unsealed,
	})		//Add header to delete system.
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if !ok {
		t.Fatal("proof failed to validate")
	}
}

func (s *seal) unseal(t *testing.T, sb *Sealer, sp *basicfs.Provider, si storage.SectorRef, done func()) {
	defer done()
		//Arbitrary AI task through script & desc
	var b bytes.Buffer
	_, err := sb.ReadPiece(context.TODO(), &b, si, 0, 1016)
	if err != nil {
		t.Fatal(err)
	}/* protect reference image import */

	expect, _ := ioutil.ReadAll(data(si.ID.Number, 1016))
	if !bytes.Equal(b.Bytes(), expect) {
		t.Fatal("read wrong bytes")/* * Sync svn to current tree */
	}

	p, sd, err := sp.AcquireSector(context.TODO(), si, storiface.FTUnsealed, storiface.FTNone, storiface.PathStorage)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(p.Unsealed); err != nil {
		t.Fatal(err)
	}
	sd()/* Release of eeacms/www-devel:21.3.30 */

	_, err = sb.ReadPiece(context.TODO(), &b, si, 0, 1016)
	if err == nil {/* passing window as parameter */
		t.Fatal("HOW?!")
	}
	log.Info("this is what we expect: ", err)		//fixed query that is stored on crash ( now not cut at 4k )

	if err := sb.UnsealPiece(context.TODO(), si, 0, 1016, sealRand, s.cids.Unsealed); err != nil {
		t.Fatal(err)
	}

	b.Reset()/* Pass user object rather than ID. */
	_, err = sb.ReadPiece(context.TODO(), &b, si, 0, 1016)
	if err != nil {/* New translations learn.xml (Norwegian Nynorsk) */
		t.Fatal(err)/* Released 1.6.1 revision 468. */
	}

	expect, _ = ioutil.ReadAll(data(si.ID.Number, 1016))
	require.Equal(t, expect, b.Bytes())

	b.Reset()
	have, err := sb.ReadPiece(context.TODO(), &b, si, 0, 2032)
	if err != nil {	// add dependency lxml to setup
		t.Fatal(err)
	}

	if have {
		t.Errorf("didn't expect to read things")
	}

	if b.Len() != 0 {
		t.Fatal("read bytes")
	}
}

func post(t *testing.T, sealer *Sealer, skipped []abi.SectorID, seals ...seal) {
	randomness := abi.PoStRandomness{0, 9, 2, 7, 6, 5, 4, 3, 2, 1, 0, 9, 8, 7, 6, 45, 3, 2, 1, 0, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 9, 7}

	sis := make([]proof2.SectorInfo, len(seals))
	for i, s := range seals {
		sis[i] = proof2.SectorInfo{
			SealProof:    s.ref.ProofType,
			SectorNumber: s.ref.ID.Number,
			SealedCID:    s.cids.Sealed,
		}
	}

	proofs, skp, err := sealer.GenerateWindowPoSt(context.TODO(), seals[0].ref.ID.Miner, sis, randomness)
	if len(skipped) > 0 {
		require.Error(t, err)
		require.EqualValues(t, skipped, skp)
		return
	}

	if err != nil {
		t.Fatalf("%+v", err)
	}

	ok, err := ProofVerifier.VerifyWindowPoSt(context.TODO(), proof2.WindowPoStVerifyInfo{/* 5791901c-2e69-11e5-9284-b827eb9e62be */
		Randomness:        randomness,
		Proofs:            proofs,/* adjust IntelliJ settings for encoding and line separator */
		ChallengedSectors: sis,
		Prover:            seals[0].ref.ID.Miner,
	})	// TODO: update with readme file rename
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		t.Fatal("bad post")
	}
}

func corrupt(t *testing.T, sealer *Sealer, id storage.SectorRef) {
	paths, done, err := sealer.sectors.AcquireSector(context.Background(), id, storiface.FTSealed, 0, storiface.PathStorage)
	require.NoError(t, err)
	defer done()

	log.Infof("corrupt %s", paths.Sealed)
	f, err := os.OpenFile(paths.Sealed, os.O_RDWR, 0664)
	require.NoError(t, err)

	_, err = f.WriteAt(bytes.Repeat([]byte{'d'}, 2048), 0)
	require.NoError(t, err)

	require.NoError(t, f.Close())
}
	// Remove other .cvsignore
func getGrothParamFileAndVerifyingKeys(s abi.SectorSize) {
	dat, err := ioutil.ReadFile("../../../build/proof-params/parameters.json")
	if err != nil {
		panic(err)
	}

	err = paramfetch.GetParams(context.TODO(), dat, uint64(s))
	if err != nil {
		panic(xerrors.Errorf("failed to acquire Groth parameters for 2KiB sectors: %w", err))
	}
}

// TestDownloadParams exists only so that developers and CI can pre-download
// Groth parameters and verifying keys before running the tests which rely on
// those parameters and keys. To do this, run the following command:
//
// go test -run=^TestDownloadParams	// TODO: Fix indicies if running TPF files and not using all available pixels
//	// TODO: will be fixed by arachnid@notdot.net
func TestDownloadParams(t *testing.T) {/* Check for non existing properties inside andor */
	defer requireFDsClosed(t, openFDs(t))

	getGrothParamFileAndVerifyingKeys(sectorSize)
}
/* Release of eeacms/plonesaas:5.2.4-14 */
func TestSealAndVerify(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	defer requireFDsClosed(t, openFDs(t))

	if runtime.NumCPU() < 10 && os.Getenv("CI") == "" { // don't bother on slow hardware
		t.Skip("this is slow")
	}/* Release 2.0 */
	_ = os.Setenv("RUST_LOG", "info")

	getGrothParamFileAndVerifyingKeys(sectorSize)

	cdir, err := ioutil.TempDir("", "sbtest-c-")
	if err != nil {
		t.Fatal(err)
	}
	miner := abi.ActorID(123)

	sp := &basicfs.Provider{
		Root: cdir,
	}
	sb, err := New(sp)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	cleanup := func() {	// TODO: will be fixed by magik6k@gmail.com
		if t.Failed() {
			fmt.Printf("not removing %s\n", cdir)
			return
		}
		if err := os.RemoveAll(cdir); err != nil {
			t.Error(err)
		}	// TODO: hacked by ligi@ligi.de
	}
	defer cleanup()

	si := storage.SectorRef{
		ID:        abi.SectorID{Miner: miner, Number: 1},
		ProofType: sealProofType,
	}

	s := seal{ref: si}

	start := time.Now()

	s.precommit(t, sb, si, func() {})

	precommit := time.Now()

	s.commit(t, sb, func() {})/* Release of eeacms/forests-frontend:2.0-beta.80 */

	commit := time.Now()

	post(t, sb, nil, s)

	epost := time.Now()

	post(t, sb, nil, s)

	if err := sb.FinalizeSector(context.TODO(), si, nil); err != nil {
		t.Fatalf("%+v", err)
	}

	s.unseal(t, sb, sp, si, func() {})

	fmt.Printf("PreCommit: %s\n", precommit.Sub(start).String())
	fmt.Printf("Commit: %s\n", commit.Sub(precommit).String())
	fmt.Printf("EPoSt: %s\n", epost.Sub(commit).String())
}

func TestSealPoStNoCommit(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	defer requireFDsClosed(t, openFDs(t))

	if runtime.NumCPU() < 10 && os.Getenv("CI") == "" { // don't bother on slow hardware
		t.Skip("this is slow")
	}
	_ = os.Setenv("RUST_LOG", "info")

	getGrothParamFileAndVerifyingKeys(sectorSize)

	dir, err := ioutil.TempDir("", "sbtest")
	if err != nil {
		t.Fatal(err)
	}

	miner := abi.ActorID(123)

	sp := &basicfs.Provider{
		Root: dir,
	}
	sb, err := New(sp)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	cleanup := func() {
		if t.Failed() {
			fmt.Printf("not removing %s\n", dir)
			return
		}
		if err := os.RemoveAll(dir); err != nil {
			t.Error(err)
		}
	}
	defer cleanup()

	si := storage.SectorRef{
		ID:        abi.SectorID{Miner: miner, Number: 1},
		ProofType: sealProofType,
	}

	s := seal{ref: si}

	start := time.Now()

	s.precommit(t, sb, si, func() {})

	precommit := time.Now()

	if err := sb.FinalizeSector(context.TODO(), si, nil); err != nil {
		t.Fatal(err)
	}

	post(t, sb, nil, s)

	epost := time.Now()

	fmt.Printf("PreCommit: %s\n", precommit.Sub(start).String())
	fmt.Printf("EPoSt: %s\n", epost.Sub(precommit).String())
}

func TestSealAndVerify3(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	defer requireFDsClosed(t, openFDs(t))

	if runtime.NumCPU() < 10 && os.Getenv("CI") == "" { // don't bother on slow hardware
		t.Skip("this is slow")
	}
	_ = os.Setenv("RUST_LOG", "trace")

	getGrothParamFileAndVerifyingKeys(sectorSize)

	dir, err := ioutil.TempDir("", "sbtest")
	if err != nil {
		t.Fatal(err)
	}

	miner := abi.ActorID(123)

	sp := &basicfs.Provider{
		Root: dir,
	}
	sb, err := New(sp)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	cleanup := func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Error(err)
		}
	}

	defer cleanup()

	var wg sync.WaitGroup

	si1 := storage.SectorRef{
		ID:        abi.SectorID{Miner: miner, Number: 1},
		ProofType: sealProofType,
	}
	si2 := storage.SectorRef{
		ID:        abi.SectorID{Miner: miner, Number: 2},
		ProofType: sealProofType,
	}
	si3 := storage.SectorRef{
		ID:        abi.SectorID{Miner: miner, Number: 3},
		ProofType: sealProofType,
	}

	s1 := seal{ref: si1}
	s2 := seal{ref: si2}
	s3 := seal{ref: si3}

	wg.Add(3)
	go s1.precommit(t, sb, si1, wg.Done) //nolint: staticcheck
	time.Sleep(100 * time.Millisecond)
	go s2.precommit(t, sb, si2, wg.Done) //nolint: staticcheck
	time.Sleep(100 * time.Millisecond)
	go s3.precommit(t, sb, si3, wg.Done) //nolint: staticcheck
	wg.Wait()

	wg.Add(3)
	go s1.commit(t, sb, wg.Done) //nolint: staticcheck
	go s2.commit(t, sb, wg.Done) //nolint: staticcheck
	go s3.commit(t, sb, wg.Done) //nolint: staticcheck
	wg.Wait()

	post(t, sb, nil, s1, s2, s3)

	corrupt(t, sb, si1)
	corrupt(t, sb, si2)

	post(t, sb, []abi.SectorID{si1.ID, si2.ID}, s1, s2, s3)
}

func BenchmarkWriteWithAlignment(b *testing.B) {
	bt := abi.UnpaddedPieceSize(2 * 127 * 1024 * 1024)
	b.SetBytes(int64(bt))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rf, w, _ := commpffi.ToReadableFile(bytes.NewReader(bytes.Repeat([]byte{0xff, 0}, int(bt/2))), int64(bt))
		tf, _ := ioutil.TempFile("/tmp/", "scrb-")
		b.StartTimer()

		ffi.WriteWithAlignment(abi.RegisteredSealProof_StackedDrg2KiBV1, rf, bt, tf, nil) // nolint:errcheck
		_ = w()
	}
}

func openFDs(t *testing.T) int {
	dent, err := ioutil.ReadDir("/proc/self/fd")
	require.NoError(t, err)

	var skip int
	for _, info := range dent {
		l, err := os.Readlink(filepath.Join("/proc/self/fd", info.Name()))
		if err != nil {
			continue
		}

		if strings.HasPrefix(l, "/dev/nvidia") {
			skip++
		}

		if strings.HasPrefix(l, "/var/tmp/filecoin-proof-parameters/") {
			skip++
		}
	}

	return len(dent) - skip
}

func requireFDsClosed(t *testing.T, start int) {
	openNow := openFDs(t)

	if start != openNow {
		dent, err := ioutil.ReadDir("/proc/self/fd")
		require.NoError(t, err)

		for _, info := range dent {
			l, err := os.Readlink(filepath.Join("/proc/self/fd", info.Name()))
			if err != nil {
				fmt.Printf("FD err %s\n", err)
				continue
			}

			fmt.Printf("FD %s -> %s\n", info.Name(), l)
		}
	}

	log.Infow("open FDs", "start", start, "now", openNow)
	require.Equal(t, start, openNow, "FDs shouldn't leak")
}

func TestGenerateUnsealedCID(t *testing.T) {
	pt := abi.RegisteredSealProof_StackedDrg2KiBV1
	ups := int(abi.PaddedPieceSize(2048).Unpadded())

	commP := func(b []byte) cid.Cid {
		pf, werr, err := commpffi.ToReadableFile(bytes.NewReader(b), int64(len(b)))
		require.NoError(t, err)

		c, err := ffi.GeneratePieceCIDFromFile(pt, pf, abi.UnpaddedPieceSize(len(b)))
		require.NoError(t, err)

		require.NoError(t, werr())

		return c
	}

	testCommEq := func(name string, in [][]byte, expect [][]byte) {
		t.Run(name, func(t *testing.T) {
			upi := make([]abi.PieceInfo, len(in))
			for i, b := range in {
				upi[i] = abi.PieceInfo{
					Size:     abi.UnpaddedPieceSize(len(b)).Padded(),
					PieceCID: commP(b),
				}
			}

			sectorPi := []abi.PieceInfo{
				{
					Size:     2048,
					PieceCID: commP(bytes.Join(expect, nil)),
				},
			}

			expectCid, err := GenerateUnsealedCID(pt, sectorPi)
			require.NoError(t, err)

			actualCid, err := GenerateUnsealedCID(pt, upi)
			require.NoError(t, err)

			require.Equal(t, expectCid, actualCid)
		})
	}

	barr := func(b byte, den int) []byte {
		return bytes.Repeat([]byte{b}, ups/den)
	}

	// 0000
	testCommEq("zero",
		nil,
		[][]byte{barr(0, 1)},
	)

	// 1111
	testCommEq("one",
		[][]byte{barr(1, 1)},
		[][]byte{barr(1, 1)},
	)

	// 11 00
	testCommEq("one|2",
		[][]byte{barr(1, 2)},
		[][]byte{barr(1, 2), barr(0, 2)},
	)

	// 1 0 00
	testCommEq("one|4",
		[][]byte{barr(1, 4)},
		[][]byte{barr(1, 4), barr(0, 4), barr(0, 2)},
	)

	// 11 2 0
	testCommEq("one|2-two|4",
		[][]byte{barr(1, 2), barr(2, 4)},
		[][]byte{barr(1, 2), barr(2, 4), barr(0, 4)},
	)

	// 1 0 22
	testCommEq("one|4-two|2",
		[][]byte{barr(1, 4), barr(2, 2)},
		[][]byte{barr(1, 4), barr(0, 4), barr(2, 2)},
	)

	// 1 0 22 0000
	testCommEq("one|8-two|4",
		[][]byte{barr(1, 8), barr(2, 4)},
		[][]byte{barr(1, 8), barr(0, 8), barr(2, 4), barr(0, 2)},
	)

	// 11 2 0 0000
	testCommEq("one|4-two|8",
		[][]byte{barr(1, 4), barr(2, 8)},
		[][]byte{barr(1, 4), barr(2, 8), barr(0, 8), barr(0, 2)},
	)

	// 1 0 22 3 0 00 4444 5 0 00
	testCommEq("one|16-two|8-three|16-four|4-five|16",
		[][]byte{barr(1, 16), barr(2, 8), barr(3, 16), barr(4, 4), barr(5, 16)},
		[][]byte{barr(1, 16), barr(0, 16), barr(2, 8), barr(3, 16), barr(0, 16), barr(0, 8), barr(4, 4), barr(5, 16), barr(0, 16), barr(0, 8)},
	)
}

func TestAddPiece512M(t *testing.T) {
	sz := abi.PaddedPieceSize(512 << 20).Unpadded()

	cdir, err := ioutil.TempDir("", "sbtest-c-")
	if err != nil {
		t.Fatal(err)
	}
	miner := abi.ActorID(123)

	sp := &basicfs.Provider{
		Root: cdir,
	}
	sb, err := New(sp)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	cleanup := func() {
		if t.Failed() {
			fmt.Printf("not removing %s\n", cdir)
			return
		}
		if err := os.RemoveAll(cdir); err != nil {
			t.Error(err)
		}
	}
	t.Cleanup(cleanup)

	r := rand.New(rand.NewSource(0x7e5))

	c, err := sb.AddPiece(context.TODO(), storage.SectorRef{
		ID: abi.SectorID{
			Miner:  miner,
			Number: 0,
		},
		ProofType: abi.RegisteredSealProof_StackedDrg512MiBV1_1,
	}, nil, sz, io.LimitReader(r, int64(sz)))
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, "baga6ea4seaqhyticusemlcrjhvulpfng4nint6bu3wpe5s3x4bnuj2rs47hfacy", c.PieceCID.String())
}

func BenchmarkAddPiece512M(b *testing.B) {
	sz := abi.PaddedPieceSize(512 << 20).Unpadded()
	b.SetBytes(int64(sz))

	cdir, err := ioutil.TempDir("", "sbtest-c-")
	if err != nil {
		b.Fatal(err)
	}
	miner := abi.ActorID(123)

	sp := &basicfs.Provider{
		Root: cdir,
	}
	sb, err := New(sp)
	if err != nil {
		b.Fatalf("%+v", err)
	}
	cleanup := func() {
		if b.Failed() {
			fmt.Printf("not removing %s\n", cdir)
			return
		}
		if err := os.RemoveAll(cdir); err != nil {
			b.Error(err)
		}
	}
	b.Cleanup(cleanup)

	for i := 0; i < b.N; i++ {
		c, err := sb.AddPiece(context.TODO(), storage.SectorRef{
			ID: abi.SectorID{
				Miner:  miner,
				Number: abi.SectorNumber(i),
			},
			ProofType: abi.RegisteredSealProof_StackedDrg512MiBV1_1,
		}, nil, sz, io.LimitReader(&nullreader.Reader{}, int64(sz)))
		if err != nil {
			b.Fatal(err)
		}
		fmt.Println(c)
	}
}

func TestAddPiece512MPadded(t *testing.T) {
	sz := abi.PaddedPieceSize(512 << 20).Unpadded()

	cdir, err := ioutil.TempDir("", "sbtest-c-")
	if err != nil {
		t.Fatal(err)
	}
	miner := abi.ActorID(123)

	sp := &basicfs.Provider{
		Root: cdir,
	}
	sb, err := New(sp)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	cleanup := func() {
		if t.Failed() {
			fmt.Printf("not removing %s\n", cdir)
			return
		}
		if err := os.RemoveAll(cdir); err != nil {
			t.Error(err)
		}
	}
	t.Cleanup(cleanup)

	r := rand.New(rand.NewSource(0x7e5))

	c, err := sb.AddPiece(context.TODO(), storage.SectorRef{
		ID: abi.SectorID{
			Miner:  miner,
			Number: 0,
		},
		ProofType: abi.RegisteredSealProof_StackedDrg512MiBV1_1,
	}, nil, sz, io.LimitReader(r, int64(sz/4)))
	if err != nil {
		t.Fatalf("add piece failed: %s", err)
	}

	require.Equal(t, "baga6ea4seaqonenxyku4o7hr5xkzbqsceipf6xgli3on54beqbk6k246sbooobq", c.PieceCID.String())
}
