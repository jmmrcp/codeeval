package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func simple(s string) {
	lista := strings.Split(s, " ")
	list := make([]float64, 0)
	for _, v := range lista {
		i, _ := strconv.ParseFloat(v, 32)
		list = append(list, i)
	}
	sort.Float64s(list)
	for _, v := range list {
		fmt.Printf("%.3f ", v)
	}
	fmt.Println()
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

		//fmt.Println(test)
		simple(test)
	}
}
