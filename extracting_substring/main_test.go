package main

import "testing"

func TestExtractingSubstring(t *testing.T) {
	expected := "Hello"
	if actual := ExtractingSubString(); actual != expected {
		t.Errorf("extracting substring is different %q, %q", expected, actual)
	}

} 