package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func readFile(path string) (input string) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func getPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item) - (65 - 27)
	}
	return int(item) - (97 - 1)

}

func main() {
	sacks := readFile("./day-3/input.txt")
	sum := 0

	for _, sack := range strings.Split(sacks, "\n") {
		halfsack := sack[:len(sack)/2]
		halfsack2 := sack[len(sack)/2:]
		charset := make(map[rune]struct{})
		for _, item := range halfsack {
			if _, ok := charset[item]; ok {
				continue
			}
			if strings.ContainsRune(halfsack2, item) {
				sum += getPriority(item)
				charset[item] = struct{}{}
			}
		}
	}
	fmt.Printf("priority sum: %v", sum)
}
