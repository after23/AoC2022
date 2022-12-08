package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input, err := ioutil.ReadFile("input.in")
	if err != nil {
		log.Fatal(err.Error())
	}
	data := strings.Split(string(input), "\n")
	for i, str := range data {
		str = strings.TrimSpace(str)
		data[i] = str
	}
	count := 0
	score := 0
	for i, str := range data {
		for j, rune := range str {
			if i == 0 || i == len(data)-1 || j == 0 || j == len(data[i])-1 {
				count++
				continue
			}
			tree, _ := strconv.Atoi(string(rune))
			visibleLeft, scoreLeft := checkLeft(tree, str[:j])
			if visibleLeft {
				count++
			}
			visibleRight, scoreRight := checkRight(tree, str[j+1:])
			if !visibleLeft && visibleRight {
				count++
			}
			visibleUp, scoreUp := checkUp(tree, i, j, data)
			if !visibleLeft && !visibleRight && visibleUp {
				count++
			}
			visibleDown, scoreDown := checkDown(tree, i, j, data)
			if !visibleLeft && !visibleRight && !visibleUp && visibleDown {
				count++
			}
			tempScore := scoreLeft * scoreRight * scoreUp * scoreDown
			if tempScore > score {
				score = tempScore
			}
		}
	}
	fmt.Println("part 1 :", count)
	fmt.Println("part 2 :", score)
	fmt.Println(time.Since(start))
}

func checkLeft(tree int, str string) (bool, int) {
	score := 0
	for i := len(str) - 1; i >= 0; i-- {
		score++
		n, _ := strconv.Atoi(string(str[i]))
		if n >= tree {
			return false, score
		}
	}
	return true, score
}

func checkRight(tree int, str string) (bool, int) {
	score := 0
	for _, val := range str {
		score++
		n, _ := strconv.Atoi(string(val))
		if n >= tree {
			return false, score
		}
	}
	return true, score
}

func checkUp(tree, i, j int, data []string) (bool, int) {
	score := 0
	for y := i - 1; y >= 0; y-- {
		score++
		n, _ := strconv.Atoi(string(data[y][j]))
		if n >= tree {
			return false, score
		}
	}
	return true, score
}

func checkDown(tree, i, j int, data []string) (bool, int) {
	score := 0
	for y := i + 1; y < len(data); y++ {
		n, _ := strconv.Atoi(string(data[y][j]))
		score++
		if n >= tree {
			return false, score
		}
	}
	return true, score
}
