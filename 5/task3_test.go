package main

import (
	"math/rand"
	"sync"
	"testing"
)

type Set struct {
	sync.Mutex
	mm map[int]float32
}

type SetRW struct {
	sync.RWMutex
	mm map[int]float32
}

func NewSet() *Set {
	return &Set{
		mm: map[int]float32{},
	}
}

func NewSetRW() *SetRW {
	return &SetRW{
		mm: map[int]float32{},
	}
}

func (s *Set) Add(numberIndex int, numberValue float32) {
	s.Lock()
	s.mm[numberIndex] = numberValue
	s.Unlock()
}

func (s *SetRW) AddRW(numberIndex int, numberValue float32) {
	s.Lock()
	s.mm[numberIndex] = numberValue
	s.Unlock()
}

func (s *Set) Has(numberIndex int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[numberIndex]
	return ok
}

func (s *SetRW) HasRW(numberIndex int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[numberIndex]
	return ok
}

func BenchmarkSet10WriteM(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.1 {
					set.Add(1, 3.1415926)
				}
				set.Has(1)
			}
		})
	})
}

func BenchmarkSet50WriteM(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.5 {
					set.Add(1, 3.1415926)
				}
				set.Has(1)
			}
		})
	})
}

func BenchmarkSet90WriteM(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.9 {
					set.Add(1, 3.1415926)
				}
				set.Has(1)
			}
		})
	})
}

func BenchmarkSet10WriteRWM(b *testing.B) {
	var set = NewSetRW()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.1 {
					set.AddRW(1, 3.1415926)
				}
				set.HasRW(1)
			}
		})
	})
}

func BenchmarkSet50WriteRWM(b *testing.B) {
	var set = NewSetRW()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.5 {
					set.AddRW(1, 3.1415926)
				}
				set.HasRW(1)
			}
		})
	})
}

func BenchmarkSet90WriteRWM(b *testing.B) {
	var set = NewSetRW()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.9 {
					set.AddRW(1, 3.1415926)
				}
				set.HasRW(1)
			}
		})
	})
}
