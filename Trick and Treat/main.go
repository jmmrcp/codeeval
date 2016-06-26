package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ajustar(s []string) (string, int) {
	st := strings.Trim(s[0], " ")
	n := strings.Trim(s[1], " ")
	m, _ := strconv.Atoi(n)
	return st, m
}

func dividirlinea(s string) {
	datos := strings.Split(s, ",")
	var caramelos, disfraz int
	for _, v := range datos[:] {
		s, n := ajustar(strings.Split(v, ":"))
		if s != "Houses" {
			disfraz += n
			caramelos += trick(s, n)
		} else {
			caramelos *= n
		}

	}
	fmt.Println(caramelos / disfraz)
}

func trick(s string, n int) int {
	switch s {
	case "Vampires":
		return (3 * n)
	case "Zombies":
		return (4 * n)
	case "Witches":
		return (5 * n)
	}
	return 0
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		test := scanner.Text()

		dividirlinea(test)
	}
}
