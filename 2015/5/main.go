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

	// fmt.Println(getNiceStrings(string(inputData)))
	fmt.Println(getNiceStrings2(string(inputData)))
	// fmt.Println(isNice2("xyxy"))
	// fmt.Println(isNice2("aaabcdefgaa"))
	// fmt.Println(isNice2("uurcxstgmygtbstg"))
	// fmt.Println(isNice2("ieodomkazucvgmuy"))
}

func getNiceStrings2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		if isNice2(scanner.Text()) {
			count++
		}
	}

	return count
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
	if containsForbidden(str) {
		return false
	}

	// Should contains at least three vowels
	if vowelsCount(str) < 3 {
		return false
	}

	// Should contains at least one letter that appears twice in a row
	if !containsPair(str) {
		return false
	}

	return true
}

func containsForbidden(str string) bool {
	for _, forbiddenStr := range forbiddenStrings {
		if strings.Contains(str, forbiddenStr) {
			return true
		}
	}

	return false
}

func vowelsCount(str string) int {
	vowelCount := 0
	for _, vowel := range vowels {
		vowelCount += strings.Count(str, vowel)
	}

	return vowelCount
}

func containsPair(str string) bool {
	var prevChar rune
	containsPair := false
	for _, currentChar := range str {
		if prevChar == currentChar {
			containsPair = true
		}
		prevChar = currentChar
	}

	if !containsPair {
		return false
	}

	return true
}

func isNice2(str string) bool {
	return containsPairWithoutOverlapping(str) && containsRepeatingLetterWithOneBetween(str)
}

func containsPairWithoutOverlapping(str string) bool {
	pairs := make(map[string]int)

	for i := range len(str) - 1 {
		pair := str[i : i+2]
		if lastIndex, exists := pairs[pair]; exists {
			if lastIndex < i-1 {
				return true
			}
		} else {
			pairs[pair] = i
		}
	}
	return false
}

func containsRepeatingLetterWithOneBetween(str string) bool {
	for i := range len(str) - 2 {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
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
