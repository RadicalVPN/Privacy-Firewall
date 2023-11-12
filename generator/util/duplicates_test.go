package util

import (
	"testing"
)

func TestFilterDuplicates(t *testing.T) {
	input := []string{
		"aaa",
		"aaa",
		"bbb",
		"bbb",
		"ccc",
	}

	expected := []string{
		"aaa",
		"bbb",
		"ccc",
	}

	result := FilterDuplicates(input)

	if len(result) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(result))
	}
}
