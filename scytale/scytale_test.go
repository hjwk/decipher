package scytale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncipher(t *testing.T) {
	testCases := map[string]struct {
		plaintext          string
		keyLength          int
		expectedCiphertext string
		expectedError      bool
	}{
		"Key of 0": {
			"test",
			0,
			"",
			true,
		},
		"Key of 1": {
			"ABCDEF",
			1,
			"ABCDEF",
			false,
		},
		"Key of -1": {
			"ABCDEF",
			-1,
			"",
			true,
		},
		"Key of 3": {
			"HelloWorld",
			3,
			"HlodeorZlWlZ",
			false,
		},
		"Key of 25": {
			"ABCDEF",
			25,
			"",
			true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := Encipher(tc.plaintext, tc.keyLength)
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tc.expectedCiphertext, got)
			}
		})
	}
}
