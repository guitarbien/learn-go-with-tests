package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://youcantreachme.ooxx" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://youcantreachme.ooxx",
	}

	actualResult := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResult)

	if want != got {
		t.Fatalf("wanted %v, got %v", want, got)
	}

	expectedResult := map[string]bool{
		"https://google.com": true,
		"https://facebook.com": true,
		"https://youcantreachme.ooxx": false,
	}

	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Fatalf("wanted %v, got %v", expectedResult, actualResult)
	}
}
