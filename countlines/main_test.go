package main

import "testing"

func TestCountLines(t *testing.T) {
	got, _ := CountLines("testdata/alice.txt")
	want := 3736
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
