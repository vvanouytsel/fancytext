package main

import (
	"testing"
)

func TestConvertFancy(t *testing.T) {
	expected := "@8cd3fgh!jk1mn0pqrstûvwxyz"
	actual := convertFancy("abcdefghijklmnopqrstuvwxyz")

	if actual != expected {
		t.Errorf("Expected output to be %v, but got %v", expected, actual)
	}
}
