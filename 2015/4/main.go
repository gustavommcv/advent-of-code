package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix(hash([]byte("abcdef609043")), "00000"))
	fmt.Println(hash([]byte("abcdef609043")))
	fmt.Println(mine([]byte("abcdef")))
	fmt.Println(mine([]byte("pqrstuv")))
	fmt.Println(mine([]byte("yzbqklnj")))
	fmt.Println(minePart2([]byte("yzbqklnj")))
}

func mine(input []byte) (string, int) {
	i := 1
	found := false
	result := ""
	for !found {
		result = hash(fmt.Appendf(nil, "%s%d", input, i))
		if strings.HasPrefix(result, "00000") {
			found = true
			break
		}
		i++
	}

	return result, i
}

func minePart2(input []byte) (string, int) {
	i := 1
	found := false
	result := ""
	for !found {
		result = hash(fmt.Appendf(nil, "%s%d", input, i))
		if strings.HasPrefix(result, "000000") {
			found = true
			break
		}
		i++
	}

	return result, i
}

func hash(input []byte) string {
	return fmt.Sprintf("%x", md5.Sum(input))
}
