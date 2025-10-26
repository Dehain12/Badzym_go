package main

import "fmt"

func main() {
	fmt.Println(sort([]int{5, 1, 13, -1, -5}))
}

func sort(badzim []int) []int {
	for i := 0; i < len(badzim)-1; i++ {
		for j := i + 1; j < len(badzim); j++ {
			if badzim[j] > badzim[i] {
				badzim[j], badzim[i] = badzim[i], badzim[j]
			}
		}
	}
	return badzim
}
