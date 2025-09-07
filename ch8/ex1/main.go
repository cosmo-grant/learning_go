package main

import "fmt"

type IntOrFloat interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

func Double[T IntOrFloat](v T) (r T) {
	return v + v
}

func main() {
	var x int = 1
	var y int8 = 1
	var z uint = 1
	var w float32 = 1
	fmt.Println(Double(x))
	fmt.Println(Double(y))
	fmt.Println(Double(z))
	fmt.Println(Double(w))
}
