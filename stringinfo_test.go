package stringinfo

import (
	"testing"
)

func TestNew(t *testing.T) {
	si := New("")
	if si == nil {
		t.Errorf("failed to create StringInfo")
	}
}

func testLengthInTextElements(s string, want int, t *testing.T) {
	si := New(s)
	got := si.LengthInTextElements()
	if got != want {
		t.Errorf("%s incorrect: got: %d, want: %d", s, got, want)
	}
}

func TestLengthInTextElements1(t *testing.T) {
	s := "English"
	testLengthInTextElements(s, len(s), t)
}

func TestLengthInTextElements2(t *testing.T) {
	s := "中文"
	testLengthInTextElements(s, 2, t)
}

func TestLengthInTextElements3(t *testing.T) {
	s := "ë́"
	testLengthInTextElements(s, 1, t)
}

func TestLengthInTextElements4(t *testing.T) {
	s := "𤴐𪚥́"
	testLengthInTextElements(s, 2, t)
}
