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

type NineHead [10]Point

var visitedPos map[string]int
var visitedPos2 map[string]int

func (point *Point) NewPoint() {
	point.x, point.y = 0, 0
}

func (heads *NineHead) NewHeads() {
	for i := 0; i <= 8; i++ {
		heads[i].x, heads[i].y = 0, 0
	}
}

func isNear(tail Point, head Point) bool {
	deltaX := head.x - tail.x
	deltaY := head.y - tail.y

	if moreThanOne(deltaX) || moreThanOne(deltaY) {
		return false
	}
	return true
}

func (head *Point) step(direction string, sign int) {
	if direction == "D" || direction == "U" {
		head.y += sign
	} else {
		head.x += sign
	}
}

func (head *NineHead) steps(direction string, sign int) {
	for i := 1; i < len(head); i++ {
		if i == len(head)-1 {
			head[i].follow(head[i-1], 2)
			return
		}
		head[i].follow(head[i-1], 3)
	}
}

func (tail *Point) follow(head Point, choice int) {
	if isNear(Point(*tail), Point(head)) {
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

	if diagonalMove(deltaX, deltaY) {
		tail.x += signX
		tail.y += signY
		updateVisitedPos(tail.x, tail.y, choice)
		return
	}

	if moreThanOne(deltaX) {
		tail.y = head.y
		tail.x += signX
		updateVisitedPos(tail.x, tail.y, choice)
		return
	}

	if moreThanOne(deltaY) {
		tail.y += signY
		tail.x = head.x
		updateVisitedPos(tail.x, tail.y, choice)
		return
	}
}

func diagonalMove(x, y int) bool {
	if moreThanOne(x) && moreThanOne(y) {
		return true
	}
	return false
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

func updateVisitedPos(x, y, choice int) {
	if choice == 1 {
		position := fmt.Sprintf("%d,%d", x, y)
		visitedPos[position] = 1
		return
	} else if choice == 2 {
		position := fmt.Sprintf("%d,%d", x, y)
		visitedPos2[position] = 1
		return
	}
}

func sign(direction string) int {
	sign := 1
	if direction == "D" || direction == "L" {
		sign = -1
	}
	return sign
}

func move(direction string, steps int, head *Point, tail *Point) {
	sign := sign(direction)

	for i := 1; i <= steps; i++ {
		head.step(direction, sign)
		tail.follow(*head, 1)
	}
}

func partTwoMove(direction string, steps int, head *NineHead) {
	sign := sign(direction)

	for i := 0; i < steps; i++ {
		head[0].step(direction, sign)
		head.steps(direction, sign)
	}
}

func main() {
	f, err := os.Open("../input.in")
	handleError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var head Point
	head.NewPoint()

	var tail Point
	tail.NewPoint()

	var nineHead NineHead
	nineHead.NewHeads()

	visitedPos = make(map[string]int)
	visitedPos2 = make(map[string]int)

	visitedPos["0,0"] = 1
	visitedPos2["0,0"] = 1

	for scanner.Scan() {
		moveSet := strings.Split(scanner.Text(), " ")
		direction := moveSet[0]
		steps, err := strconv.Atoi(moveSet[1])
		handleError(err)

		move(direction, steps, &head, &tail)
		partTwoMove(direction, steps, &nineHead)
	}
	fmt.Println("Part 1:", len(visitedPos))
	fmt.Println("Part 2:", len(visitedPos2))

}
