package main

import "fmt"

func UpdateSlice(sl []string, st string) {
	sl[len(sl)-1] = st
	fmt.Println("at end of UpdateSlice:", sl) // [a b update]
}

func GrowSlice(sl []string, st string) {
	sl = append(sl, st)
	fmt.Println("at end of GrowSlice:", sl) // [a b update grow]
}

func main() {
	sl := []string{"a", "b", "c"}

	fmt.Println("at start of main:", sl) // [a b c]
	UpdateSlice(sl, "update")
	fmt.Println("in main, after calling UpdateSlice:", sl) // [a b update]
	GrowSlice(sl, "grow")
	fmt.Println("in main, after calling GrowSlice:", sl) // [a b update]
}

// Slices are implemented as pointers to structs with 3 fields:
// len int, cap int, pointer to memory block.
// When you call a function with the slice as argument, you copy that struct.
// Assigning in UpdateSlice() changes the bits in the pointed at memory block.
// main's slice points at the same memory block.
// So it sees the change.
// Appending to the slice in GrowSlice() exceeds the capacity, so a new memory block is allocated and pointed at.
// main's slice still points at the original memory block, so it doesn't see the change.
// Even if appending hadn't exceeded the capacity, the appended item would be beyond main's slice's length, so inaccessible.
