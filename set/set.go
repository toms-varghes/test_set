package set

type Set struct {
	// TODO: Protect elements array using a lock.
	// TODO: Make it a map instead of array.
	elements []int
}

func (s *Set) Add(elem int) {
	for _, e := range s.elements {
		if e == elem {
			return
		}
	}
	s.elements = append(s.elements, elem)
}

func (s *Set) Union(s2 *Set) *Set {
	newSet := &Set{}
	for _, e := range s.elements {
		newSet.Add(e)
	}
	for _, e := range s2.elements {
		newSet.Add(e)
	}
	return newSet
}

func (s *Set) GetElements() []int {
	return s.elements
}

func NewSet() *Set {
	return &Set{
		elements: make([]int, 0),
	}
}
