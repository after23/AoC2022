package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var rucksack1, rucksack2, rucksack3 string
	var priority, count int = 0, 1

	for scanner.Scan() {
		temp := scanner.Text()
		temp = strings.TrimSpace(temp)
		if count == 1 {
			rucksack1 = temp
			count++
		} else if count == 2 {
			rucksack2 = temp
			count++
		} else if count == 3 {
			rucksack3 = temp
			for _, val := range rucksack1 {
				if strings.Contains(rucksack2, string(val)) && strings.Contains(rucksack3, string(val)) {
					priority += prio(int(val))
					break
				}
			}
			count = 1
		}
	}

	fmt.Println(priority)

}

func prio(char int) int {
	var res int
	if char < 97 {
		res = char - 38
	} else {
		res = char - 96
	}
	return res
}
