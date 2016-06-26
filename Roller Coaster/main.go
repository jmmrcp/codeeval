/*Sample code to read in test cases:*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var escambiable = true

func convertir(r rune) rune {
	if esLetra(r) {
		if escambiable {
			escambiable = false
			r -= 'a' - 'A'
			return r
		}
		escambiable = true
		return r
	}
	return r
}

func esLetra(r rune) bool {
	if 'a' <= r && r <= 'z' {
		return true
	}
	return false
}

func formatoTexto(s string) string {
	s = strings.ToLower(s)
	formato := strings.Map(convertir, s)
	return formato
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		test := scanner.Text()
		escambiable = true
		fmt.Println(formatoTexto(test))
	}
}
