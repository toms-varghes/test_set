package main

import (
	"fmt"

	"github.com/toms-varghes/test_set/set"
)

func main() {
	s1 := set.NewSet()
	fmt.Printf("S1: %v\n", s1.GetElements())

	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s1.Add(2)
	s1.Add(3)
	fmt.Printf("S1: %v\n", s1.GetElements())

	s2 := set.NewSet()
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)
	s2.Add(6)
	s2.Add(7)
	fmt.Printf("S2: %v\n", s2.GetElements())
	union := s1.Union(s2)
	fmt.Printf("S1 U S2: %v\n", union.GetElements())

}
