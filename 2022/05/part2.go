package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	index := 1
	var stackSize int
	m := make(map[int]string)
	for scanner.Scan() {
		line := scanner.Text()
		if index == 1 {
			stackSize = (len(line) + 1) / 4
		}
		x := 1
		for i := 1; i <= stackSize; i++ {
			c := strings.Split(line, "")[x]
			if unicode.IsUpper([]rune(c)[0]) {
				m[i] += c
			}
			x += 4
		}
		index++
	}

	f, err = os.Open("./instructions")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner = bufio.NewScanner(f)

	fmt.Println(m)
	for scanner.Scan() {
		line := scanner.Text()
		instruction := strings.Split(line, " ")
		count, _ := strconv.Atoi(instruction[1])
		src, _ := strconv.Atoi(instruction[3])
		dest, _ := strconv.Atoi(instruction[5])

		s1 := m[src][:count]
		s2 := m[src][count:]

		m[dest] = s1 + m[dest]
		m[src] = s2
		fmt.Println(m)
	}
	fmt.Println(m)
	fmt.Println()
	r := ""
	for i := 1; i <= len(m); i++ {
		s, _ := m[i]
		r += s[:1]
	}
	fmt.Println(r)
}
