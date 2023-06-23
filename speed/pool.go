package speed

import (
	"log"
	"reflect"
)

type Recyclable interface {
	Reset()
	Return()
	SetReturnerFn(func(any))
}

type Pool[v Recyclable] interface {
	Created() int
	Flying() int
	Lease() v
	Release(v)
}

// leanPool implements a non-thread safe pool
type leanPool[v Recyclable] struct {
	objects ExpRingBuffer[v]
	factory func() v
	created int
}

func NewLeanPool[v Recyclable](initialCapacity int, factory func() v) Pool[v] {
	return &leanPool[v]{objects: NewExpRingBuffer[v](initialCapacity), factory: factory, created: 0}
}

func (lp *leanPool[v]) Created() int {
	return lp.created
}

func (lp *leanPool[v]) Flying() int {
	return lp.created - lp.objects.Used()
}

func (lp *leanPool[v]) Lease() v {
	if lp.objects.Empty() {
		elem := lp.factory()
		elem.SetReturnerFn(func(val any) { lp.Release(val.(v)) })
		lp.objects.Push(elem)
		lp.created++
		if lp.created%lp.objects.Capacity() == 0 {
			log.Printf("pool of %v has created %d elements\n", reflect.TypeOf(elem), lp.created)
		}
	}
	return lp.objects.Pop()
}

func (lp *leanPool[v]) Release(val v) {
	if lp.objects.Full() {
		log.Printf("pool of %v ringbuffer is expanding from %d to %d\n", reflect.TypeOf(val), lp.objects.Capacity(), lp.objects.Capacity()*2)
		lp.objects.ExpandBy(lp.objects.Capacity() * 2)
	}
	val.Reset()
	lp.objects.Push(val)
}
