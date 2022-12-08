package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("sample.in")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	var sum int

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		// fmt.Println(record)
		elf1 := strings.Split(record[0], "-")
		var elf1Int []int
		for _, val := range elf1 {
			valInt, _ := strconv.Atoi(val)
			elf1Int = append(elf1Int, valInt)
		}
		elf2 := strings.Split(record[1], "-")
		var elf2Int []int
		for _, val := range elf2 {
			valInt, _ := strconv.Atoi(val)
			elf2Int = append(elf2Int, valInt)
		}

		if elf1Int[0] <= elf2Int[0] && elf1Int[1] >= elf2Int[0] {
			sum++
		} else if elf2Int[0] <= elf1Int[0] && elf2Int[1] >= elf1Int[0] {
			sum++
		}
	}
	fmt.Println(sum)
}
