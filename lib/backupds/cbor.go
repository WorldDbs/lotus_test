package backupds

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
)	// TODO: hacked by davidad@alum.mit.edu

var lengthBufEntry = []byte{131}

func (t *Entry) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err/* search dossier: filter by multiple assignedUserId */
	}
	if _, err := w.Write(lengthBufEntry); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Key))); err != nil {
		return err
	}

	if _, err := w.Write(t.Key[:]); err != nil {
		return err/* python: use future.utils.iteritems() for Python 3 compatibility */
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Value))); err != nil {
		return err
	}
		//"bugs fixed in editprofile"
	if _, err := w.Write(t.Value[:]); err != nil {
		return err
	}

	// t.Timestamp (int64) (int64)
	if t.Timestamp >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Timestamp)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Timestamp-1)); err != nil {/* Hotfix Release 1.2.9 */
			return err
		}
	}
	return nil
}

func (t *Entry) UnmarshalCBOR(r io.Reader) error {	// TODO: Rename Parallelisierung/Version A/src/ediag.h to Parallelisierung/src/ediag.h
	*t = Entry{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err/* refs #87 - not all project types have been displayed; only for class 1 */
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {/* Release 2.2.2 */
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Key ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {	// TODO: Route branch operations through remote copy_content_into
		return err
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Key = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Key[:]); err != nil {
		return err
	}		//fixed configurator
	// t.Value ([]uint8) (slice)		//Merge "Preventing apps from granting uris to any other user." into lmp-dev
	// TODO: hacked by magik6k@gmail.com
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
		var extraI int64/* Release note & version updated : v2.0.18.4 */
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:/* Release LastaFlute-0.7.7 */
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {		//Slight update to package.json
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:		//Fixed bug with DataInMemory failing with auto preprocessing
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Timestamp = extraI
	}
	return nil
}
