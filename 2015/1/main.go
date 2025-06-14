package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const LEFT_PARENTHESIS string = "("
const RIGHT_PARENTHESIS string = ")"

func main() {
	relativePath, err := filepath.Abs("input.txt")
	checkError(err)

	input := scanFile(relativePath)
	floor := GetFloor(input)
	position := GetPosition(input)

	fmt.Println("Floor: ", floor)
	fmt.Println("Position: ", position)
}

func scanFile(filePath string) []byte {
	input, err := os.ReadFile(filePath)
	checkError(err)

	return input
}

func GetPosition(input []byte) int {
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	scanner.Split(bufio.ScanRunes)

	position := 0
	floor := 0
	for scanner.Scan() {
		if scanner.Text() == LEFT_PARENTHESIS || scanner.Text() == RIGHT_PARENTHESIS {
			position++
		}

		if scanner.Text() == LEFT_PARENTHESIS {
			floor++
		}

		if scanner.Text() == RIGHT_PARENTHESIS {
			floor--
		}

		if floor == -1 {
			break
		}
	}

	return position
}

func GetFloor(input []byte) int {
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
