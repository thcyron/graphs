package graphs

// A Set is a container that contains each element just once.
type Set[T comparable] map[T]struct{}

// NewSet creates a new empty set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{}
}

// NewSetWithElements creates a new set with the given
// arguments as elements.
func NewSetWithElements[T comparable](elements ...T) *Set[T] {
	set := NewSet[T]()
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

// Add adds an element to the set. It returns true if the
// element has been added and false if the set already contained
// that element.
func (s *Set[T]) Add(element T) bool {
	_, exists := (*s)[element]
	(*s)[element] = struct{}{}
	return !exists
}

// Len returns the number of elements.
func (s *Set[T]) Len() int {
	return len(*s)
}

// Equals returns whether the given set is equal to the receiver.
func (s *Set[T]) Equals(s2 *Set[T]) bool {
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
func (s *Set[T]) Contains(element T) bool {
	_, exists := (*s)[element]
	return exists
}

// Merge adds the elements of the given set to the receiver.
func (s *Set[T]) Merge(s2 *Set[T]) {
	if s2 == nil {
		return
	}

	for element, _ := range *s2 {
		s.Add(element)
	}
}

// Remove removes the given element from the set and returns
// whether the element was removed from the set.
func (s *Set[T]) Remove(element T) bool {
	if _, exists := (*s)[element]; exists {
		delete(*s, element)
		return true
	}
	return false
}

// Any returns any element from the set.
func (s *Set[T]) Any() T {
	for v, _ := range *s {
		return v
	}
	panic("graphs: empty set")
}

// Each executes the given function for each element
// in the set.
func (s *Set[T]) Each(f func(T, func())) {
	stopped := false
	stop := func() { stopped = true }
	for v, _ := range *s {
		f(v, stop)
		if stopped {
			return
		}
	}
}
