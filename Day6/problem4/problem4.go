package main

import (
	"fmt"
	"strings"
)

func encode(sentence string) string {
	var key int
	key = 10

	abjad := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sentence = strings.ToUpper(sentence)
	hasilEncode := ""
	for _, karakter := range sentence {
		if strings.Contains(abjad, string(karakter)) {
			indexLama := strings.Index(abjad, string(karakter))
			indexBaru := (indexLama + key) % len(abjad)
			abjadBaru := string(abjad[indexBaru])
			hasilEncode += abjadBaru
		} else {
			hasilEncode += " "
		}
	}
	return hasilEncode
}

func main() {
	fmt.Println(encode("SEPULSA OKE"))
	fmt.Println(encode("ALTERRA ACADEMY"))
	fmt.Println(encode("INDONESIA"))
	fmt.Println(encode("GOLANG"))
	fmt.Println(encode("PROGRAMMER"))
}
