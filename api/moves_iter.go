package api

import (
	"github.com/ComputerSaysYeah/RookMills/speed"
)

// MovesIterator an iterator backed by an expandable RingBuffer, which can be pooled/recycled; so it can be reused
type MovesIterator interface {
	speed.Iterator[Move]
	speed.Recyclable
	Add(Move)
}

type movesIter struct {
	MovesIterator
	moves    speed.ExpRingBuffer[Move]
	returned func(any)
}

func NewMovesIterator() MovesIterator {
	return &movesIter{moves: speed.NewExpRingBuffer[Move](32)}
}

func (m *movesIter) Reset() {
	m.moves.Reset()
}

func (m *movesIter) Return() {
	m.returned(m)
}

func (m *movesIter) SetReturnerFn(returner func(any)) {
	m.returned = returner
}

func (m *movesIter) HasNext() bool {
	return !m.moves.IsEmpty()
}

func (m *movesIter) Next() Move {
	return m.moves.Pop()
}

func (m *movesIter) Add(move Move) {
	if m.moves.IsFull() {
		m.moves.ExpandBy(m.moves.Capacity())
	}
	m.moves.Push(move)
}
