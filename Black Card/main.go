/*Sample code to read in test cases:*/
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
        entrada(test)
    }
}

func entrada(s string) {
    valores := strings.Split(s, "|")
    for i := 0; i < len(valores); i++ {
        valores[i] = strings.Trim(valores[i], " ")
    }
    jugadores := valores[0]
    numerostr := valores[1]
    numero, _ := strconv.Atoi(numerostr)
    listajugadores := strings.Split(jugadores, " ")
    blackcard(listajugadores, numero)
    // fmt.Printf("%#v - %d\n", listajugadores, numero)
}

func blackcard(s []string, n int) {
    numjugadores := len(s)
    for numjugadores != 1 {
        i := (n % numjugadores) - 1
        if i == -1 {
            s = s[:numjugadores-1]
        } else {
            s = s[:i+copy(s[i:], s[i+1:])]
        }
        numjugadores--
    }
    fmt.Println(s[0])
}
