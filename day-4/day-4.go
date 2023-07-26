package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readFile(path string) (input string) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func extract_range(r string) (int, int) {
	nums := strings.Split(r, "-")
	left, err := strconv.Atoi(strings.TrimSpace(nums[0]))
	if err != nil {
		log.Fatal(err)
	}
	right, err := strconv.Atoi(strings.TrimSpace(nums[1]))
	if err != nil {
		log.Fatal(err)
	}

	return left, right

}

func is_range_within_range(min1, max1, min2, max2 int) bool {
	if (min1 >= min2) && (max1 <= max2) {
		return true
	}
	return false
}

func main() {
	content := readFile("./day-4/input.txt")
	sum := 0
	for _, pairs := range strings.Split(content, "\n") {
		if pairs == "" {
			continue
		}
		assignements := strings.Split(pairs, ",")
		left, right := extract_range(assignements[0])
		left2, right2 := extract_range(assignements[1])
		if is_range_within_range(left, right, left2, right2) || is_range_within_range(left2, right2, left, right) {
			sum += 1
		}
	}
	fmt.Printf("sum: %v", sum)

}
