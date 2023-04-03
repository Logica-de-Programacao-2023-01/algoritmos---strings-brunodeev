package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "T e s t e"

	newWord := strings.ReplaceAll(word, " ", "")

	fmt.Printf("A string sem espaços será: %s", newWord)
}
