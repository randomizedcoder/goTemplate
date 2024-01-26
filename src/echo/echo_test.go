package echo

import (
	"testing"
)

func TestLoudly(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},
		{"world", "WORLD"},
		{"", ""},
	}

	for _, test := range tests {
		result := EchoLoudly(test.input)
		if result != test.expected {
			t.Errorf("Loudly(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestEchoRepeat(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "hello"},
		{"world", "world"},
		{"", ""},
	}

	for _, test := range tests {
		result := EchoRepeat(test.input)
		if result != test.expected {
			t.Errorf("EchoRepeat(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
