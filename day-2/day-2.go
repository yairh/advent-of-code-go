package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readFile(path string) (input string) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}


func codeToMove(code string) Move {
	codemap := map[string]Move{
		"A":Rock,
		"B":Paper,
		"C":Scissor,
		"X":Rock,
		"Y":Paper,
		"Z":Scissor,
	}
	return codemap[code]
}

type Move int 

const (
	Rock Move = iota + 1  	
	Paper  
	Scissor 
)

const (
	Lose = 0
	Draw = 3
	Win = 6
)


func winner(move1 Move, move2 Move ) int {
	if move1 == move2 {
		return Draw
	}

	if move2 == Rock {
		if move1 == Scissor {
			return Win
		}
			return Lose
	}
	if move2 == Paper {
		if move1 == Rock {
			return Win
		}
			return Lose
	}
	
	if move2 == Scissor {
		if move1 == Paper {
			return Win
		}
		return Lose
	}
	return 0 // should not happen
}

type Player struct {
	id int
	points int
}

func (p *Player) addPoints(points int) {
	p.points += points
}

func main() {
	guide := readFile("./day-2/guide.txt")
	
	// player1 := Player{id:1, points:0}
	player2 := Player{id:2, points:0}

	for _, g := range strings.Split(guide, "\n") {
		if g == "" {
			continue }
		moves := strings.Fields(g)
		player2move := codeToMove(moves[1])
		result := winner(codeToMove(moves[0]), player2move)
		player2.addPoints(result)
		player2.addPoints(int(player2move))
	}
		fmt.Printf("Total points: %v\n", player2.points)
	
}
