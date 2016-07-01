package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
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
        if len(test) != 0 {
            datos := strings.Split(test, " ")
            for i := len(datos) - 1; i >= 0; i-- {
                fmt.Printf("%s", datos[i])
                if i != 0 {
                    fmt.Printf(" ")
                }
            }
            fmt.Println()
        }
    }
}
