package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
		integers = re.FindAllString(line, -1)
		d, _ := strconv.Atoi((fmt.Sprintf("%s%s", integers[0], integers[len(integers)-1])))
		total += d
	}

	readFile.Close()

	fmt.Println(total)
}
