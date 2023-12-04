package main

import (
	"embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	data, _ := f.ReadFile("input")
	splitByLine := strings.Split(string(data), "\n")

	allowed := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	addedUp := 0

gameLoop:
	for i := 0; i < len(splitByLine)-1; i++ {
		gameID := i + 1
		handful, _ := strings.CutPrefix(splitByLine[i], ":")

		for _, color := range []string{"blue", "green", "red"} {
			re := regexp.MustCompile(fmt.Sprintf("\\d{2,3}\\s%s", color))
			findings := re.FindAllString(handful, -1)
			if len(findings) == 0 {
				continue
			}
			for _, found := range findings {
				number := strings.Split(found, " ")[0]
				n, err := strconv.Atoi(number)
				if err != nil {
					log.Panic(err)
				}
				fmt.Printf("Saw %d of %s \n", n, color)
				if n > allowed[color] {
					continue gameLoop
				}
			}
		}
		addedUp += gameID
	}
	fmt.Println(addedUp)

}
