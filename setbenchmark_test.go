package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// реализуйте быстрое множество
type IntSet map[int]struct{}

func MakeIntSet() IntSet {
	return IntSet{}
}

func (s IntSet) Contains(elem int) bool {
	if _, ok := s[elem]; ok {
		return true
	}
	return false
}

func (s IntSet) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	s[elem] = struct{}{}
	return true
}

func BenchmarkIntSet(b *testing.B) {
	rand.Seed(0)
	for _, size := range []int{1, 10, 100, 1000, 10000, 100000} {
		set := randomSet(size)
		name := fmt.Sprintf("Contains-%d", size)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				elem := rand.Intn(100000)
				set.Contains(elem)
			}
		})

	}
}

func randomSet(size int) IntSet {
	set := MakeIntSet()
	for i := 0; i < size; i++ {
		n := rand.Intn(100000)
		set.Add(n)
	}
	return set
}
