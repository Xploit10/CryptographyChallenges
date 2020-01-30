package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func Xor(stringBytes []byte, keyBytes []byte) ([]byte, error) {

	if len(stringBytes) != len(keyBytes) {
		return nil, errors.New("Bytes slice are wrong sizes.")
	}

	store := make([]byte, len(stringBytes))

	for i := range stringBytes {
		store[i] = stringBytes[i] ^ keyBytes[i]
	}

	return store, nil
}

func main() {
	// Hex strings
	s := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	// to bytes
	stringBytes, _ := hex.DecodeString(s)
	keyBytes, _ := hex.DecodeString(key)

	// return bytes xor
	result, _ := Xor(stringBytes, keyBytes)

	// Encode to Hex
	finalString := hex.EncodeToString(result)
	fmt.Println(finalString)

}
