package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var score = map[string]map[string]int{
	"A": {
		"X": 3,
		"Y": 4,
		"Z": 8,
	},
	"B": {
		"X": 1,
		"Y": 5,
		"Z": 9,
	},
	"C": {
		"X": 2,
		"Y": 6,
		"Z": 7,
	},
}

type RPS struct {
	opponentHand string
	hand         string
}

func (rps RPS) duel() int {
	return score[rps.opponentHand][rps.hand]
}

func main() {
	var rps RPS

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var score int

	for scanner.Scan() {
		temp := scanner.Text()
		tempSlice := strings.Split(temp, " ")
		rps.opponentHand, rps.hand = tempSlice[0], tempSlice[1]
		score += rps.duel()
	}
	fmt.Println(score)
}
