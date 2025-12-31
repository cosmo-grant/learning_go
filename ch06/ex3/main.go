package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{}
	for range 10_000_000 {
		people = append(people, Person{"cosmo", "grant", 33})
	}
}

// GOGC=default: 0.48s user 0.27s system 118% cpu 0.638 total
// GOGC=200:     0.34s user 0.26s system 134% cpu 0.450 total
// GOGC=300:     0.25s user 0.20s system 142% cpu 0.319 total
// GOGC=400:     0.30s user 0.27s system 64% cpu 0.873 total

// export GODEBUG=gctrace=1
// go run ch6/ex3 2>&1 | wc -l  # think this counts the number of collections
// GOGC=default: 19
// GOGC=200: 11
// GOGC=300: 7
// GOGC=400: 6
// which makes sense: the higher GOGC, the faster the heap is allowed to grow between collections
