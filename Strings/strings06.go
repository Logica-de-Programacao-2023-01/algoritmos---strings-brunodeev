package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Digite uma frase para contarmos as palavras: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	line := scanner.Text()

	words := strings.Fields(line)

	count := len(words)

	fmt.Printf("Sua frase conta com %d palavras!", count)
}
