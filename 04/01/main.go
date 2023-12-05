package main

import (
	"embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	data, _ := f.ReadFile("input")
	splitByLine := strings.Split(string(data), "\n")

	sum := 0
	for _, l := range splitByLine {
		if l == "" {
			continue
		}
		card := 0
		_, l, _ = strings.Cut(l, ": ")
		before, after, _ := strings.Cut(l, " | ")
		re := regexp.MustCompile(`\d{1,2}`)
		cardNumbers := re.FindAllString(after, -1)

	myNumberLoop:
		for _, myNumber := range re.FindAllString(before, -1) {
			for _, cardNumber := range cardNumbers {
				if myNumber == cardNumber {
					fmt.Printf("incremented %s because it matched with %s \n", myNumber, cardNumber)
					card = increment(card)
					continue myNumberLoop
				}
			}
		}
		sum += card
	}
	fmt.Println(sum)
}

func increment(i int) int {
	j := 0
	if i == 0 {
		j = 1
	} else {
		j = i * 2
	}
	return j
}
