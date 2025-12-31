package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := []int{}
	for range 100 {
		s = append(s, rand.Intn(100))
	}
	for _, v := range s {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
