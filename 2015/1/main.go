package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	relativePath, err := filepath.Abs("input.txt")
	checkError(err)

	input := scanFile(relativePath)
	integer := WhatFloor(input)

	fmt.Println(integer)
}

func scanFile(filePath string) []byte {
	input, err := os.ReadFile(filePath)
	checkError(err)

	return input
}

func WhatFloor(input []byte) int {
	const LEFT_PARENTHESIS string = "("
	const RIGHT_PARENTHESIS string = ")"

	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	scanner.Split(bufio.ScanRunes)

	floor := 0
	for scanner.Scan() {
		if scanner.Text() == LEFT_PARENTHESIS {
			floor++
			continue
		}

		if scanner.Text() == RIGHT_PARENTHESIS {
			floor--
		}
	}

	return floor
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
