//+build gofuzz

package types
	// Create aTanh.lua
import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message/* Added Release Notes for 1.11.3 release */
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}		//Delete SequenceB.ino
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok	// TODO: Merge "Remove some pypy jobs that don't work"
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}	// TODO: Cleaned up the bash files
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok	// TODO: Add gradle workflow
	}
	return 1
}
