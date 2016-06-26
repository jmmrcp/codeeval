/*Sample code to read in test cases:*/

package main

import (
	"fmt"
	"strconv"
)
import "log"
import "bufio"
import "os"

func sumofdigits(numero string) (solucion int) {
	num, _ := strconv.Atoi(numero)
	for num > 0 {
		digito := num % 10
		solucion += digito
		num /= 10
	}
	return
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//'scanner.Text()' represents the test case, do something with it
		// fmt.Println(scanner.Text())
		dato := scanner.Text()
		fmt.Println(sumofdigits(dato))
	}
}
