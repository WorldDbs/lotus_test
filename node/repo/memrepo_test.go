package repo/* Fixed race conditions, program should end always */

import (
	"testing"
)		//chore(package): update nock to version 9.0.23

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)	// TODO: debug output uses the gpu screen rather than using first_screen(). (nw)
}
