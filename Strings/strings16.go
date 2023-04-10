package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Print("\nDiga uma string em camelCase (sem espaços): ")
	fmt.Scan(&s)

	camelcase := false
	for _, c := range s {
		if unicode.IsUpper(c) {
			camelcase = true
			break
		}
	}
	if camelcase {
		fmt.Println("\nA string está em CamelCase.")
	} else {
		fmt.Println("\nA string não está em CamelCase.")
	}

	numerodepalavras := 1
	for _, c := range s {
		if unicode.IsUpper(c) {
			numerodepalavras++
		}
	}
	fmt.Printf("\nA string contém %d palavras.\n", numerodepalavras)
}
