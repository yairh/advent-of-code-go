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

type Range struct {
	min int
	max int
}

func (r Range) within(r2 Range) bool {
	if (r.min >= r2.min) && (r.max <= r2.max) {
		return true
	}
	return false
}

func (r Range) overlaps(r2 Range) bool {
	if r.min <= r2.min && r.max >= r2.min {
		return true
	}
	if r2.min <= r.min && r2.max >= r.min {
		return true
	}
	return false
}

func extract_range(r string) Range {
	nums := strings.Split(r, "-")
	left, err := strconv.Atoi(strings.TrimSpace(nums[0]))
	if err != nil {
		log.Fatal(err)
	}
	right, err := strconv.Atoi(strings.TrimSpace(nums[1]))
	if err != nil {
		log.Fatal(err)
	}

	return Range{min: left, max: right}
}

func main() {
	content := readFile("./day-4/input.txt")
	contained := 0
	overlaped := 0
	for _, pairs := range strings.Split(content, "\n") {
		if pairs == "" {
			continue
		}
		assignements := strings.Split(pairs, ",")
		range1 := extract_range(assignements[0])
		range2 := extract_range(assignements[1])
		if range1.within(range2) || range2.within(range1) {
			contained += 1
		}
		if range1.overlaps(range2) || range2.overlaps(range1) {
			overlaped += 1
		}
	}
	fmt.Printf("contained: %v\n", contained)
	fmt.Printf("overlaped: %v", overlaped)
}
