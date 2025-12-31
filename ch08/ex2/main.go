package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type MyInt int

func (mi MyInt) String() string {
	return strconv.Itoa(int(mi))
}

type MyFloat float64

func (mf MyFloat) String() string {
	return strconv.FormatFloat(float64(mf), 'f', 3, 64)
}

func Print[T Printable](p T) {
	fmt.Println(p)
}

func main() {
	mi := MyInt(123)
	mf := MyFloat(456.7890123)
	Print(mi) // 123
	Print(mf) // 456.789
}
