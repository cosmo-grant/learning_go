// Package add provides a function for adding two ints.
//
// That's it.
// For more on addition, see [this article].
//
// [this article]: https://www.mathisfun.com/numbers/addition.html
package add

import "golang.org/x/exp/constraints"

// Number is a type constraint for numeric types.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns the sum of two ints.
func Add[T Number](i T, j T) T {
	return i + j
}
