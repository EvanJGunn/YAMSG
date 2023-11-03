package data

import (
	"errors"
)

// CONTEXT - Meant to use this to do some classic stack based validation on command syntax
// but after seeing that instead of one big command, I am doing many little ones,
// I dont think its necessary after all
// also means this is all untested :/

const maxuint32 = 4294967295

type Stack[T any] interface {
	Push(element T)
	Pop() T
	Peak() (element *T, ok bool)
}

type stack[T any] struct {
	size     uint32
	maxSize  uint32
	elements []*T
}

// NewStack Create a new stack, size will always end up 1 at minimum
func NewStack[T any](size uint32) Stack[T] {
	if size == 0 {
		size = 1
	}
	return &stack[T]{
		size:     0,
		maxSize:  size,
		elements: make([]*T, size),
	}
}

func (s *stack[T]) Push(element T) {
	// increasing size
	if s.size+1 > s.maxSize {
		if s.maxSize == maxuint32 {
			// Not really worried about hitting this one, minecraft has a 32k character limit on command blocks
			panic(errors.New("failed to increase stack size, hit size limit"))
		}

		// 2*max size is greater than the max uint32, set it to the max
		if (maxuint32 - s.maxSize) > s.maxSize {
			s.maxSize = maxuint32
		} else {
			s.maxSize = 2 * s.maxSize
		}

		newElements := make([]*T, s.maxSize)

		// maybe there is a more efficient way... idk, wanted to do this on my own
		// also its a slice so probably under the hood it would get taken care of anyways
		for i := uint32(0); i < s.size; i++ {
			newElements[i] = s.elements[i]
		}
		s.elements = newElements
	}

	// normally I would have to allocate on the heap here... pretty sure go handles this somehow?
	*s.elements[s.size] = element
	s.size += 1
}

func (s *stack[T]) Pop() (element T) {
	element = *s.elements[s.size-1]
	s.size -= 1

	// normally I would free memory here... assuming go will handle this as well
	s.elements[s.size-1] = nil
	return element
}

func (s *stack[T]) Peak() (element *T, ok bool) {
	if s.size > 0 {
		return s.elements[s.size-1], true
	}
	return nil, false
}
