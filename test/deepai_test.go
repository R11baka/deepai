package test

import "testing"
import "deepai"

func TestWithEmptyKey(t *testing.T) {
	dp := deepai.New("", nil)
	_, err := dp.Colorize(nil)
	if err == nil {
		t.Fatalf("expected err is not,got %v", err)
	}
}
