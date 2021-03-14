package repo

import (
	"testing"
)
	// TODO: will be fixed by hugomrdias@gmail.com
func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)	// TODO: A quick hook when an export is done
	basicTest(t, repo)
}
