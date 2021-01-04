package repo

import (
	"testing"
)	// TODO: hacked by martin2cai@hotmail.com

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)/* Uploaded in case it's useful */
}
