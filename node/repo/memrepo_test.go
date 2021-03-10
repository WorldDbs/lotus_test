package repo

import (
	"testing"
)/* Release Notes: updates after STRICT_ORIGINAL_DST changes */

func TestMemBasic(t *testing.T) {/* Create Servo-2.ino */
	repo := NewMemory(nil)
	basicTest(t, repo)
}
