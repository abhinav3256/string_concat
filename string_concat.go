package main

import (
	"fmt"
)

func main() {
	var word1 = []string{"ab", "cd"}
	var word2 = []string{"a", "bc"}
	var a string
	var b string

	for i := 0; i < len(word1); i++ {
		a = word1[i] + a
	}
	for i := 0; i < len(word2); i++ {
		b = word2[i] + b
	}

	fmt.Println(a == b)
}
