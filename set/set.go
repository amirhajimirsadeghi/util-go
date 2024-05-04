package set

import "golang.org/x/exp/maps"

// Set allows you to create a typed Set when you want to keep track of a pool of
// objects, but don't care about the order.
type Set[T comparable] map[T]struct{}

var Exists = struct{}{}

func New[T comparable]() Set[T] {
	s := make(Set[T])
	return s
}

// Contains returns true if a given value is currently in the set
func (s Set[T]) Contains(value T) bool {
	if s == nil {
		s = make(Set[T])
	}

	_, c := s[value]
	return c
}

func (s Set[T]) Add(value T) {
	if s == nil {
		s = make(Set[T])
	}

	s[value] = Exists
}

// AddSet copies the contents of another set into the recieving set
func (s Set[T]) AddSet(input Set[T]) {
	for _, value := range input.Slice() {
		s.Add(value)
	}
}

func (s Set[T]) Copy() Set[T] {
	newSet := New[T]()

	for k := range s {
		newSet.Add(k)
	}

	return newSet
}

// Remove deletes an item from the Set.  No effect if it is already
// absent
func (s Set[T]) Remove(value T) {
	if s == nil {
		s = make(map[T]struct{})
	}

	delete(s, value)
}

// Size returns the current number of elements stored in the Set()
// This does NOT calculate the amount of memory it's using!
func (s Set[T]) Size() int {
	if s == nil {
		s = make(map[T]struct{})
	}

	return len(s)
}

func (s Set[T]) Empty() bool {
	if s == nil {
		s = make(Set[T])
	}

	return s.Size() == 0
}

// Slice returns a Slice representation of the Set's keys
// This is useful when you'd like to iterate over the Set's contents
// e.g.
// for item := range mySet.Slice() {}
func (s Set[T]) Slice() []T {
	if s == nil {
		s = make(Set[T])
	}

	return maps.Keys(s)
}

// DeleteIntersection Remove any elements that exist in both input sets
// from both input sets.  Will return the set of elements
// that were ejected from each set
func DeleteIntersection[K comparable](set1, set2 *Set[K]) *Set[K] {
	removed := New[K]()

	for _, val := range set1.Slice() {
		if set2.Contains(val) {
			set1.Remove(val)
			set2.Remove(val)
			removed.Add(val)
		}
	}

	return &removed
}

// DeleteIntersectionWithLimit Remove up to N elements that exist in both input sets
// from both input sets.  Will return the set of elements
// that were ejected from each set
func DeleteIntersectionWithLimit[K comparable](set1, set2 *Set[K], limit uint) *Set[K] {
	removed := New[K]()

	for _, val := range set1.Slice() {
		if limit == 0 {
			break
		}

		if set2.Contains(val) {
			set1.Remove(val)
			set2.Remove(val)
			removed.Add(val)

			limit -= 1
		}
	}

	return &removed
}
