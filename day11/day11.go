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
	items        []int64
	operation    func(int64) int64
	test         func(int64) bool
	testTrue     int
	testFalse    int
	inspectCount int64
	divisors     int64
}

type Monkeys []*Monke

func (monke *Monke) pop() {
	monke.items = monke.items[1:]
}

func (monke *Monke) push(item int64) {
	monke.items = append(monke.items, item)
}

func (monkeys Monkeys) evaluateTest(i, testTrue, testFalse int, item int64) {
	if monkeys[i].test(item) {
		monkeys[testTrue].push(item)
	} else {
		monkeys[testFalse].push(item)
	}
	monkeys[i].inspectCount++
}

func (monkeys Monkeys) examine(n int64) {
	for i, val := range monkeys {
		for _, item := range val.items {
			item = val.operation(item)
			item %= n
			monkeys.evaluateTest(i, val.testTrue, val.testFalse, item)
			val.pop()
		}
	}
}

func (monkeys Monkeys) elp() int64 {
	var res int64 = 1
	for _, val := range monkeys {
		res *= val.divisors
	}
	return res
}

func monkeBusiness(monkeys Monkeys) int64 {
	var inspectCounts []int64
	for _, monke := range monkeys {
		inspectCounts = append(inspectCounts, monke.inspectCount)
	}
	sort.Slice(inspectCounts, func(i, j int) bool { return inspectCounts[i] < inspectCounts[j] })
	fmt.Println(inspectCounts)
	// sort.Ints(inspectCounts)
	return inspectCounts[len(inspectCounts)-1] * inspectCounts[len(inspectCounts)-2]
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (monkeys Monkeys) addItems(idx int, items []string) {
	for _, item := range items {
		// itemInt, err := strconv.Atoi(strings.TrimSpace(item))
		itemInt, err := strconv.ParseInt(strings.TrimSpace(item), 10, 64)
		handleErr(err)
		monkeys[idx].push(itemInt)
	}
}

func (monke *Monke) plusOperation(param string) {
	switch param {
	case "old":
		monke.operation = func(item int64) int64 {
			return item + item
		}
	default:
		// paramInt, err := strconv.Atoi(param)
		paramInt, err := strconv.ParseInt(strings.TrimSpace(param), 10, 64)
		handleErr(err)
		monke.operation = func(item int64) int64 {
			return item + paramInt
		}
	}
}

func (monke *Monke) sumOperation(param string) {
	switch param {
	case "old":
		monke.operation = func(item int64) int64 {
			return item * item
		}
	default:
		// paramInt, err := strconv.Atoi(param)
		paramInt, err := strconv.ParseInt(strings.TrimSpace(param), 10, 64)
		handleErr(err)
		monke.operation = func(item int64) int64 {
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

func (monke *Monke) setTest(param int64) {
	monke.test = func(item int64) bool {
		return item%param == 0
	}
	monke.divisors = param
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
			// paramInt, err := strconv.Atoi(param)
			paramInt, err := strconv.ParseInt(strings.TrimSpace(param), 10, 64)
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
	n := monkeys.elp()
	for i := 0; i < 10000; i++ {
		monkeys.examine(n)
	}

	fmt.Println("Part 2 :", monkeBusiness(monkeys))
}
