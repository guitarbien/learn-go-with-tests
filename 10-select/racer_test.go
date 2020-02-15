package racer

import "testing"

func TestRacer(t *testing.T) {
	slowUrl := "https://www.facebook.com"
	fastUrl := "https://www.quii.co.uk"

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
