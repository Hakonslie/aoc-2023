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

	addedUp := 0

	for i := 0; i < len(splitByLine)-1; i++ {
		handful, _ := strings.CutPrefix(splitByLine[i], ":")

		powerOfCubes := 1
		for _, color := range []string{"blue", "green", "red"} {
			largest := 0
			re := regexp.MustCompile(fmt.Sprintf("\\d{1,2}\\s%s", color))
			findings := re.FindAllString(handful, -1)
			for _, found := range findings {
				number := strings.Split(found, " ")[0]
				n, err := strconv.Atoi(number)
				if err != nil {
					log.Panic(err)
				}
				fmt.Printf("Saw %d of %s \n", n, color)
				if n > largest {
					largest = n
				}
			}
			powerOfCubes = powerOfCubes * largest
		}
		addedUp += powerOfCubes
	}
	fmt.Println(addedUp)

}
