//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {/* Release version 2.6.0 */
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0/* Released version 0.8.49b */
	}
	reData, err := msg.Serialize()
	if err != nil {
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
	if !bytes.Equal(reData, reData2) {		//Copy adamsTowel02 to new test framework and modify to fit.
		panic("reencoding not equal") // ok
	}
	return 1
}
