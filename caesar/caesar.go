package caesar

import (
	"math"
	"strings"

	"github.com/hjwk/decipher/common"
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

// Decipher attempts to guess the shift of a caesar ciphertext.
func Decipher(in, lang string) (int, string) {
	var freqsRef []float32
	switch lang {
	case "eng":
		freqsRef = common.FreqsEng
	case "fr":
		freqsRef = common.FreqsFr
	default:
		freqsRef = common.FreqsEng
	}

	in = strings.ToLower(in)
	freqs := countFrequencies(in)

	min := float32(math.MaxFloat32)
	shift := 0
	for i := 0; i < 26; i++ {
		err := computeErrorSquared(freqsRef, freqs, i)
		if err < min {
			min = err
			shift = i
		}
	}

	return shift, Encipher(in, -shift)
}

func countFrequencies(in string) []float32 {
	freqs := common.FreqsInit

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
