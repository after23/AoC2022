package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Head Point
type Tail Point

var visitedPos map[string]int

func (head *Head) NewHead() {
	head.x, head.y = 0, 0
}

func (tail *Tail) NewTail() {
	tail.x, tail.y = 0, 0
}

func (tail Tail) isNear(head Head) bool {
	deltaX := head.x - tail.x
	deltaY := head.y - tail.y

	if moreThanOne(deltaX) || moreThanOne(deltaY) {
		return false
	}
	return true
}

func (head *Head) step(direction string, sign int) {
	if direction == "D" || direction == "U" {
		head.y += sign
	} else {
		head.x += sign
	}
}

func (tail *Tail) follow(head Head) {
	if tail.isNear(head) {
		return
	}

	signX := 1
	signY := 1

	deltaX := head.x - tail.x
	deltaY := head.y - tail.y

	if deltaX < -1 {
		signX = -1
	}

	if deltaY < -1 {
		signY = -1
	}

	if moreThanOne(deltaX) {
		tail.y = head.y
		tail.x += signX
		position := fmt.Sprintf("%d,%d", tail.x, tail.y)
		visitedPos[position] = 1
		return
	}

	if moreThanOne(deltaY) {
		tail.y += signY
		tail.x = head.x
		position := fmt.Sprintf("%d,%d", tail.x, tail.y)
		visitedPos[position] = 1
		return
	}
}

func moreThanOne(n int) bool {
	if n < -1 || n > 1 {
		return true
	}
	return false
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func move(direction string, steps int, head *Head, tail *Tail) {
	sign := 1
	if direction == "D" || direction == "L" {
		sign = -1
	}

	for i := 1; i <= steps; i++ {
		head.step(direction, sign)
		tail.follow(*head)
	}
}

func main() {
	f, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var head Head
	head.NewHead()

	var tail Tail
	tail.NewTail()
	visitedPos = make(map[string]int)
	visitedPos["0,0"] = 1

	for scanner.Scan() {
		moveSet := strings.Split(scanner.Text(), " ")
		direction := moveSet[0]
		steps, err := strconv.Atoi(moveSet[1])
		handleError(err)

		move(direction, steps, &head, &tail)
	}
	fmt.Println("Part 1:", len(visitedPos))
}
