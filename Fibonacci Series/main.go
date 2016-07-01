package main

import (
    "bufio"
    "fmt"
    "log"
)

import "os"

func main() {

    var n int

    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //test := scanner.Text()
        fmt.Sscanf(scanner.Text(), "%d", &n)
        fibonacci(n)
    }
}

func fibonacci(n int) {

    numantes := 0
    numahora := 1

    for i := 0; i < n; i++ {
        temp := numantes
        numantes = numahora
        numahora = numantes + temp
    }
    fmt.Println(numantes)
}
