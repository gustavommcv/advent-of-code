package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// north (^), south (v), east (>), or west (<)
const (
	NORTH = "^"
	SOUTH = "v"
	EAST  = ">"
	WEST  = "<"
)

func main() {
	absolutePath, err := filepath.Abs("input.txt")
	checkError(err)
	inputData := scanFile(absolutePath)

	for _, char := range inputData {
		fmt.Printf("%t\n", isDirectionChar(char))
	}

}

func scanFile(filePath string) []byte {
	fileContent, err := os.ReadFile(filePath)
	checkError(err)

	return fileContent
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func isDirectionChar(char byte) bool {
	if string(char) == NORTH || string(char) == SOUTH || string(char) == EAST || string(char) == WEST {
		return true
	}
	return false
}
