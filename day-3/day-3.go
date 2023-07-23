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
	var sack_list []string
	for _, sack := range strings.Split(sacks, "\n") {
		sack_list = append(sack_list, sack)
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

	type void struct{}
	sum2 := 0
	for i := 0; i < len(sack_list); i += 3 {
		charset := make(map[rune]void)
		for _, item := range sack_list[i] {
			if _, ok := charset[item]; ok {
				continue
			}
			if strings.ContainsRune(sack_list[i+1], item) && strings.ContainsRune(sack_list[i+2], item) {
				sum2 += getPriority(item)
				charset[item] = void{}
			}
		}
	}
	fmt.Printf("priority sum: %v", sum)
	fmt.Printf("badge priority sum: %v", sum2)
}
