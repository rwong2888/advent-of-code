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
	i := 0
	j := 0

	var a []string

	for scanner.Scan() {
		line := scanner.Text()

		if j == 0 {
			a = append(a, line)
		} else {
			a[i] += " " + line
		}
		if j == 2 {
			j = 0
			i++
		} else {
			j++
		}
	}

	for i := 0; i < len(a); i++ {
		x := strings.Split(a[i], " ")
		y := intersect.Simple(strings.Split(x[0], ""), strings.Split(x[1], ""))
		z := intersect.Simple(strings.Split(x[1], ""), strings.Split(x[2], ""))
		c := fmt.Sprintf("%s", intersect.Simple(y, z)[0])
		priority := strings.Index(alphabet, c) + 1
		fmt.Printf("%s: %d\n", c, priority)
		score += priority
	}
	fmt.Println(score)
}
