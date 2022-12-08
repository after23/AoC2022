package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	MAX   = 100000
	TOTAL = 70000000
	MIN   = 30000000
)

type Directory struct {
	name   string
	size   int
	depth  int
	parent *Directory
	subdir map[string]*Directory
}

func (dir *Directory) NewDir() {
	dir.size = 0
	dir.subdir = make(map[string]*Directory)
}

func (dir *Directory) addSubDir(name string, subdir *Directory) {
	dir.subdir[name] = subdir
}

func (dir *Directory) addToParent(n int) {
	var curDir *Directory
	for curDir = dir; curDir.parent != nil; curDir = curDir.parent {
		curDir.parent.size += n
	}
}

func (dir Directory) sum() int {
	var sum int
	if len(dir.subdir) == 0 {
		if dir.size <= MAX {
			return dir.size
		} else {
			return 0
		}
	}
	if dir.size <= MAX {
		sum += dir.size
	}
	for _, subdir := range dir.subdir {
		sum += subdir.sum()
	}
	return sum
}

func (dir Directory) findOptimal(opti, free int) int {
	newOpti := Optimal(free, dir.size)
	if len(dir.subdir) == 0 {
		if newOpti < opti {
			return newOpti
		}
		return opti
	}

	if opti < newOpti {
		newOpti = opti
	}

	for _, subdir := range dir.subdir {
		if subdir.findOptimal(newOpti, free) < newOpti {
			newOpti = subdir.findOptimal(newOpti, free)
		}
	}
	return newOpti
}

func Optimal(free, dirsize int) int {
	optimal := (free + dirsize) - MIN
	if optimal >= 0 {
		return dirsize
	} else {
		return TOTAL
	}
}

func main() {
	start := time.Now()
	f, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	var root Directory
	root.NewDir()
	root.name = "/"
	rootPtr := &root
	curDir := &root

	for scanner.Scan() {
		input := scanner.Text()
		if input == "$ ls" {
			continue
		} else if strings.Contains(input, "dir") {
			var subdir Directory
			subdir.NewDir()
			inputSlice := strings.Split(input, " ")
			subdir.name = inputSlice[1]
			curDir.addSubDir(inputSlice[1], &subdir)
		} else if strings.Contains(input, "cd ") {
			inputSlice := strings.Split(input, " ")
			if inputSlice[2] == ".." {
				curDir = curDir.parent
			} else {
				tempParent := curDir
				curDir = curDir.subdir[inputSlice[2]]
				curDir.parent = tempParent
				curDir.depth = curDir.parent.depth + 1
				tempParent = nil
			}
		} else {
			inputSlice := strings.Split(input, " ")
			size, _ := strconv.Atoi(inputSlice[0])
			curDir.size += size
			curDir.addToParent(size)
		}
	}
	free := TOTAL - rootPtr.size
	baseOpti := TOTAL
	fmt.Println(rootPtr.sum())
	fmt.Println(rootPtr.findOptimal(baseOpti, free))
	fmt.Println(time.Since(start))
}
