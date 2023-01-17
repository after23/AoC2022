package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func errHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type qItem struct {
	col   int
	row   int
	value int
	dist  int
}

const UintSize = 32 << (^uint(0) >> 32 & 1)

const maxInt = 1<<(UintSize-1) - 1

var start qItem
var destination qItem

var xLen int
var yLen int

var startingPoints []qItem

func sliceInput(byteInput []byte, intInput *[]int) {
	for i, n := range byteInput {
		if int(n) == 83 {
			start.col = i
			start.row = yLen
			start.value = 1
			(*intInput)[i] = 1
			enqueue(&startingPoints, start)
			continue
		}
		if int(n) == 69 {
			destination.col = i
			destination.row = yLen
			destination.value = 27
			(*intInput)[i] = 27
			continue
		}
		tempInt := int(n) - 96
		if tempInt == 1 {
			var tempItem qItem
			tempItem.col = i
			tempItem.row = yLen
			tempItem.value = tempInt
			enqueue(&startingPoints, tempItem)
		}
		(*intInput)[i] = tempInt
	}
}

func enqueue(queue *[]qItem, item qItem) {
	*queue = append(*queue, item)
}

func dequeue(queue *[]qItem) qItem {
	var res qItem = (*queue)[0]
	if len(*queue) == 1 {
		*queue = []qItem{}
	} else {
		*queue = (*queue)[1:]
	}
	return res
}

func resetVisited(visitedItem [][]bool) {
	for i := 0; i < yLen; i++ {
		for j := 0; j < xLen; j++ {
			visitedItem[i][j] = false
		}
	}
}

func minDist(queue *[]qItem, visitedItem [][]bool, intInput *[][]int) int {
	for {
		curItem := dequeue(queue)
		visitedItem[curItem.row][curItem.col] = true

		if curItem.value == destination.value {
			// fmt.Println(curItem.dist)
			resetVisited(visitedItem)
			return curItem.dist
		}

		//go up
		if curItem.row-1 >= 0 && visitedItem[curItem.row-1][curItem.col] == false && (*intInput)[curItem.row-1][curItem.col]-curItem.value <= 1 {
			var tempItem qItem
			tempItem.row = curItem.row - 1
			tempItem.col = curItem.col
			tempItem.value = (*intInput)[curItem.row-1][curItem.col]
			tempItem.dist = curItem.dist + 1
			enqueue(queue, tempItem)
			visitedItem[curItem.row-1][curItem.col] = true
		}

		//go down
		if curItem.row+1 < yLen && visitedItem[curItem.row+1][curItem.col] == false && (*intInput)[curItem.row+1][curItem.col]-curItem.value <= 1 {
			var tempItem qItem
			tempItem.row = curItem.row + 1
			tempItem.col = curItem.col
			tempItem.value = (*intInput)[curItem.row+1][curItem.col]
			tempItem.dist = curItem.dist + 1
			enqueue(queue, tempItem)
			visitedItem[curItem.row+1][curItem.col] = true
		}

		//go left
		if curItem.col-1 >= 0 && visitedItem[curItem.row][curItem.col-1] == false && (*intInput)[curItem.row][curItem.col-1]-curItem.value <= 1 {
			var tempItem qItem
			tempItem.row = curItem.row
			tempItem.col = curItem.col - 1
			tempItem.value = (*intInput)[curItem.row][curItem.col-1]
			tempItem.dist = curItem.dist + 1
			enqueue(queue, tempItem)
			visitedItem[curItem.row][curItem.col-1] = true
		}

		//go right
		if curItem.col+1 < xLen && visitedItem[curItem.row][curItem.col+1] == false && (*intInput)[curItem.row][curItem.col+1]-curItem.value <= 1 {
			var tempItem qItem
			tempItem.row = curItem.row
			tempItem.col = curItem.col + 1
			tempItem.value = (*intInput)[curItem.row][curItem.col+1]
			tempItem.dist = curItem.dist + 1
			enqueue(queue, tempItem)
			visitedItem[curItem.row][curItem.col+1] = true
		}

		if len(*queue) == 0 {
			return maxInt
		}
	}
}

func main() {
	fileName := "input.in"
	f, err := os.Open(fileName)
	errHandler(err)
	defer f.Close()

	cmd := exec.Command("wc", "-l", fileName)
	cmdOut, err := cmd.CombinedOutput()
	errHandler(err)
	fileLenStr := strings.Fields(string(cmdOut))[0]
	fileLen, err := strconv.Atoi(fileLenStr)
	errHandler(err)
	intInput := make([][]int, fileLen+1)
	visitedItem := make([][]bool, fileLen+1)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input := scanner.Text()
		if xLen == 0 {
			xLen = len(input)
		}
		intInput[yLen] = make([]int, xLen)
		visitedItem[yLen] = make([]bool, xLen)
		byteInput := []byte(input)
		sliceInput(byteInput, &intInput[yLen])

		yLen++
	}

	minInt := maxInt

	for len(startingPoints) > 0 {
		curItem := dequeue(&startingPoints)
		queue := []qItem{}
		enqueue(&queue, curItem)
		tempMin := minDist(&queue, visitedItem, &intInput)
		if tempMin < minInt {
			minInt = tempMin
		}
	}
	fmt.Println(minInt)
}
