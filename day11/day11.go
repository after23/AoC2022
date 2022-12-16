package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monke struct {
	items        []int
	operation    func(int) int
	test         func(int) bool
	testTrue     int
	testFalse    int
	inspectCount int
}

type Monkeys []*Monke

func (monke *Monke) pop() {
	monke.items = monke.items[1:]
}

func (monke *Monke) push(item int) {
	monke.items = append(monke.items, item)
}

func (monkeys Monkeys) evaluateTest(i, item, testTrue, testFalse int) {
	if monkeys[i].test(item) {
		monkeys[testTrue].push(item)
	} else {
		monkeys[testFalse].push(item)
	}
	monkeys[i].inspectCount++
}

func (monkeys Monkeys) examine() {
	for i, val := range monkeys {
		for _, item := range val.items {
			item = val.operation(item)
			item /= 3
			monkeys.evaluateTest(i, item, val.testTrue, val.testFalse)
			val.pop()
		}
	}
}

func monkeBusiness(monkeys Monkeys) int {
	var inspectCounts []int
	for _, monke := range monkeys {
		inspectCounts = append(inspectCounts, monke.inspectCount)
	}
	sort.Ints(inspectCounts)
	return inspectCounts[len(inspectCounts)-1] * inspectCounts[len(inspectCounts)-2]
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (monkeys Monkeys) addItems(idx int, items []string) {
	for _, item := range items {
		itemInt, err := strconv.Atoi(strings.TrimSpace(item))
		handleErr(err)
		monkeys[idx].push(itemInt)
	}
}

func (monke *Monke) plusOperation(param string) {
	switch param {
	case "old":
		monke.operation = func(item int) int {
			return item + item
		}
	default:
		paramInt, err := strconv.Atoi(param)
		handleErr(err)
		monke.operation = func(item int) int {
			return item + paramInt
		}
	}
}

func (monke *Monke) sumOperation(param string) {
	switch param {
	case "old":
		monke.operation = func(item int) int {
			return item * item
		}
	default:
		paramInt, err := strconv.Atoi(param)
		handleErr(err)
		monke.operation = func(item int) int {
			return item * paramInt
		}
	}
}

func (monke *Monke) setOperation(param, operator string) {
	switch operator {
	case "+":
		monke.plusOperation(param)
	default:
		monke.sumOperation(param)
	}
}

func (monkeys Monkeys) addOperation(idx int, input string) {
	var operator string
	if strings.Contains(input, "*") {
		operator = "*"
	} else {
		operator = "+"
	}
	lines := strings.Split(input, operator)
	param := strings.TrimSpace(lines[1])
	monkeys[idx].setOperation(param, operator)
}

func (monke *Monke) setTest(param int) {
	monke.test = func(item int) bool {
		return item%param == 0
	}
}

func main() {
	f, err := os.Open("input.in")
	handleErr(err)
	defer f.Close()

	var monkeys Monkeys
	scanner := bufio.NewScanner(f)
	idx := -1

	for scanner.Scan() {
		input := scanner.Text()
		if strings.HasPrefix(input, "Monkey") {
			var monke Monke
			monkeys = append(monkeys, &monke)
			idx++
			continue
		}
		if strings.Contains(input, "Starting") {
			lines := strings.Split(input, ":")
			items := strings.Split(lines[1], ",")
			monkeys.addItems(idx, items)
			continue
		}
		if strings.Contains(input, "Operation") {
			monkeys.addOperation(idx, input)
			continue
		}
		if strings.Contains(input, "Test") {
			lines := strings.Split(input, "by")
			param := strings.TrimSpace(lines[1])
			paramInt, err := strconv.Atoi(param)
			handleErr(err)
			monkeys[idx].setTest(paramInt)
			continue
		}
		if strings.Contains(input, "true") {
			param := string(input[len(input)-1])
			paramInt, err := strconv.Atoi(param)
			handleErr(err)
			monkeys[idx].testTrue = paramInt
			continue
		}
		if strings.Contains(input, "false") {
			param := string(input[len(input)-1])
			paramInt, err := strconv.Atoi(param)
			handleErr(err)
			monkeys[idx].testFalse = paramInt
			continue
		}
	}
	for i := 0; i < 20; i++ {
		monkeys.examine()
	}
	fmt.Println("Part 1 :", monkeBusiness(monkeys))
}
