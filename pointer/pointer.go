package pointer

// A few functions that I've found helpful when working with Pointers

// Of is a helper routine that allocates a new any value
// to store v and returns a pointer to it.
// This is useful when you want to initialize a struct with pointers in it.
func Of[Value any](v Value) *Value {
	return &v
}

// Equal compares pointers that might be nil, without risking
// a nil dereference
func Equal[Value comparable](a, b *Value) bool {
	return a == b || (a != nil && b != nil && *a == *b)
}

// SDeref (Safe Dereference) is a helper function to access a pointer that *might*
// be nil, and you'd like to return an empty value if it is
// This comes up a lot for "might be set" string parameters
func SDeref[T any](p *T) T {
	if p == nil {
		var v T
		return v
	}
	return *p
}
