package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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

func ParseElfs(content string) (elfs []elf){

	for idx, l := range strings.Split(content, "\n\n") {
		str_items :=  strings.Split(l, "\n")
		items := make([]int, 0, len(str_items)) 
		for _, i := range strings.Split(l, "\n"){
			if i == "" {
				continue
			}
			cal, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			items = append(items, cal)
		}
		e := elf{name:string(rune(idx)), items:items}
		elfs = append(elfs, e)
	}
	return elfs
	
}

func sumIntArray(arr []int) (s int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum

}

type elf struct {
	name  string
	items []int
}

func (e elf) Calories() (c int) {
	return sumIntArray(e.items)
}

func sortElfsDecreasing(elfs []elf) {
	sort.Slice(elfs[:], func(i, j int ) bool {return elfs[i].Calories() > elfs[j].Calories() })
 
}

func find_max_calory_elf(elfs []elf) (e elf) {

	max_elf := elfs[0]

	for _, e := range elfs {
		if e.Calories() > max_elf.Calories() {
			max_elf = e
		}

	}
	return max_elf

}

func main() {
	input := readFile("./day-1/day-1-input.txt")
	elfs := ParseElfs(input)
	sortElfsDecreasing(elfs)
	fmt.Printf("Max calory elf is %+v and holds %v\n", elfs[0], elfs[0].Calories())
	top_three := elfs[0].Calories() + elfs[1].Calories() + elfs[2].Calories()

	fmt.Printf("Top three %v", top_three)
}
