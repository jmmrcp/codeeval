package main

import (
    "bufio"
    "fmt"
    "log"
)

import "os"

func main() {
    var n, valor int
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d", &n)
        valor += n
    }
    fmt.Println(valor)
}
