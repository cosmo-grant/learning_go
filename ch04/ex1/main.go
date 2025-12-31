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
	fmt.Println(s)
}
