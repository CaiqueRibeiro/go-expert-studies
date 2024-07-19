package main

import "fmt"

func patternMatching(pattern string, source string) int {
	count := 0
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'y': true,
	}

	n := 0

	for i := 0; i < len(source); i++ {
		letter := rune(source[i])
		_, ok := vowels[letter]

		if string(pattern[n]) == "0" && !ok {
			n = 0
			continue
		}

		if n == len(pattern)-1 {
			count++
			n = 0
			i--
			continue
		}

		n++
	}

	return count
}

func main() {
	result := patternMatching("010", "amazing")
	fmt.Println(result)
}
