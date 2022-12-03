package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/juliangruber/go-intersect"
)

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabet += strings.ToUpper(alphabet)

	scanner := bufio.NewScanner(f)

	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		a := line[0 : length/2]
		b := line[length/2:]

		c := fmt.Sprintf("%s", intersect.Simple(strings.Split(a, ""), strings.Split(b, ""))[0])
		priority := strings.Index(alphabet, c) + 1
		fmt.Printf("%s: %d\n", c, priority)
		score += priority
	}
	fmt.Println(score)
}
