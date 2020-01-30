package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

type Keysize struct {
	Length int
}

func HammingDistance(a, b []byte) int {
	if len(a) != len(b) {
		log.Fatalln(errors.New("String Length Error Computing HammingDistance"))
	}
	var distance int
	for i := range a {
		if a[i] != b[i] {

		}
	}

}

func main() {

	b, err := ioutil.ReadFile("encryptedFile")
	if err != nil {
		log.Fatal(err)
	}

	// Cipher text as []bytes
	cipherText, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cipherText)

}
