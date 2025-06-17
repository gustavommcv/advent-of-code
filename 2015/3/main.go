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

type position struct {
	x int
	y int
}

func (p *position) String() string {
	return fmt.Sprintf("X: %d, Y: %d", p.x, p.y)
}

func main() {
	absolutePath, err := filepath.Abs("input.txt")
	checkError(err)
	inputData := scanFile(absolutePath)

	santasMap := startDeliveringWithRobot(inputData)
	printMap(santasMap)
	fmt.Println(len(santasMap))
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

func startDelivering(input []byte) map[position]int {
	santasMap := make(map[position]int)
	initialPosition := position{0, 0}

	currentPosition := initialPosition
	santasMap[initialPosition]++

	for _, element := range input {
		switch string(element) {
		case NORTH:
			currentPosition.y++
		case SOUTH:
			currentPosition.y--
		case WEST:
			currentPosition.x--
		case EAST:
			currentPosition.x++
		}

		santasMap[currentPosition]++
	}

	return santasMap
}

func startDeliveringWithRobot(input []byte) map[position]int {
	santasMap := make(map[position]int)
	initialPosition := position{0, 0}

	santasCurrentPosition := initialPosition
	robotCurrentPosition := initialPosition

	santasMap[initialPosition]++

	for i, element := range input {
		var currentPosition *position
		if isEven(i) {
			currentPosition = &santasCurrentPosition
		} else {
			currentPosition = &robotCurrentPosition
		}
		switch string(element) {
		case NORTH:
			currentPosition.y++
		case SOUTH:
			currentPosition.y--
		case WEST:
			currentPosition.x--
		case EAST:
			currentPosition.x++
		}

		santasMap[*currentPosition]++
	}

	return santasMap
}

func isEven(number int) bool {
	return number%2 == 0
}

func printMap(santasMap map[position]int) {
	for key, value := range santasMap {
		fmt.Printf("{%s} = %d\n", key.String(), value)
	}
}

func isDirectionChar(char byte) bool {
	if string(char) == NORTH || string(char) == SOUTH || string(char) == EAST || string(char) == WEST {
		return true
	}
	return false
}
