//+build gofuzz
/* fix bug that results in unneeded slects when using find */
package types

import "bytes"

func FuzzMessage(data []byte) int {	// Merge branch 'develop' into 4227-omit-private-changelog
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}		//Updated to 1.11 - for OSX
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}/* [ADD] Debian Ubuntu Releases */
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}		//Delete eta3.launch~
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok		//Imported Debian patch 1.0b2-10
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}
