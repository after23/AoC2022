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

var start qItem
var destination qItem

var xLen int
var yLen int

func sliceInput(byteInput []byte, intInput *[]int) {
	for i, n := range byteInput {
		if int(n) == 83 {
			start.col = i
			start.row = yLen
			start.value = int(n)
			(*intInput)[i] = int(n)
			continue
		}
		if int(n) == 69 {
			destination.col = i
			destination.row = yLen
			destination.value = 27
			(*intInput)[i] = 27
			continue
		}
		(*intInput)[i] = int(n) - 96
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

func main() {
	fileName := "input.in"
	f, err := os.Open(fileName)
	errHandler(err)
	defer f.Close()

	cmd := exec.Command("wc", "-l", "input.in")
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

	queue := []qItem{}
	enqueue(&queue, start)
	for {
		curItem := dequeue(&queue)
		visitedItem[curItem.row][curItem.col] = true

		if curItem.value == destination.value {
			fmt.Println(curItem.dist)
			break
		}

		//go up
		if curItem.row-1 >= 0 && visitedItem[curItem.row-1][curItem.col] == false && intInput[curItem.row-1][curItem.col]-curItem.value <= 1 {
			var tempItem qItem
			tempItem.row = curItem.row - 1
			tempItem.col = curItem.col
			tempItem.value = intInput[curItem.row-1][curItem.col]
			tempItem.dist = curItem.dist + 1
			enqueue(&queue, tempItem)
			visitedItem[curItem.row-1][curItem.col] = true
		}

		//go down
		if curItem.row+1 < yLen && visitedItem[curItem.row+1][curItem.col] == false && intInput[curItem.row+1][curItem.col]-curItem.value <= 1 {
			var tempItem qItem
			tempItem.row = curItem.row + 1
			tempItem.col = curItem.col
			tempItem.value = intInput[curItem.row+1][curItem.col]
			tempItem.dist = curItem.dist + 1
			enqueue(&queue, tempItem)
			visitedItem[curItem.row+1][curItem.col] = true
		}

		//go left
		if curItem.col-1 >= 0 && visitedItem[curItem.row][curItem.col-1] == false && intInput[curItem.row][curItem.col-1]-curItem.value <= 1 {
			var tempItem qItem
			tempItem.row = curItem.row
			tempItem.col = curItem.col - 1
			tempItem.value = intInput[curItem.row][curItem.col-1]
			tempItem.dist = curItem.dist + 1
			enqueue(&queue, tempItem)
			visitedItem[curItem.row][curItem.col-1] = true
		}

		//go right
		if curItem.col+1 < xLen && visitedItem[curItem.row][curItem.col+1] == false && intInput[curItem.row][curItem.col+1]-curItem.value <= 1 {
			var tempItem qItem
			tempItem.row = curItem.row
			tempItem.col = curItem.col + 1
			tempItem.value = intInput[curItem.row][curItem.col+1]
			tempItem.dist = curItem.dist + 1
			enqueue(&queue, tempItem)
			visitedItem[curItem.row][curItem.col+1] = true
		}
	}
}
