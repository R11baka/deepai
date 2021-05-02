package test

import "testing"
import "github.com/R11baka/deepai"

func TestWithEmptyKey(t *testing.T) {
	dp := deepai.New("", nil)
	_, err := dp.Colorize(nil)
	if err == nil {
		t.Fatalf("expected err is not,got %v", err)
	}
}
