package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var x int = 1

type Sprite struct {
	position [3]int
}

func (sprite *Sprite) newSprite() {
	sprite.position = [3]int{0, 1, 2}
}

func (sprite *Sprite) updatePosition(n int) {
	sprite.position = [3]int{n - 1, n, n + 1}
}

func (sprite *Sprite) isIntersect(pos int) bool {
	for _, val := range sprite.position {
		if val == pos {
			return true
		}
	}
	return false
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func signalStrength(cycle, x int) (int, bool) {
	if (cycle-20)%40 != 0 {
		return 0, false
	}
	return x * cycle, true
}

func addSum(sum, val int, ok bool) int {
	if !ok {
		return sum
	}

	sum += val
	return sum
}

func printLine(line string) string {
	if len(line) == 40 {
		fmt.Println(line)
		return ""
	}
	return line
}

func addLine(sprite Sprite, line string) string {
	length := len(line)
	if sprite.isIntersect(length) {
		line += "#"
	} else {
		line += "."
	}
	return line
}

func main() {
	f, err := os.Open("input.in")
	handleErr(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	cycle := 1
	sum := 0
	var line string
	var sprite Sprite
	sprite.newSprite()

	fmt.Println("Part 2:")
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		command := input[0]
		var v int
		if len(input) > 1 {
			v, err = strconv.Atoi(input[1])
			handleErr(err)
		}
		switch command {
		case "noop":
			line = addLine(sprite, line)
			cycle++
			line = printLine(line)
			val, ok := signalStrength(cycle, x)
			sum = addSum(sum, val, ok)
		case "addx":
			line = addLine(sprite, line)
			cycle++
			line = printLine(line)
			val, ok := signalStrength(cycle, x)
			sum = addSum(sum, val, ok)
			line = addLine(sprite, line)
			cycle++
			line = printLine(line)
			x += v
			sprite.updatePosition(x)
			val, ok = signalStrength(cycle, x)
			sum = addSum(sum, val, ok)
		}
	}
	fmt.Println("Part 1:", sum)
}
