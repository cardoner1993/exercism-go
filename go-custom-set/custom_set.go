package stringset

import (
	"fmt"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	items map[string]struct{}
}

func New() Set {
	return Set{
		items: make(map[string]struct{}),
	}
}

func NewFromSlice(l []string) Set {
	s := Set{
		items: make(map[string]struct{}, len(l)),
	}
	for _, item := range l {
		s.items[item] = struct{}{}
	}
	return s
}

func (s Set) String() string {
	if len(s.items) == 0 {
		return "{}"
	}

	var elements []string
	for item := range s.items {
		elements = append(elements, fmt.Sprintf("%q", item))
	}

	return "{" + strings.Join(elements, ", ") + "}"
}

func (s Set) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s Set) Has(elem string) bool {
	_, exists := s.items[elem]
	if exists {
		return true
	}
	return false
}

func (s Set) Add(elem string) {
	_, exists := s.items[elem]
	if !exists {
		s.items[elem] = struct{}{}
	}
}

func Subset(s1, s2 Set) bool {
	// s1 is a subset of s2 if every element in s1 is also in s2
	for item := range s1.items {
		if _, exists := s2.items[item]; !exists {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for item := range s1.items {
		if _, exists := s2.items[item]; exists {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.items) == 0 && len(s2.items) == 0 {
		return true
	}
	if len(s1.items) != len(s2.items) {
		return false
	}
	for item := range s1.items {
		if _, exists := s2.items[item]; !exists {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	s := Set{items: make(map[string]struct{})}
	for item := range s1.items {
		if _, exists := s2.items[item]; exists {
			s.items[item] = struct{}{}
		}
	}
	return s
}

func Difference(s1, s2 Set) Set {
	s := Set{items: make(map[string]struct{})}
	for item := range s1.items {
		if _, exists := s2.items[item]; !exists {
			s.items[item] = struct{}{}
		}
	}
	return s
}

func Union(s1, s2 Set) Set {
	for item := range s1.items {
		if _, exists := s2.items[item]; !exists {
			s2.items[item] = struct{}{}
		}
	}
	return s2
}
