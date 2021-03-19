package repo

import (
	"testing"
)/* certdb/Main: remove obsolete option "--all" */
	// TODO: removed output files from svn
func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}/* Merge branch 'master' into travis_Release */
