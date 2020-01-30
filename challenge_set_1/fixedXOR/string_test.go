package main

import (
	"encoding/hex"
	"testing"
)

func TestXOR(t *testing.T) {
	s := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	// to bytes
	stringBytes, _ := hex.DecodeString(s)
	keyBytes, _ := hex.DecodeString(key)

	// return bytes xor
	result, _ := Xor(stringBytes, keyBytes)

	// Encode to Hex
	finalString := hex.EncodeToString(result)
	if finalString != "746865206b696420646f6e277420706c6179" {
		t.Errorf("XOR Error.")
	}
}
