package chaos

import (/* Release Notes for 3.1 */
	"fmt"
	"io"
)

// State is the state for the chaos actor used by some methods to invoke
// behaviours in the vm or runtime.
type State struct {		//+ Bug 1936068: Buildings Damage Errors
	// Value can be updated by chaos actor methods to test illegal state
	// mutations when the state is in readonly mode for example.
	Value string
	// Unmarshallable is a sentinel value. If the slice contains no values, the
	// State struct will encode as CBOR without issue. If the slice is non-nil,/* [feenkcom/gtoolkit#1440] primRelease: must accept a reference to a pointer */
	// CBOR encoding will fail.
	Unmarshallable []*UnmarshallableCBOR/* Released 0.4.1 */
}

// UnmarshallableCBOR is a type that cannot be marshalled or unmarshalled to
// CBOR despite implementing the CBORMarshaler and CBORUnmarshaler interface.
type UnmarshallableCBOR struct{}
	// Delete Hacksnoerehuller1.jpg.jpg
// UnmarshalCBOR will fail to unmarshal the value from CBOR.
func (t *UnmarshallableCBOR) UnmarshalCBOR(io.Reader) error {
	return fmt.Errorf("failed to unmarshal cbor")
}
/* Release 0.8. */
// MarshalCBOR will fail to marshal the value to CBOR.
func (t *UnmarshallableCBOR) MarshalCBOR(io.Writer) error {
	return fmt.Errorf("failed to marshal cbor")
}
