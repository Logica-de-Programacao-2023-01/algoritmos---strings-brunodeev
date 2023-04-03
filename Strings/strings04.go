package main

import (
	"fmt"
)

func main() {
	var first, second string

	fmt.Print("Digite a primeira string: ")
	fmt.Scan(&first)

	fmt.Print("Digite a segunda string: ")
	fmt.Scan(&second)

	if first == second {
		fmt.Println("Essas duas strings são EXATAMENTE iguais")
	} else {
		fmt.Println("Essas duas strings não são exatamente iguais")
	}
}
