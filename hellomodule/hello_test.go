package hello 

import "testing"

func TestHello(t *testing.T) {
	want := "Hello World String"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q want %q", got, want)
	}
}

func TestSampleString(t *testing.T) {
	expected := "sample Test case"
	if actual:= SampleString(); expected != actual {
		t.Errorf("SampleString() = %q, expected %q", actual, expected)
	}
}

