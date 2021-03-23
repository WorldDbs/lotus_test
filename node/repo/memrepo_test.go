package repo/* Version 3 Release Notes */

import (
	"testing"
)
/* Complete the "Favorite" feature for PatchReleaseManager; */
func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
