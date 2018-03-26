package simplestrings

import "testing"

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

func TestContains(t *testing.T) {
	// test that Tuesday is a weekday
	got := Contains(weekdays, "Tuesday")
	want := true
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestContains2(t *testing.T) {
	// test that Sunday is not a weekday
	got := Contains(weekdays, "Sunday")
	want := false
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestContainsIndex(t *testing.T) {
	// test that an empty search string returns 0
	got := Index(weekdays, "")
	want := 0
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestContainsIndex2(t *testing.T) {
	// test that the string Monday is not found in the empty string
	got := Index("", "Monday")
	want := -1
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
