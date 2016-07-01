package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fichero, _ := file.Stat()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // test := scanner.Text()
    }
    fmt.Println(fichero.Size())
}
