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
	re2 := regexp.MustCompile(`[*]`)
	symbolsIndexes := re2.FindAllIndex(data, -1)

	sum := 0
	for _, symbols := range symbolsIndexes {
		mapOfMatchedNumbers := make(map[int]bool)
		for _, dir := range possibleDirections {
		numberLoop:
			for i, num := range numbersIndexes {
				if symbols[0]+dir >= num[0] && symbols[0]+dir < num[len(num)-1] {
					mapOfMatchedNumbers[i] = true
					break numberLoop
				}
			}
		}
		var numbers []int
		if len(mapOfMatchedNumbers) == 2 {
			for k, _ := range mapOfMatchedNumbers {
				no, _ := strconv.Atoi(string(data[numbersIndexes[k][0]:numbersIndexes[k][1]]))
				numbers = append(numbers, no)
			}
			sum += (numbers[0] * numbers[1])
		}
	}
	fmt.Println(sum)
}
