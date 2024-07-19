package main

import "fmt"

func arrayManipulation(a []int) []int {
	b := make([]int, len(a))

	for i, _ := range a {
		if i > 0 {
			b[i] += a[i-1]
		}

		b[i] += a[i]

		if i < len(a)-1 {
			b[i] += a[i+1]
		}
	}

	return b
}

func main() {
	a := []int{4, 0, 1, -2, 3}
	b := arrayManipulation(a)
	fmt.Println(b)
}
