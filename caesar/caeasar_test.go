package caesar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncipher(t *testing.T) {
	testCases := map[string]struct {
		in       string
		shift    int
		expected string
	}{
		"Shift of 0": {
			"test",
			0,
			"test",
		},
		"Shift of 1": {
			"ABCDEF",
			1,
			"BCDEFG",
		},
		"Shift of -1": {
			"ABCDEF",
			-1,
			"ZABCDE",
		},
		"Shift of 25": {
			"ABCDEF",
			25,
			"ZABCDE",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := Encipher(tc.in, tc.shift)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestDecipher(t *testing.T) {
	testCases := map[string]struct {
		in            string
		lang          string
		expectedShift int
		expectedText  string
	}{
		"Not encoded eng": {
			"this string is not encoded, and is quite short",
			"eng",
			0,
			"this string is not encoded, and is quite short",
		},
		"Not encoded fr": {
			"ce message n'est pas encode",
			"fr",
			0,
			"ce message n'est pas encode",
		},
		"Basic eng": {
			"wklv lv d whvw",
			"eng",
			3,
			"this is a test",
		},
		"Basic fr": {
			"fhfl hvw xq whvw",
			"fr",
			3,
			"ceci est un test",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			shift, msg := Decipher(tc.in, tc.lang)
			assert.Equal(t, tc.expectedShift, shift)
			assert.Equal(t, tc.expectedText, msg)
		})
	}
}
