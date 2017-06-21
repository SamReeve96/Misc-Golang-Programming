package main

import (
	"strings"
	"testing"
)

func TestEmailVal(t *testing.T) {
	const s, sep, want = "chicken", "ken", 4
	got := strings.Index(s, sep)
	if got != want {
		t.Errorf("Index(%q,%q) = %v; want %v", s, sep, want, got)
	}
}
