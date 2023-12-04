package main

import (
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	data, _ := f.ReadFile("input")
	splitByLine := strings.Split(string(data), "\n")

	mapOfConversion := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var totalNumber int
	for _, line := range splitByLine {
		if line == "" {
			break
		}
		var firstNumber string
		var lastNumber string
		for i := 0; i <= len(line); i++ {
			letter := line[i : i+1]
			if _, err := strconv.Atoi(letter); err == nil {
				firstNumber = letter
				break
			}
			if len(line)-i < 4 {
				continue
			}
			if letter, ok := mapOfConversion[line[i:i+3]]; ok {
				firstNumber = letter
				break
			}
			if len(line)-i < 5 {
				continue
			}
			if letter, ok := mapOfConversion[line[i:i+4]]; ok {
				firstNumber = letter
				break
			}
			if len(line)-i < 6 {
				continue
			}
			if letter, ok := mapOfConversion[line[i:i+5]]; ok {
				firstNumber = letter
				break
			}
		}
		for i := len(line); i >= 0; i-- {
			letter := line[i-1 : i]
			if _, err := strconv.Atoi(letter); err == nil {
				lastNumber = letter
				break
			}
			if i < 3 {
				continue
			}
			if letter, ok := mapOfConversion[line[i-3:i]]; ok {
				lastNumber = letter
				break
			}
			if i < 4 {
				continue
			}
			if letter, ok := mapOfConversion[line[i-4:i]]; ok {
				lastNumber = letter
				break
			}
			if i < 5 {
				continue
			}
			if letter, ok := mapOfConversion[line[i-5:i]]; ok {
				lastNumber = letter
				break
			}
		}

		no, err := strconv.Atoi(firstNumber + lastNumber)
		if err != nil {
			log.Panic("wtf")
			return
		}

		totalNumber += no
		fmt.Printf("%d added \n", no)

	}
	fmt.Println(totalNumber)
}
