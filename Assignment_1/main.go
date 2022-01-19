package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//odd_even_slice := []int{3, 57, 354, 463, 645, 456, 3256, 46, 2134, 354354, 472547}
	odd_even_slice := []int{}

	for i := 0; i <= 10; i++ {
		odd_even_slice = append(odd_even_slice, rand.Intn(100))
	}

	for _, value := range odd_even_slice {
		if value%2 == 0 {
			fmt.Printf("%v is even.\n", value)
		} else {
			fmt.Printf("%v is odd.\n", value)
		}

	}
}
