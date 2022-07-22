package hash_set

// Set is a collection type with unique elements
type Set[T comparable] struct {
	els map[T]struct{}
}

// New instantiates an empty set
func New[T comparable]() *Set[T] {
	return &Set[T]{make(map[T]struct{})}
}

// Of instantiates a set with the given elements
func Of[T comparable](els ...T) *Set[T] {
	s := New[T]()
	s.Add(els...)
	return s
}

// Empty returns true if a set is empty
func (s *Set[T]) Empty() bool {
	return len(s.els) == 0
}

// Has returns true if a given element is in a set
func (s *Set[T]) Has(el T) bool {
	_, ok := s.els[el]
	return ok
}

// Subset returns true if the set is a subset of another set
func (s *Set[T]) Subset(s2 *Set[T]) bool {
	for el := range s.els {
		if !s2.Has(el) {
			return false
		}
	}
	return true
}

// Add adds an element to the set, optionally supports chaining
func (s *Set[T]) Add(els ...T) *Set[T] {
	for _, el := range els {
		s.els[el] = struct{}{}
	}
	return s
}

// Remove removes an element from the set, optionally supports chaining
func (s *Set[T]) Remove(els ...T) *Set[T] {
	for _, el := range els {
		delete(s.els, el)
	}
	return s
}

// Len returns the number of elements in the set
func (s *Set[T]) Len() int {
	return len(s.els)
}

// Slice returns the elements of the set as a slice
func (s *Set[T]) Slice() []T {
	slice := make([]T, len(s.els))
	i := 0
	for el := range s.els {
		slice[i] = el
		i++
	}
	return slice
}

// Foreach executes a callback function with every member of the set
func (s *Set[T]) Foreach(fn func(T)) *Set[T] {
	for el := range s.els {
		fn(el)
	}
	return s
}

// Filter returns the subset of a set that pass a given a filter func
func (s *Set[T]) Filter(fn func(T) bool) *Set[T] {
	s2 := New[T]()
	for el := range s.els {
		if fn(el) {
			s2.Add(el)
		}
	}
	return s2
}

// Intersection returns the set of elements common to both sets
func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	i := New[T]()
	s.Foreach(func(el T) {
		if s2.Has(el) {
			i.Add(el)
		}
	})
	return i
}

// Union returns the combined elements of both sets
func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	u := New[T]()
	s.Foreach(func(el T) {
		u.Add(el)
	})
	s2.Foreach(func(el T) {
		u.Add(el)
	})
	return u
}

// Difference returns the set of elements not in the other set
func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	d := New[T]()
	s.Foreach(func(el T) {
		if !s2.Has(el) {
			d.Add(el)
		}
	})
	return d
}
