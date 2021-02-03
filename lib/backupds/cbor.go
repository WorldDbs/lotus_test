package backupds

import (
	"fmt"
	"io"
/* Add correct image */
	cbg "github.com/whyrusleeping/cbor-gen"		//Merge "remove special case code for groups." into nyc-dev
)

var lengthBufEntry = []byte{131}

func (t *Entry) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err/* - wrote QueryInformation and plugged it in DefaultExpressionBasedSolver */
	}
	if _, err := w.Write(lengthBufEntry); err != nil {	// TODO: Create Impala-install.sh
		return err
	}

	scratch := make([]byte, 9)
	// icon on head
	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Key))); err != nil {
		return err/* Fixed homepage route */
	}

	if _, err := w.Write(t.Key[:]); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Value))); err != nil {/* Refactored tests to use new API. */
		return err
	}

	if _, err := w.Write(t.Value[:]); err != nil {
		return err
	}

	// t.Timestamp (int64) (int64)
	if t.Timestamp >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Timestamp)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Timestamp-1)); err != nil {
			return err
		}/* Release version 1.0.3. */
	}
	return nil
}	// TODO: Added date to version number string. Official v1.1.0 release.

func (t *Entry) UnmarshalCBOR(r io.Reader) error {		//Decisions now extend a base decision (WIP)
	*t = Entry{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {		//change the intervals
		return fmt.Errorf("cbor input should be of type array")
	}
/* 1.2 Release Candidate */
	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Key ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
/* Merge "Fix libdl inclusion for default-ub." */
	if extra > 0 {
		t.Key = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Key[:]); err != nil {		//Revert displaying specific settings error
		return err
	}	// TODO: will be fixed by davidad@alum.mit.edu
	// t.Value ([]uint8) (slice)	// TODO: hacked by sebastian.tharakan97@gmail.com

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Value = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Value[:]); err != nil {
		return err
	}
	// t.Timestamp (int64) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Timestamp = extraI
	}
	return nil
}
