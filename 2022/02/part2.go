package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func convert(s string) string {
	var result string
	switch s {
	case "A", "X":
		result = "Rock"
	case "B", "Y":
		result = "Paper"
	case "C", "Z":
		result = "Scissor"
	}
	return result
}

func winLoseDraw(a string, b string) string {
	result := "Draw"
	if a == b {
		return result
	}

	switch b {
	case "Rock":
		switch a {
		case "Paper":
			result = "Lose"
		case "Scissor":
			result = "Win"
		}
	case "Paper":
		switch a {
		case "Rock":
			result = "Win"
		case "Scissor":
			result = "Lose"
		}
	case "Scissor":
		switch a {
		case "Rock":
			result = "Lose"
		case "Paper":
			result = "Win"
		}
	}

	return result
}

func score(a string, b string) int {
	x := convert(a)
	y := convert(b)

	var score int
	switch y {
	case "Rock":
		score = 1
	case "Paper":
		score = 2
	case "Scissor":
		score = 3
	}

	result := winLoseDraw(x, y)
	switch result {
	case "Draw":
		score += 3
	case "Win":
		score += 6
	default:
	}

	fmt.Printf("%s %s %d\n", x, y, score)

	return score
}

func pick(a string, action string) string {
	var pick string
	switch action {
	case "Draw":
		pick = a
	case "Win":
		switch a {
		case "Rock":
			pick = "Paper"
		case "Paper":
			pick = "Scissor"
		case "Scissor":
			pick = "Rock"
		}
	case "Lose":
		switch a {
		case "Rock":
			pick = "Scissor"
		case "Paper":
			pick = "Rock"
		case "Scissor":
			pick = "Paper"
		}
	}
	fmt.Printf("Opponent: %s Action: %s Pick: %s\n", a, action, pick)
	return pick
}

func deconvert(s string) string {
	var x string
	switch s {
	case "Rock":
		x = "X"
	case "Paper":
		x = "Y"
	case "Scissor":
		x = "Z"
	}
	return x
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	points := 0

	for scanner.Scan() {
		line := scanner.Text()
		a := strings.Split(line, " ")[0]
		b := strings.Split(line, " ")[1]

		var action string

		switch b {
		case "X":
			action = "Lose"
		case "Y":
			action = "Draw"
		case "Z":
			action = "Win"
		}

		i := pick(convert(a), action)
		j := deconvert(i)

		points += score(a, j)
	}

	fmt.Printf("%d\n", points)
}
