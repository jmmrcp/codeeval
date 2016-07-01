package main

import (
	"fmt"
	"math"
)

// esPrimo devuelve si un numero n es primo o no.
//
// esPrimo uint -> bool
// esPrimo(12)
// false
func esPrimo(n uint) bool {
	var i uint
	f := float64(n)
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n < 9 {
		return true
	}
	if n%3 == 0 {
		return false
	}
	r := (math.Sqrt(f + 0.0)) + 1

	for i = 5; float64(i) <= r; i += 6 {
		if n%i == 0 {
			return false
		}
		if n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// sumaPrimos devuelve la suma de los numeros primos hasta n.
//
// sumaPrimos uint -> uint64
// sumaPrimos(100)
// 1060
func sumaPrimos(n uint) uint64 {
	var contador, num uint = 0, 1
	var suma uint64
	for contador != n {
		if esPrimo(num) {
			suma += uint64(num)
			contador++
		}
		num++
	}
	return suma
}

func main() {
	suma := sumaPrimos(1000)
	fmt.Println(suma)
	//fmt.Printf("Suma de los numeros primos: %v.\n", suma) // 142913828922
}
