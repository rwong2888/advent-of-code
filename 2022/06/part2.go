package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	fmt.Println(line)

	var i int
	var j int
	var s string
	x := 14
	for i = 0; i < len(line)-x; i++ {
		s = line[i : i+x]
		fmt.Println(s)
		for j = 0; j < len(s); j++ {
			if strings.Count(s, string(s[j])) != 1 {
				break
			}
		}
		if j == len(s) {
			break
		}
	}
	fmt.Println(i + x)
}
