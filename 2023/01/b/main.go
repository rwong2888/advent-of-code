package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var line, match string
var integers []string
var d int

func main() {

	total := 0
	re, _ := regexp.Compile(`^(\d|one|two|three|four|five|six|seven|eight|nine)`)

	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		integers = make([]string, 0)
		line = fileScanner.Text()
		for i, _ := range line {
			match = re.FindString(line[i:])
			switch match {
			case "one", "1":
				match = "1"
			case "two", "2":
				match = "2"
			case "three", "3":
				match = "3"
			case "four", "4":
				match = "4"
			case "five", "5":
				match = "5"
			case "six", "6":
				match = "6"
			case "seven", "7":
				match = "7"
			case "eight", "8":
				match = "8"
			case "nine", "9":
				match = "9"
			}
			if match != "" {
				integers = append(integers, match)
			}
		}
		d, _ := strconv.Atoi((fmt.Sprintf("%s%s", integers[0], integers[len(integers)-1])))
		total += d
	}

	readFile.Close()

	fmt.Println(total)
}
