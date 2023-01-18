package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const fileName = "sample.in"

type signal struct {
	packets []int
}

type signalList []signal

var firstSignal signalList
var secondSignal signalList

func errHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	f, err := os.Open(fileName)
	errHandler(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	signalCounter := 1
	// counter := 1

	for scanner.Scan() {
		rawInput := scanner.Text()
		if rawInput == "" {
			signalCounter = 1
			continue // or do something
		}
		if signalCounter == 1 {
			//process first signal input
			for len(rawInput) > 0 {
				subString := strings.Trim(rawInput, "[]")
				re := regexp.MustCompile(`\[.*\]`)
				subString = re.ReplaceAllString(subString, "")
				// subString = strings.Split(subString, ",")
				fmt.Println(subString)
				break
			}
		} else {
			//process second signal input
		}
		fmt.Println("hey you")
	}
}
