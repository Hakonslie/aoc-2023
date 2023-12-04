package main

import (
	"embed"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	data, _ := f.ReadFile("input")
	splitByLine := strings.Split(string(data), "\n")

}
