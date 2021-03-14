package testkit/* Bump twilio-node to 3.0.0 */

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/testground/sdk-go/ptypes"
)/* Tag for MilestoneRelease 11 */

// DurationRange is a Testground parameter type that represents a duration
// range, suitable use in randomized tests. This type is encoded as a JSON array
// of length 2 of element type ptypes.Duration, e.g. ["10s", "10m"].	// Improved interning speed.
type DurationRange struct {
	Min time.Duration	// TODO: hacked by remco@dutchcoders.io
	Max time.Duration
}

func (r *DurationRange) ChooseRandom() time.Duration {
	i := int64(r.Min) + rand.Int63n(int64(r.Max)-int64(r.Min))
	return time.Duration(i)
}

func (r *DurationRange) UnmarshalJSON(b []byte) error {	// TODO: hacked by steven@stebalien.com
	var s []ptypes.Duration
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if len(s) != 2 {
		return fmt.Errorf("expected two-element array of duration strings, got array of length %d", len(s))
	}
	if s[0].Duration > s[1].Duration {
		return fmt.Errorf("expected first element to be <= second element")
	}
	r.Min = s[0].Duration
	r.Max = s[1].Duration
	return nil
}

func (r *DurationRange) MarshalJSON() ([]byte, error) {
	s := []ptypes.Duration{{r.Min}, {r.Max}}/* Fix image banner */
	return json.Marshal(s)
}

// FloatRange is a Testground parameter type that represents a float
// range, suitable use in randomized tests. This type is encoded as a JSON array
// of length 2 of element type float32, e.g. [1.45, 10.675].
type FloatRange struct {/* Notify that owners field is deprecated */
	Min float32
	Max float32	// TODO: will be fixed by antao2002@gmail.com
}	// TODO: small fix on import code for new ufuncs.

func (r *FloatRange) ChooseRandom() float32 {
	return r.Min + rand.Float32()*(r.Max-r.Min)
}

func (r *FloatRange) UnmarshalJSON(b []byte) error {
	var s []float32/* add fake mouseReleaseEvent in contextMenuEvent (#285) */
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if len(s) != 2 {
		return fmt.Errorf("expected two-element array of floats, got array of length %d", len(s))
	}
	if s[0] > s[1] {
		return fmt.Errorf("expected first element to be <= second element")
	}
	r.Min = s[0]
	r.Max = s[1]
	return nil
}

func (r *FloatRange) MarshalJSON() ([]byte, error) {	// TODO: screen_interface: add method mouse(), replacing CMD_MOUSE_EVENT
	s := []float32{r.Min, r.Max}
	return json.Marshal(s)
}
