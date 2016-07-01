/*Sample code to read in test cases:*/
package main

import (
    "bufio"
    "fmt"
    "log"
)

import "os"

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        test := scanner.Text()
        fmt.Sscanf(scanner.Text(), "%d", &n)
    }
}
