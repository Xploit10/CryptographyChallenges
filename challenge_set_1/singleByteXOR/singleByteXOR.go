package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

type Data struct {
	Message []byte
	Score   float64
	Key     int
}

type Store struct {
	Items []Data
}

func (s *Store) AddItem(item Data) []Data {
	s.Items = append(s.Items, item)
	return s.Items
}

func single_char_xor(input_bytes []byte, char_value int) []byte {
	bytes := make([]byte, len(input_bytes))
	for i := range input_bytes {
		bytes[i] = input_bytes[i] ^ byte(char_value)
	}
	return bytes
}

func sum(input []float64) float64 {
	sum := 0.0
	for i := range input {
		sum += input[i]
	}
	return sum
}

func get_english_score(message []byte) float64 {
	character_frequencies := map[uint8]float64{
		'a': .08167, 'b': .01492, 'c': .02782, 'd': .04253,
		'e': .12702, 'f': .02228, 'g': .02015, 'h': .06094,
		'i': .06094, 'j': .00153, 'k': .00772, 'l': .04025,
		'm': .02406, 'n': .06749, 'o': .07507, 'p': .01929,
		'q': .00095, 'r': .05987, 's': .06327, 't': .09056,
		'u': .02758, 'v': .00978, 'w': .02360, 'x': .00150,
		'y': .01974, 'z': .00074, ' ': .13000,
	}

	scores := make([]float64, len(message))
	var score float64
	for _, char := range bytes.ToLower([]byte(message)) {
		value, ok := character_frequencies[char]
		if ok {
			score = value
		} else {
			score = 0
		}
		scores = append(scores, score)
	}
	return sum(scores)

}

func main() {
	const s = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	s_byte, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	store := Store{}

	for i := 0; i <= 256; i++ {
		message := single_char_xor(s_byte, i)
		frequency := get_english_score(message)
		data := Data{
			Message: message,
			Score:   frequency,
			Key:     i,
		}

		store.AddItem(data)

	}
	var value string
	var scorer float64

	var maxScore float64
	var maxS string

	for i := range store.Items {
		value = string(store.Items[i].Message)
		scorer = store.Items[i].Score
		if scorer > maxScore {
			maxS = value
			maxScore = scorer
		}
	}
	fmt.Println(maxScore, maxS)

}
