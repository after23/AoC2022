package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func unique(str string) (bool, int) {
	for i, char := range str {
		if strings.Contains(str[i+1:], string(char)) {
			return false, i
		}
	}
	return true, 5
}

func main() {
	f, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	text := string(bytes)
	i := 0
	for {
		unique, n := unique(text[i : i+14])
		if !unique {
			i += n + 1
		} else {
			fmt.Println(text[i : i+14])
			fmt.Println(i + 14)
			break
		}
	}
}
