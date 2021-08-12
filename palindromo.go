package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	text = strings.ToLower(text)
	text = strings.Replace(text, " ", "", -1)
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}

func main() {
	isPalindromo(" Anita  lava la tina ")
	isPalindromo(" Ansdfs ")
	isPalindromo("amA ")
	isPalindromo("hola")
}
