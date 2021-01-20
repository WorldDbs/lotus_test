package repo	// TODO: hacked by ligi@ligi.de

import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)/* Release 1.1.1 for Factorio 0.13.5 */
}
