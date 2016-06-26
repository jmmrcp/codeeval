/*Sample code to read in test cases:*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func stringMask(s string) string {
	var res string
	dato := strings.Split(s, " ")
	st := dato[0]
	n := dato[1]
	for i, v := range st {
		letra := mask(v, n[i])
		res += letra
	}
	return res
}

func mask(r rune, u uint8) string {
	if u == 49 {
		r -= 'a' - 'A'
		return string(r)
	}
	return string(r)
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
		fmt.Println(stringMask(test))
	}
}
