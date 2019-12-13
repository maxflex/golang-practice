package main

import (
	"fmt"
	"math/rand"
)

func insertionSort(slice []int) {
	n := len(slice)
	for i := 1; i < n; i++ {
		// fmt.Printf("Iteration %d:\n", i)
		for j := i; j > 0; j-- {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			} else {
				break
			}
		}
		// fmt.Println(slice)
	}
}

func generateSlice(size int) []int {
	slice := make([]int, size)

	for i := range slice {
		slice[i] = rand.Intn(100) + 1
	}

	return slice
}

func main() {
	sl := generateSlice(10)
	fmt.Println(sl)
	insertionSort(sl)
	fmt.Println(sl)
}
