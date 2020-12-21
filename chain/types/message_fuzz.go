//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message	// TODO: Fix misspelling of "tries"
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))/* Added cynthia's picture */
	if err != nil {	// TODO: hacked by fkautz@pseudocode.cc
		panic(err) // ok
	}/* Release des locks ventouses */
	reData2, err := msg.Serialize()	// TODO: Update Updater.cs
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}
