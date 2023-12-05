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
	cards := strings.Split(string(data), "\n")

	extraCards := make(map[int]int)
	sum := 0
	for i := 0; i < len(cards); i++ {
		l := cards[i]
		if l == "" {
			continue
		}
		_, l, _ = strings.Cut(l, ": ")
		before, after, _ := strings.Cut(l, " | ")
		re := regexp.MustCompile(`\d{1,2}`)
		cardNumbers := re.FindAllString(after, -1)

		wins := 0
	myNumberLoop:
		for _, myNumber := range re.FindAllString(before, -1) {
			for _, cardNumber := range cardNumbers {
				if myNumber == cardNumber {
					wins++
					continue myNumberLoop
				}
			}
		}
		for j := 1; j <= wins; j++ {
			extraCards[i+j] = extraCards[i+j] + 1
		}
		sum++

		if extraCards[i] > 0 {
			extraCards[i] = extraCards[i] - 1
			i--
		}
	}
	fmt.Println(sum)
}
