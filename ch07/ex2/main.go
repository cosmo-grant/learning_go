package main

import (
	"fmt"
	"sort"
)

type Team struct {
	name    string
	players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) MatchResult(t1 string, s1 int, t2 string, s2 int) {
	if s1 > s2 {
		l.Wins[t1]++
	} else if s2 > s1 {
		l.Wins[t2]++
	}
}

func (l League) Ranking() []string {
	ranking := make([]string, 0, len(l.Teams))
	for _, team := range l.Teams {
		ranking = append(ranking, team.name)
	}
	sort.Slice(ranking, func(i, j int) bool {
		return l.Wins[ranking[i]] < l.Wins[ranking[j]]
	})
	return ranking
}

func main() {
	t1 := Team{"A", []string{"a1", "a2", "a3"}}
	t2 := Team{"B", []string{"b1", "b2", "b3"}}
	t3 := Team{"C", []string{"c1", "c2"}}

	l := League{[]Team{t1, t2, t3}, map[string]int{}}

	l.MatchResult("A", 0, "B", 1)
	l.MatchResult("A", 0, "C", 1)
	l.MatchResult("B", 1, "C", 0)
	// A has no wins, B has 2, C has 1

	fmt.Println(l.Ranking())

}
