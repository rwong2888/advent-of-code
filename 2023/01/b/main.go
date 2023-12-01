package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var line string
var integers []string
var d int

func main() {

	total := 0
	re := regexp.MustCompile("[0-9]")

	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line = fileScanner.Text()
		fmt.Println(line)
		line = strings.ReplaceAll(line, "one", "1")
		line = strings.ReplaceAll(line, "two", "2")
		line = strings.ReplaceAll(line, "three", "3")
		line = strings.ReplaceAll(line, "four", "4")
		line = strings.ReplaceAll(line, "five", "5")
		line = strings.ReplaceAll(line, "six", "6")
		line = strings.ReplaceAll(line, "seven", "7")
		line = strings.ReplaceAll(line, "eight", "8")
		line = strings.ReplaceAll(line, "nine", "9")
		integers = re.FindAllString(line, -1)
		d, _ := strconv.Atoi((fmt.Sprintf("%s%s", integers[0], integers[len(integers)-1])))
		total += d
		fmt.Println(line, integers, d, total)
	}

	readFile.Close()

	fmt.Println(total)
}
