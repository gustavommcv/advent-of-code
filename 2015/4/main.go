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
	fmt.Println(mineWithPrefix([]byte("yzbqklnj"), "000000"))
}

func mine(input []byte) (string, int) {
	return mineWithPrefix(input, "00000")
}

func mineWithPrefix(input []byte, prefix string) (string, int) {
	i := 1
	found := false
	result := ""
	for !found {
		result = hash(fmt.Appendf(nil, "%s%d", input, i))
		if strings.HasPrefix(result, prefix) {
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
