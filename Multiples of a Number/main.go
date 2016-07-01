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
        dato := strings.Split(test, ",")
        num01, _ := strconv.Atoi(dato[0])
        num02, _ := strconv.Atoi(dato[1])
        var valor = 0
        for i := 0; num01 > valor; i++ {
            valor = num02 * i
        }
        fmt.Println(valor)
    }
}
