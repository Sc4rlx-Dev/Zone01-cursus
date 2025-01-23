package goreloded

import (
	"testing"
)

func TestProcessText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			"Don not be sad ,because sad backwards is das ",
			"Don not be sad, because sad backwards is das.",
		},
		{
			"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
	}

	for _, tt := range tests {
		result := ProcessText(tt.input)
		if result != tt.expected {
			t.Errorf("For input: %q\nExpected: %q\nGot: %q", tt.input, tt.expected, result)
		}
	}
}
