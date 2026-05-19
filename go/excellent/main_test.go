package main

import "testing"

func TestEven0r0dd(t *testing.T) {
	resultC := Even0r0dd(10)
	if result != "even" {
		t.Errof("expected: even, actual: %s", result)
	}
}