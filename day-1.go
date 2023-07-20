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
	input := readFile("./day-1-input.txt")
	elfs := ParseElfs(input)
	max_elf := find_max_calory_elf(elfs)
	fmt.Printf("Max calory elf is %+v and holds %v", max_elf, max_elf.Calories())
}
