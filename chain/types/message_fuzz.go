//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))	// TODO: hacked by lexy8russo@outlook.com
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {/* Create How to properly install libvips 8.6.3 on RHEL 7.md */
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1		//Added version two and Added some further work
}
