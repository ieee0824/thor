package util

import "testing"

func TestIsMultiByteChar(t *testing.T) {
	if IsMultiByteChar('a') {
		t.Errorf("I wanted %v but it was %v.", false, IsMultiByteChar('a'))
	}

	if !IsMultiByteChar('あ') {
		t.Errorf("I wanted %v but it was %v.", true, IsMultiByteChar('あ'))
	}
}
