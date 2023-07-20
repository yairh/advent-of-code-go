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

func codeToMove(code string) int {
	codemap := map[string]int{
		"A": Rock,
		"B": Paper,
		"C": Scissor,
		"X": Rock,
		"Y": Paper,
		"Z": Scissor,
	}
	return codemap[code]
}

func codeToOutcome(code string) int {
	codemap := map[string]int{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}
	return codemap[code]
}

const (
	Rock = iota + 1
	Paper
	Scissor
)

const (
	Lose = 0
	Draw = 3
	Win  = 6
)

func shouldPlay(move, outcome int) int {
	if outcome == Draw {
		return move
	}
	if move == Rock {
		if outcome == Win {
			return Paper
		}
		return Scissor
	}
	if move == Paper {
		if outcome == Win {
			return Scissor
		}
		return Rock
	}
	// Scissor
	if outcome == Win {
		return Rock
	}
	return Paper
}

func play(move1, move2 int) int {
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
	if move1 == Paper {
		return Win
	}
	return Lose
}

type Player struct {
	id     int
	points int
}

func (p *Player) addPoints(points int) {
	p.points += points
}

func main() {
	guide := readFile("./day-2/guide.txt")

	player2 := Player{id: 2, points: 0}

	for _, g := range strings.Split(guide, "\n") {
		if g == "" {
			continue
		}
		moves := strings.Fields(g)
		// player2move := codeToMove(moves[1])
		outcome := codeToOutcome(moves[1])
		player2move := shouldPlay(codeToMove(moves[0]), outcome)
		// result := play(codeToMove(moves[0]), player2move)
		player2.addPoints(outcome)
		player2.addPoints(player2move)
	}
	fmt.Printf("Total points: %v\n", player2.points)

}
