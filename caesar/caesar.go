package caesar

import (
	"math"
	"strings"
)

var (
	freqs_eng = []float32{
		'a': 8.12,
		'b': 1.49,
		'c': 2.71,
		'd': 4.32,
		'e': 12.02,
		'f': 2.3,
		'g': 2.03,
		'h': 5.92,
		'i': 7.31,
		'j': 0.10,
		'k': 0.69,
		'l': 3.98,
		'm': 2.61,
		'n': 6.95,
		'o': 7.68,
		'p': 1.82,
		'q': 0.11,
		'r': 6.02,
		's': 6.28,
		't': 9.10,
		'u': 2.88,
		'v': 1.11,
		'w': 2.09,
		'x': 0.17,
		'y': 2.11,
		'z': 0.07}

	freqs_fr = []float32{}

	freqs_init = []float32{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0}
)

// Encipher encrypts plaintext with a given shift.
func Encipher(plaintext string, shift int) string {
	shift = (shift%26 + 26) % 26 // [0, 25]
	bytes := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		c := plaintext[i]
		var a int
		switch {
		case 'a' <= c && c <= 'z':
			a = 'a'
		case 'A' <= c && c <= 'Z':
			a = 'A'
		default:
			bytes[i] = c
			continue
		}
		bytes[i] = byte(a + ((int(c)-a)+shift)%26)
	}

	return string(bytes)
}

// Decipher attempts to guess the shift of a ciphertext.
func Decipher(in, lang string) (int, string) {
	var freqs_ref []float32
	switch lang {
	case "eng":
		freqs_ref = freqs_eng
	case "fr":
		freqs_ref = freqs_fr
	default:
		freqs_ref = freqs_eng
	}

	in = strings.ToLower(in)
	freqs := countFrequencies(in)

	min := float32(math.MaxFloat32)
	shift := 0
	for i := 0; i < 26; i++ {
		err := computeErrorSquared(freqs_ref, freqs, i)
		if err < min {
			min = err
			shift = i
		}
	}

	return shift, Encipher(in, -shift)
}

func countFrequencies(in string) []float32 {
	freqs := freqs_init

	var inputChars float32
	for _, r := range in {
		if r != ' ' && r != ',' && r != '.' {
			freqs[r]++
			inputChars++
		}
	}

	for i := range freqs {
		freqs[i] = freqs[i] / inputChars
	}

	return freqs
}

func computeErrorSquared(ref, input []float32, shift int) float32 {
	var errorSquared float32
	for i := 'a'; i <= 'z'; i++ {
		j := i + rune(shift)
		if j > 'z' {
			j = 'a' + j - ('z' + 1)
		}
		errorSquared += (ref[i] - input[j]) * (ref[i] - input[j])
	}

	return errorSquared
}
