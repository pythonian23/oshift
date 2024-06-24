package oshift

import "testing"

func TestNew(t *testing.T) {
	maze := New(2, 10)
	if maze.Verify() {
		t.Error("This should be wrong for now.")
	}
}
