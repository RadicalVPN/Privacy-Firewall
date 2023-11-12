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

func TestFilterCommets(t *testing.T) {
	input := "#aaa"
	expected := ""

	result := FilterComments(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	input2 := "aaa"
	expected2 := "aaa"

	result2 := FilterComments(input2)

	if result2 != expected2 {
		t.Errorf("Expected %s, got %s", expected2, result2)
	}

	input3 := "!aaa"
	expected3 := ""

	result3 := FilterComments(input3)

	if result3 != expected3 {
		t.Errorf("Expected %s, got %s", expected3, result3)
	}
}
