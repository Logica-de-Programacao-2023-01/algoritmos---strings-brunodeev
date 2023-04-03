package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "hello"

	newWord := strings.ToUpper(word)

	fmt.Printf("\nA frase %s inteiramente maiúscula é %s\n", word, newWord)
}
