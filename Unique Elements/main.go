package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        test := scanner.Text()
        evaluar(test)
    }
}

func evaluar(s string) {
    dato := strings.Split(s, ",")
    c := len(dato)
    numeros := make([]int, c)
    for i := 0; i < c; i++ {
        num, _ := strconv.Atoi(dato[i])
        numeros[i] = num
    }
    unique(numeros)
}

func unique(n []int) {
    longitud := len(n)
    anterior := n[0]
    actual := 0
    lista := strconv.Itoa(anterior)
    for i := 1; i < longitud; i++ {
        actual = n[i]
        if anterior != actual {
            lista = lista + "," + strconv.Itoa(actual)
        }
        anterior = n[i]
    }
    fmt.Println(lista)
}
