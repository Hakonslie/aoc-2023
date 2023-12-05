package main

import (
	"embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed *
var f embed.FS

func main() {
	data, _ := f.ReadFile("input")

	rowLength := 0
	for _, col := range data {
		rowLength++
		if col == byte('\n') {
			break
		}
	}
	possibleDirections := []int{
		-(rowLength - 1), -(rowLength), -(rowLength + 1),
		-1, +1,
		rowLength - 1, rowLength, rowLength + 1,
	}

	re := regexp.MustCompile(`\d+`)
	numbersIndexes := re.FindAllIndex(data, -1)
	re2 := regexp.MustCompile(`[^0-9.\n]`)
	symbolsIndexes := re2.FindAllIndex(data, -1)

	sum := 0
	for _, numbers := range numbersIndexes {
		match := false
	symbolsLoop:
		for _, symbols := range symbolsIndexes {
			for i := numbers[0]; i < numbers[1]; i++ {
				for _, direction := range possibleDirections {
					fmt.Printf("comparing %d and %d \n", symbols[0]+direction, i)
					if symbols[0]+direction == i {
						match = true
						break symbolsLoop
					}
				}
			}
		}
		if match == true {
			nr, _ := strconv.Atoi(string(data[numbers[0]:numbers[1]]))
			sum += nr
		}
	}

	fmt.Println(sum)
}
