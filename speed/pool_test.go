package speed

import (
	"testing"
)

type PoolObj struct {
	value    int
	returner func(any)
}

func (o *PoolObj) Reset() {
	//o.value = 0
	//in a normal object, I want this to be reset but I just want to check they are being recycled
}

func (o *PoolObj) Return() {
	o.returner(o)
}

func (o *PoolObj) SetReturnerFn(returner func(any)) {
	o.returner = returner
}

func NewPoolObj(value int) PoolObj {
	return PoolObj{value: value}
}

func TestNewLeanPool_Basics(t *testing.T) {
	next := 0
	pool := NewLeanPool[*PoolObj](32, func() *PoolObj {
		next++
		obj := NewPoolObj(next)
		return &obj
	})

	var flying []*PoolObj

	for i := 0; i < 1024; i++ {
		if pool.Created() != i {
			t.Fatalf("pool should have created %d values, not %d\n", i, pool.Created())
		}
		if pool.Flying() != i {
			t.Fatalf("pool should have  %d flying objects, not %d\n", i, pool.Flying())
		}
		flying = append(flying, pool.Lease())
	}

	for len(flying) > 0 {
		flying[0].Return()
		flying = flying[1:]
		if pool.Flying() != len(flying) {
			t.Fatal("flying objects should get updated")
		}
	}

	if pool.Flying() != 0 {
		t.Fatal("there should not be more flying objects")
	}

	for i := 0; i < 1024; i++ {
		obj := pool.Lease()
		if obj.value != i+1 {
			t.Fatalf("it is not reusing the objects, or not reusing them in order, i=%d, obj=%v\n", i, obj)
		}
	}

}
