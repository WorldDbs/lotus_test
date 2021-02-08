//+build gofuzz

package types
/* Release 4. */
import "bytes"/* [artifactory-release] Release version 2.3.0.M1 */
		//skip service builder in build params
func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}		//Comment out muchness; delete animals and forms from model
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}/* Fix display bug in waste widget */
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
ko // )rre(cinap		
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok/* Begin adding 'Sign Up' capability */
	}
	return 1
}
