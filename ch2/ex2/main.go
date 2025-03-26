package main

import "fmt"

func main() {
	const value = 1
	var i int = value
	var f float64 = value
	fmt.Println(i, f)
}

// notable, value i.e. 1 does not have an intrinsic type
// const is just name for literal (not the value of that literal)
// it will become whatever type it is called on to be
// it does have a default type though
//
// literals don't have intrinsic types, but do have default types
// constants are names for literals
