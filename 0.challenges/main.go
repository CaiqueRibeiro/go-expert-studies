package main

import "fmt"

func main() {
	sports := make([]string, 5) // um slice
	sports[0] = "ski"
	sports[1] = "surf"
	sports[2] = "swim"
	sports[3] = "sail"
	sports[4] = "sumo wrestling"

	xs := sports[1:3:3]
	xs[0] = "CHANGED"

	inspectSlice(sports)
	inspectSlice(xs)
}

func inspectSlice(xs []string) {
	fmt.Printf("len=%d \ncap=%v \n", len(xs), cap(xs))
	for i := range xs {
		fmt.Printf("%p \t %v \n", &xs[i], xs[i])
	}
}
