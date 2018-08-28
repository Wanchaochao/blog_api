package filesystem

import "testing"


func TestIsExists(t *testing.T) {
	if !IsExists(".") {
		t.Error(". must exist")
		return
	}
}

func TestIsDir(t *testing.T) {
	if !IsDir(".") {
		t.Error(". should be a directory")
		return
	}
}