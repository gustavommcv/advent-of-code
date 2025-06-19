package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var forbiddenStrings [4]string = [...]string{"ab", "cd", "pq", "xy"}
var vowels [5]string = [...]string{"a", "e", "i", "o", "u"}

func main() {
	absolutePath, err := filepath.Abs("input.txt")
	checkError(err)
	inputData := scanFile(absolutePath)

	fmt.Println(getNiceStrings(string(inputData)))
}

func getNiceStrings(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		if isNice(scanner.Text()) {
			count++
		}
	}

	return count
}

func isNice(str string) bool {
	// Should not contain ab, cd, pq, or xy
	for _, forbiddenStr := range forbiddenStrings {
		if strings.Contains(str, forbiddenStr) {
			return false
		}
	}

	// Should contains at least three vowels
	vowelCount := 0
	for _, vowel := range vowels {
		vowelCount += strings.Count(str, vowel)
	}

	if vowelCount < 3 {
		return false
	}

	// Should contains at least one letter that appears twice in a row
	var prevChar rune
	isNice := false
	for _, currentChar := range str {
		if prevChar == currentChar {
			isNice = true
		}
		prevChar = currentChar
	}

	if !isNice {
		return false
	}

	return true
}

func scanFile(filePath string) []byte {
	fileContent, err := os.ReadFile(filePath)
	checkError(err)

	return fileContent
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}
