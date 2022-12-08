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
	var priority int = 0

	for scanner.Scan() {
		temp := scanner.Text()
		temp = strings.TrimSpace(temp)
		half := len(temp) / 2
		rucksack1 := temp[:half]
		rucksack2 := temp[half:]

		for _, val := range rucksack1 {
			if strings.Contains(rucksack2, string(val)) {
				priority += prio(int(val))
				rucksack2 = strings.ReplaceAll(rucksack2, string(val), "@")
			}
		}
	}

	fmt.Println(priority)
	a := 'z'
	fmt.Println(prio(int(a)))

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
