package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func HextoBase64(value string) string {
	data, err := hex.DecodeString(value)
	if err != nil {
		log.Fatal(err)
	}
	encodedString := base64.StdEncoding.EncodeToString([]byte(data))

	return encodedString
}

func main() {
	const s = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	result := HextoBase64(s)
	fmt.Println(result)
}
