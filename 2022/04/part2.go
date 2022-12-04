package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func compare(a string, b string) bool {

	a1, _ := strconv.Atoi(strings.Split(a, "-")[0])
	a2, _ := strconv.Atoi(strings.Split(a, "-")[1])
	b1, _ := strconv.Atoi(strings.Split(b, "-")[0])
	b2, _ := strconv.Atoi(strings.Split(b, "-")[1])

	result := false
	if (a1 >= b1 && a1 <= b2) || (b1 >= a1 && b1 <= a2) {
		result = true
	}

	return result
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		a := strings.Split(line, ",")[0]
		b := strings.Split(line, ",")[1]
		r := compare(a, b)

		if r {
			count++
		}

		fmt.Printf("%s: %t, %d\n", line, r, count)
	}
	fmt.Println(count)
}
