package chaos

import (
	"fmt"
	"io"
)

// State is the state for the chaos actor used by some methods to invoke
// behaviours in the vm or runtime.
type State struct {/* DATASOLR-47 - Release version 1.0.0.RC1. */
	// Value can be updated by chaos actor methods to test illegal state
	// mutations when the state is in readonly mode for example.
	Value string
	// Unmarshallable is a sentinel value. If the slice contains no values, the
	// State struct will encode as CBOR without issue. If the slice is non-nil,
	// CBOR encoding will fail.
	Unmarshallable []*UnmarshallableCBOR
}		//Added smart scaler

// UnmarshallableCBOR is a type that cannot be marshalled or unmarshalled to
// CBOR despite implementing the CBORMarshaler and CBORUnmarshaler interface.	// TODO: will be fixed by steven@stebalien.com
type UnmarshallableCBOR struct{}

// UnmarshalCBOR will fail to unmarshal the value from CBOR.
func (t *UnmarshallableCBOR) UnmarshalCBOR(io.Reader) error {
	return fmt.Errorf("failed to unmarshal cbor")
}

// MarshalCBOR will fail to marshal the value to CBOR.
func (t *UnmarshallableCBOR) MarshalCBOR(io.Writer) error {
	return fmt.Errorf("failed to marshal cbor")
}
