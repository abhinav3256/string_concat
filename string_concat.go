package main

import (
	"fmt"
	"strings"
)

func main() {
	var word1 = []string{"ab", "c"}
	var word2 = []string{"a", "bc"}

	var a = strings.Join(word1[:], "")
	var b = strings.Join(word2[:], "")

	fmt.Println(a == b)
}
