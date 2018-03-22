package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	strings := []string{"Hello", "Хелло"}
	wantStrings := []string{"olleH", "оллеХ"}

	for i := 0; i < len(strings); i++ {
		s := strings[i]
		got := Reverse(s)
		want := wantStrings[i]
		if got != want {
			t.Fatalf("got %q, expected %q", got, want)
		}
	}
}
