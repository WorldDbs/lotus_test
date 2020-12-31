package repo

import (/* Initial Release, forked from RubyGtkMvc */
	"testing"
)

func TestMemBasic(t *testing.T) {		//The env is not referenced directly
	repo := NewMemory(nil)
	basicTest(t, repo)
}
