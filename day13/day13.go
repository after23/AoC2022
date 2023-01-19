package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const fileName = "sample.in"

type signal struct {
	packets []int
}

type signalList []signal

var firstSignal signalList
var secondSignal signalList
var signalCounter = 1

func errHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (signalL *signalList) appenItem(input []string, rawInput, target string) string {
	var intInput []int
	for _, val := range input {
		intVal, err := strconv.Atoi(val)
		errHandler(err)
		intInput = append(intInput, intVal)
	}
	var tempSignal signal
	tempSignal.packets = intInput
	*signalL = append(*signalL, tempSignal)
	rawInput = strings.TrimPrefix(rawInput, target)
	return rawInput
}

func (signalL *signalList) insert(rawInput string) {
	for len(rawInput) > 0 {
		re := regexp.MustCompile(`(?U)^,*\[[\d,]*\]|(^[\[,].,|^[\[,]\d+,)(\d|\[)|^.\d\d|^,\d|^\d+[,\]\[]|[\[\],]`)
		reInt := regexp.MustCompile(`\d+`)
		res := re.FindStringSubmatch(rawInput)
		if res[1] != "" {
			input := reInt.FindString(res[1])
			rawInput = signalL.appenItem([]string{input}, rawInput, res[1])
			continue
		}
		if res[0] == "]" || res[0] == "," {
			rawInput = strings.TrimPrefix(rawInput, res[0])
			continue
		}
		if res[0][0] == '[' {
			reTemp := regexp.MustCompile(`\[`)
			loc := reTemp.FindStringIndex(rawInput)
			target := rawInput[loc[0]+1]
			if target == '[' {
				rawInput = signalL.appenItem([]string{"-1"}, rawInput, res[0])
				continue
			}
		}
		if res[0][0] == '[' || res[0][1] == '[' {
			input := reInt.FindAllString(res[0], -1)
			rawInput = signalL.appenItem(input, rawInput, res[0])
			continue
		}
		input := reInt.FindString(rawInput)
		rawInput = signalL.appenItem([]string{input}, rawInput, res[0])
	}
}

func main() {
	f, err := os.Open(fileName)
	errHandler(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// counter := 1

	for scanner.Scan() {
		rawInput := scanner.Text()
		if rawInput == "" {
			//implement the signal comparison here
			//reset the signal counter
			signalCounter = 1
			//reset the signal list var
			firstSignal = []signal{}
			secondSignal = []signal{}
			continue
		}
		if signalCounter == 1 {
			//process first signal input
			fmt.Println("1 : ", rawInput)
			firstSignal.insert(rawInput)
			signalCounter++
			fmt.Println(firstSignal)
		} else {
			//process second signal input
			fmt.Println("2 : ", rawInput)
			secondSignal.insert(rawInput)
			signalCounter++
			fmt.Println(secondSignal)
		}
	}
}
