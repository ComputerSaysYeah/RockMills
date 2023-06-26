package speed

import (
	"log"
)

type ExpRingBuffer[t any] interface {
	Push(t)
	Pop() t
	Capacity() int
	Remaining() int
	Used() int
	IsEmpty() bool
	IsFull() bool
	Reset()
	ExpandBy(int)
}

type expRingBufferSt[t any] struct {
	buffer []t
	head   int // next position where to push
	tail   int // next position where to pop
}

// NewExpRingBuffer an expandable ring buffer which enough effort has been placed in order to make it non-allocating unless when it is expanded
func NewExpRingBuffer[t any](initialCapacity int) ExpRingBuffer[t] {
	if initialCapacity <= 0 {
		log.Println("setting default ExpRingBuffer to 16 as zero or negative was given.")
		initialCapacity = 16
	}
	return &expRingBufferSt[t]{buffer: make([]t, initialCapacity), head: 0, tail: initialCapacity - 1}
}

func (e *expRingBufferSt[t]) Push(v t) {
	if e.IsFull() {
		log.Fatalln("ExpRingBuffer should not be full in order to Push() an element. check IsFull() first.")
	}
	e.buffer[e.head] = v
	e.head++
	if e.head == e.Capacity() {
		e.head = 0
	}
}

func (e *expRingBufferSt[t]) Pop() t {
	if e.IsEmpty() {
		log.Fatalln("ExpRingBuffer should not be empty in order to Pop() an element. check IsEmpty() first.")
	}
	e.tail++
	if e.tail == e.Capacity() {
		e.tail = 0
	}
	return e.buffer[e.tail]
}

func (e *expRingBufferSt[t]) Reset() {
	e.head = 0
	e.tail = e.Capacity() - 1
}

func (e *expRingBufferSt[t]) Capacity() int {
	return len(e.buffer)
}

func (e *expRingBufferSt[t]) Remaining() int {
	return (e.tail - e.head + e.Capacity()) % e.Capacity()
}

func (e *expRingBufferSt[t]) Used() int {
	return e.Capacity() - e.Remaining() - 1
}

func (e *expRingBufferSt[t]) IsEmpty() bool {
	return (e.tail+1)%e.Capacity() == e.head
}

func (e *expRingBufferSt[t]) IsFull() bool {
	return e.tail == e.head
}

func (e *expRingBufferSt[t]) ExpandBy(more int) {
	newBuffer := make([]t, e.Capacity()+more)
	next := 0
	for !e.IsEmpty() {
		newBuffer[next] = e.Pop()
		next++
	}
	e.buffer = newBuffer
	e.head = next
	e.tail = len(e.buffer) - 1
}
