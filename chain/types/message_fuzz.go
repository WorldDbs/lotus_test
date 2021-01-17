//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message/* Upgrade Maven Release Plugin to the current version */
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {/* Removed all but one reference to ActiveMQ in tests and connector.  */
		return 0
	}		//Rename split.py to Testing/split.py
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok/* Merge "docs: Release notes for support lib v20" into klp-modular-dev */
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok/* Release 0.7.1 Alpha */
	}
	reData2, err := msg.Serialize()/* temperature */
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok/* WL-2589 Switch to one map set for skills. */
	}
	return 1
}
