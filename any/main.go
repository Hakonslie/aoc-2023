package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	input, _ := f.ReadFile("test")
	lines := strings.Split(string(input), "\n")
	fmt.Println(p1(lines))
}

func p1(lines []string) int {
	ans := 0

	return ans
}

/*
func p2(lines []string) int {
	ans := 0

	return ans
}
*/
