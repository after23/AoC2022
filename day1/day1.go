package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type TopThree struct {
	top3 [3]int
}

func (t3 *TopThree) add(input int) {
	for i, val := range t3.top3 {
		if input > val {
			t3.top3[i] = input
			sort.Ints(t3.top3[:])
			break
		}
	}
}

func (t3 TopThree) sum() int {
	var res int
	for _, val := range t3.top3 {
		res += val
	}
	return res
}

func main() {
	f, err := os.Open("input-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	curSum := 0
	max := 0
	var t3 TopThree
	for scanner.Scan() {
		if scanner.Text() == "" {
			count++
			if curSum > max {
				max = curSum
			}
			t3.add(curSum)
			curSum = 0
			continue
		}
		curVal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		curSum += curVal
	}

	fmt.Println(max)
	fmt.Println(t3.sum())
	fmt.Println(t3.top3)

}
