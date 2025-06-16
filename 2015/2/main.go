package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	absolutePath, err := filepath.Abs("input.txt")
	checkError(err)
	inputData := scanFile(absolutePath)

	presents := parsePresents(inputData)
	presentWrappers := getPresentWrappers(presents)

	fmt.Println(getTotalSquareFeet(presentWrappers))
}

type Present struct {
	length int
	width  int
	height int
}

type PresentWrapper struct {
	length     int
	width      int
	height     int
	squareFeet int
}

func (pw *PresentWrapper) getExtraPaper() int {
	side1 := pw.length * pw.width
	side2 := pw.length * pw.height
	side3 := pw.width * pw.height

	return min(side1, side2, side3)
}

func scanFile(filePath string) []byte {
	fileContent, err := os.ReadFile(filePath)
	checkError(err)

	return fileContent
}

func parsePresents(inputData []byte) []Present {
	scanner := bufio.NewScanner(strings.NewReader(string(inputData)))

	var presents []Present
	for scanner.Scan() {
		dimensions := strings.Split(scanner.Text(), "x")

		length, err := strconv.Atoi(dimensions[0])
		checkError(err)

		width, err := strconv.Atoi(dimensions[1])
		checkError(err)

		height, err := strconv.Atoi(dimensions[2])
		checkError(err)

		presents = append(presents, Present{length: length, width: width, height: height})
	}

	return presents
}

func getPresentWrapper(present Present) PresentWrapper {
	presentWrapper := PresentWrapper{
		length: present.length,
		width:  present.width,
		height: present.height,
	}
	presentWrapper.squareFeet = 2*present.length*present.width +
		2*present.width*present.height +
		2*present.height*present.length

	return presentWrapper
}

func getPresentWrappers(presents []Present) []PresentWrapper {
	var presentWrappers []PresentWrapper
	for _, present := range presents {
		presentWrappers = append(presentWrappers, getPresentWrapper(present))
	}

	return presentWrappers
}

func getTotalSquareFeet(presentWrappers []PresentWrapper) int {
	totalSquareFeet := 0

	for _, wrapper := range presentWrappers {
		totalSquareFeet += wrapper.squareFeet + wrapper.getExtraPaper()
	}

	return totalSquareFeet
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
