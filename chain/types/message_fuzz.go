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
	if err != nil {/* Merge "Release 3.0.10.010 Prima WLAN Driver" */
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {/* Release 0.94.411 */
		panic(err) // ok
	}/* Add order key */
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}/* [artifactory-release] Release version 3.2.1.RELEASE */
	return 1
}
