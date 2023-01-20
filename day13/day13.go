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

// SkipToNextValidSignal is used to return the index of the next signal with valid packets // Skip signals that has -1 as packet's value
func (signals signalList) skipToNextValidSignal(index int) int {
	if len(signals[index].packets) == 0 {
		return -1
	}
	signal := signals[index]
	if !signal.isSkipSignal() {
		return index
	}
	index++
	for index < len(signals) {
		signal = signals[index]
		if !signal.isSkipSignal() {
			return index
		}
		index++
	}
	return index
}

// isSkipSignal is used to check if a signal's packet is valid or not
// non valid packet has value of -1
func (signal signal) isSkipSignal() bool {
	if len(signal.packets) == 0 {
		return false
	}
	if signal.packets[0] == -1 {
		return true
	}
	return false
}

// 1 = ordered
// 2 = false ordered
// 3 = lanjut
func isPacketOrdered(firstIndex, secondIndex, max int) int {
	firstPacket := firstSignal[firstIndex].packets
	secondPacket := secondSignal[secondIndex].packets
	if max == 0 {
		if len(firstPacket) == 0 && len(secondPacket) != 0 {
			return 1
		}
		if len(secondPacket) == 0 && len(firstPacket) != 0 {
			return 2
		}
	}
	for i := 0; i < max; i++ {
		firstItem := firstPacket[i]
		secondItem := secondPacket[i]
		fmt.Printf("%d vs %d\n", firstItem, secondItem)
		// false ordered
		if secondItem < firstItem {
			return 2
		}
		// ordered
		if firstItem < secondItem {
			return 1
		}
	}
	// first signal run out of item so it is ordered
	if max == len(firstPacket) && max != len(secondPacket) {
		return 1
	}

	// second signal run out of item so it is not ordered correctly
	if max == len(secondPacket) && max != len(firstPacket) {
		return 2
	}
	// no conclusion so check the next packet
	return 3
}

func iterateOverSignalList() bool {
	firstLen := len(firstSignal)
	secondLen := len(secondSignal)
	max := firstLen
	if secondLen < firstLen {
		max = secondLen
	}
	firstIndex := 0
	secondIndex := 0
	for i := 0; i < max; i++ {
		firstSignalMessage := firstSignal[firstIndex]
		secondSignalMessage := secondSignal[secondIndex]

		tempFirstIndex := firstSignal.skipToNextValidSignal(firstIndex)
		tempSecondIndex := secondSignal.skipToNextValidSignal(secondIndex)

		if secondIndex != tempSecondIndex && tempSecondIndex != -1 {
			secondIndex = tempSecondIndex
			secondSignalMessage = secondSignal[secondIndex]
		}

		if firstIndex != tempFirstIndex && tempFirstIndex != -1 {
			firstIndex = tempFirstIndex
			firstSignalMessage = firstSignal[firstIndex]
		}

		maxPacketLen := len(firstSignalMessage.packets)
		if len(secondSignalMessage.packets) < len(firstSignalMessage.packets) {
			maxPacketLen = len(secondSignalMessage.packets)
		}

		if secondIndex == maxPacketLen && len(secondSignalMessage.packets) == 0 {
			return false
		}

		if firstIndex == maxPacketLen && len(firstSignalMessage.packets) == 0 {
			return true
		}

		option := isPacketOrdered(firstIndex, secondIndex, maxPacketLen)
		if option == 1 {
			return true
		}
		if option == 2 {
			return false
		}
		firstIndex++
		secondIndex++
		if firstLen != secondLen && firstIndex == max {
			return true
		}
		if secondLen != firstLen && secondIndex == max {
			return false
		}
	}
	return false
}

func main() {
	f, err := os.Open(fileName)
	errHandler(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	pairCounter := 1
	sum := 0

	for scanner.Scan() {
		rawInput := scanner.Text()
		if rawInput == "" {
			//implement the signal comparison here
			//reset the signal counter
			signalCounter = 1
			fmt.Println(firstSignal)
			fmt.Println(secondSignal)
			if iterateOverSignalList() {
				sum += pairCounter
			}
			//increment the pair counter
			pairCounter++
			//reset the signal list var
			firstSignal = []signal{}
			secondSignal = []signal{}
			continue
		}
		if signalCounter == 1 {
			//process first signal input
			// fmt.Println("1 : ", rawInput)
			firstSignal.insert(rawInput)
			signalCounter++
			// fmt.Println(firstSignal)
		} else {
			//process second signal input
			// fmt.Println("2 : ", rawInput)
			secondSignal.insert(rawInput)
			signalCounter++
			// fmt.Println(secondSignal)
		}
	}
	fmt.Println(sum)

}
