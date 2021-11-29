package scytale

import (
	"errors"
	"strings"
)

// Encipher enciphers a plaintext with the given keylength (representing the diameter of the scytale).
func Encipher(plaintext string, keyLength int) (string, error) {
	if keyLength <= 0 {
		return "", errors.New("keyLength cannot be less or equal to 0")
	}

	length := len(plaintext)

	if keyLength > length {
		return "", errors.New("keyLength cannot be longer than plaintext")
	}

	rows := length / keyLength
	modulo := length % keyLength
	padding := 0
	if modulo > 0 {
		rows++
		padding = keyLength - (length % keyLength)
		plaintext += strings.Repeat("Z", padding)
	}

	ciphertext := make([]byte, length+padding)

	i := 0
	for col := 0; col < keyLength; col++ {
		for row := 0; row < rows; row++ {
			characterToAdd := col + row*keyLength
			ciphertext[i] = byte(plaintext[characterToAdd])
			i++
		}
	}

	return string(ciphertext), nil
}
