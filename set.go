package graphs

// A Set is a container that contains each element just once.
type Set map[interface{}]struct{}

// NewSet creates a new empty set.
func NewSet() *Set {
	return &Set{}
}

// NewSetWithElements creates a new set with the given
// arguments as elements.
func NewSetWithElements(elements ...interface{}) *Set {
	set := NewSet()
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

// Add adds an element to the set. It returns true if the
// element has been added and false if the set already contained
// that element.
func (s *Set) Add(element interface{}) bool {
	_, exists := (*s)[element]
	(*s)[element] = struct{}{}
	return !exists
}

// Len returns the number of elements.
func (s *Set) Len() int {
	return len(*s)
}

// Equals returns whether the given set is equal to the receiver.
func (s *Set) Equals(s2 *Set) bool {
	if s2 == nil || s.Len() != s2.Len() {
		return false
	}

	for element, _ := range *s {
		if !s2.Contains(element) {
			return false
		}
	}

	return true
}

// Contains returns whether the set contains the given element.
func (s *Set) Contains(element interface{}) bool {
	_, exists := (*s)[element]
	return exists
}

// Merge adds the elements of the given set to the receiver.
func (s *Set) Merge(s2 *Set) {
	if s2 == nil {
		return
	}

	for element, _ := range *s2 {
		s.Add(element)
	}
}
