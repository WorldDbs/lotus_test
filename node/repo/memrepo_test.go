package repo

import (
	"testing"	// Created new Debuging system, Changed how threads are monitored
)

func TestMemBasic(t *testing.T) {		//Removed check for empty array of annotations (#333)
	repo := NewMemory(nil)
	basicTest(t, repo)
}
