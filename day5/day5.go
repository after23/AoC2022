package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Stack struct {
	nodes []string
	count int
}

func NewStack() Stack {
	return Stack{}
}

func (stack *Stack) push(n string) {
	stack.nodes = append(stack.nodes, n)
	stack.count++
}

func (stack *Stack) pushFront(n string) {
	stack.nodes = append([]string{n}, stack.nodes...)
	stack.count++
}

func (stack *Stack) pop() string {
	if stack.count == 0 {
		return ""
	}
	res := stack.nodes[0]
	stack.nodes = stack.nodes[1:]
	stack.count--
	return res
}

func (stack *Stack) bulkPop(n int) []string {
	var res []string
	for n > 0 {
		res = append(res, stack.pop())
		n--
	}
	return res
}

func (stack *Stack) bulkPush(n []string) {
	stack.nodes = append(n, stack.nodes...)
	stack.count += len(n)
}

func main() {
	start := time.Now()
	f, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	var stackList []Stack
	for i := 0; i <= 8; i++ {
		stackList = append(stackList, NewStack())
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		re1 := regexp.MustCompile(`^[^m]`)
		if re1.MatchString(line) {
			re := regexp.MustCompile(`[A-Z]`)
			indexes := re.FindAllStringIndex(line, -1)
			for _, val := range indexes {
				i := val[0] / 4
				stackList[i].push(string(line[val[0]]))
			}
		} else {
			re := regexp.MustCompile(`[^\d\s]`)
			if line == "" {
				continue
			}
			line = re.ReplaceAllString(line, "")
			line = strings.TrimSpace(line)
			movement := strings.Split(line, " ")
			n, _ := strconv.Atoi(movement[0])
			src, _ := strconv.Atoi(movement[2])
			dst, _ := strconv.Atoi(movement[4])
			// for i := 1; i <= n; i++ {
			// 	stackList[dst-1].pushFront(stackList[src-1].pop())
			// }
			if n == 1 {
				stackList[dst-1].pushFront(stackList[src-1].pop())
			} else {
				stackList[dst-1].bulkPush(stackList[src-1].bulkPop(n))
			}
		}

	}
	var res string
	for _, val := range stackList {
		res += val.pop()
	}
	fmt.Println(res)
	fmt.Println(time.Since(start))
}
