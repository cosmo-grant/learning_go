package main

import (
	"fmt"
	"strings"
)

type LinkedList[T comparable] struct {
	val  T
	next *LinkedList[T]
}

func (ll LinkedList[T]) String() string {
	var res string

	for {
		res = res + " " + fmt.Sprint(ll.val)
		if ll.next == nil {
			break
		}
		ll = *ll.next
	}

	res = strings.Trim(res, " ")
	return res
}

func (ll *LinkedList[T]) Add(v T) {
	for {
		if ll.next == nil {
			break
		} else {
			ll = ll.next
		}
	}
	ll.next = &LinkedList[T]{val: v, next: nil}
}

func (ll *LinkedList[T]) Index(v T) int {
	for i := 0; ; i++ {
		if ll == nil {
			return -1
		} else if ll.val == v {
			return i
		} else {
			ll = ll.next
		}
	}
}

func (ll *LinkedList[T]) Length() int {
	length := 1
	for {
		if ll.next == nil {
			return length
		} else {
			length++
			ll = ll.next
		}
	}
}

func (ll *LinkedList[T]) Insert(v T, i int) {
	var prev *LinkedList[T]
	for j := 1; j <= i; j++ {
		prev = ll
		ll = ll.next
	}
	if prev == nil {
		*ll = LinkedList[T]{v, &LinkedList[T]{ll.val, ll.next}}
	} else if ll == nil {
		prev.next = &LinkedList[T]{v, nil}
	} else {
		prev.next = &LinkedList[T]{v, ll}
	}
}

func main() {
	ll := LinkedList[int]{0, nil}
	ll.Add(10)
	ll.Add(20)
	fmt.Println("values:", ll)
	ll.Insert(-5, 0)
	ll.Insert(5, 1)
	ll.Insert(15, 3)
	ll.Insert(25, 5)
	fmt.Println("values:", ll)
	for _, v := range []int{0, 10, 20, 40} {
		fmt.Println("index of", v, "is", ll.Index(v))
	}
	fmt.Println("length:", ll.Length())
}
