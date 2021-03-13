package repo

import (
	"testing"	// commented / improved/ edited utility classes
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
