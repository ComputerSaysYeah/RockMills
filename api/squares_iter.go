package api

import (
	"github.com/ComputerSaysYeah/RookMills/speed"
)

// SquaresIterator an iterator backed by an expandable RingBuffer, which can be pooled/recycled; so it can be reused
type SquaresIterator interface {
	speed.Iterator[Square]
	speed.Recyclable
	Add(Square) SquaresIterator
	AddIfValid(Square) SquaresIterator
}

type squaresIter struct {
	squares  speed.ExpRingBuffer[Square]
	returned func(any)
}

func NewSquaresIterator() SquaresIterator {
	return &squaresIter{squares: speed.NewExpRingBuffer[Square](32)}
}

func (s *squaresIter) Reset() {
	s.squares.Reset()
}

func (s *squaresIter) Return() {
	s.returned(s)
}

func (s *squaresIter) SetReturnerFn(returner func(any)) {
	s.returned = returner
}

func (s *squaresIter) HasNext() bool {
	return !s.squares.IsEmpty()
}

func (s *squaresIter) Next() Square {
	return s.squares.Pop()
}

func (s *squaresIter) Add(square Square) SquaresIterator {
	if s.squares.IsFull() {
		s.squares.ExpandBy(s.squares.Capacity())
	}
	s.squares.Push(square)
	return s
}

func (s *squaresIter) AddIfValid(square Square) SquaresIterator {
	if square.IsValid() {
		s.Add(square)
	}
	return s
}
