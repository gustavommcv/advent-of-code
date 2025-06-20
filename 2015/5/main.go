package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var forbiddenStrings [4]string = [...]string{"ab", "cd", "pq", "xy"}
var vowels [5]string = [...]string{"a", "e", "i", "o", "u"}

func main() {
	// absolutePath, err := filepath.Abs("input.txt")
	// checkError(err)
	// inputData := scanFile(absolutePath)

	// fmt.Println(getNiceStrings(string(inputData)))
	fmt.Println(isNice2("xyxy"))
	// fmt.Println(isNice2("aaabcdefgaa"))
	// fmt.Println(isNice2("uurcxstgmygtbstg"))
	// fmt.Println(isNice2("ieodomkazucvgmuy"))
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
	// Should contain a pair of any two letters that appears at least twice in the string
	// without overlapping
	// xyxy, aabcdefgaa
	fmt.Println("Has two pairs: ", containsPairWithoutOverlapping(str))

	// Should contain at least one letter which repeats with exactly one letter between them
	// xyx, abcdefeghi, aaa

	return true
}

func containsPairWithoutOverlapping(str string) bool {
	for i, r := range str {
		currentRune := string(r)
		previousRune := ""
		nextRune := ""

		if i > 0 {
			previousRune = string(str[i-1])
		}

		if len(str) > i+1 {
			nextRune = string(str[i+1])
		}

		fmt.Printf("Previous: %s | Current: %s | Next: %s\n", previousRune, currentRune, nextRune)
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
