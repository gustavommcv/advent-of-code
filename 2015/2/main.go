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
	relativePath, err := filepath.Abs("input.txt")
	checkError(err)
	input := scanFile(relativePath)

	presents := ParsePresents(input)
	presentWrapers := GetPresentWrapers(presents)

	fmt.Println(GetTotalSquareFeet(presentWrapers))
}

type Present struct {
	lenght int
	width  int
	height int
}

type PresentWrap struct {
	lenght     int
	width      int
	heigth     int
	squareFeet int
}

func (p *PresentWrap) GetExtraPaper() int {
	s1 := p.lenght * p.width
	s2 := p.lenght * p.heigth
	s3 := p.width * p.heigth

	return min(s1, s2, s3)
}

func scanFile(filePath string) []byte {
	input, err := os.ReadFile(filePath)
	checkError(err)

	return input
}

func ParsePresents(input []byte) []Present {
	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	presents := []Present{}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "x")

		lenght, err := strconv.Atoi(input[0])
		checkError(err)

		width, err := strconv.Atoi(input[1])
		checkError(err)

		height, err := strconv.Atoi(input[2])
		checkError(err)

		presents = append(presents, Present{lenght: lenght, width: width, height: height})
	}

	return presents
}

func GetPresentWraper(present Present) PresentWrap {
	presentWrap := PresentWrap{lenght: present.lenght, width: present.width, heigth: present.height}
	presentWrap.squareFeet = 2*present.lenght*present.width + 2*present.width*present.height + 2*present.height*present.lenght

	return presentWrap
}

func GetPresentWrapers(presents []Present) []PresentWrap {
	presentWrapers := []PresentWrap{}
	for _, present := range presents {
		presentWrapers = append(presentWrapers, GetPresentWraper(present))
	}

	return presentWrapers
}

func GetTotalSquareFeet(presentWrapers []PresentWrap) int {
	total := 0

	for _, presentWrap := range presentWrapers {
		total += presentWrap.squareFeet + presentWrap.GetExtraPaper()
	}

	return total
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
