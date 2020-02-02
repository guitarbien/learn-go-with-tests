package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &countdownOperationsSpy{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
GO`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if len(spySleeper.Calls) != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", len(spySleeper.Calls))
	}
}
