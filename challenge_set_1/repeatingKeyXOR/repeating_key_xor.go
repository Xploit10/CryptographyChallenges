package main

import (
	"encoding/hex"
	"fmt"
)

func repeatKeyXOR(text []byte, key []byte) (encrypted []byte) {
	counter := 0

	encryptedText := make([]byte, len(text))

	for i := 0; i < len(text); i++ {
		if counter == len(key) {
			counter = 0
			encryptedText[i] = text[i] ^ key[counter]
			counter += 1
		} else if counter < len(key) {
			encryptedText[i] = text[i] ^ key[counter]
			counter += 1
		}

	}
	return encryptedText
}

func main() {
	text := []byte("Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal")
	key := []byte("ICE")

	encryptedText := repeatKeyXOR(text, key)

	fmt.Println(hex.EncodeToString(encryptedText))
}

// Answer
// 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272
// a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f

// Result = True
