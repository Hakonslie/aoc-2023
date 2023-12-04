package main

import (
	"embed"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	data, _ := f.ReadFile("test")
	splitByLine := strings.Split(string(data), "\n")

}
