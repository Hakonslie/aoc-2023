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
	data, _ := f.ReadFile("test")
	splitByLine := strings.Split(string(data), "\n")

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
		}
		for i := len(line); i >= 0; i-- {
			letter := line[i-1 : i]
			if _, err := strconv.Atoi(letter); err == nil {
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
