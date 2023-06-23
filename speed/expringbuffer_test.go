package speed

import (
	"testing"
)

func TestNewExpRingBuffer_Basics(t *testing.T) {
	rb := NewExpRingBuffer[int](0)
	if !rb.Empty() {
		t.Fatal("it should be empty at start")
	}
	if rb.Remaining() != 15 {
		t.Fatal("there should be 15 slots available -one less than the total-")
	}
	if rb.Capacity() != 16 {
		t.Fatal("default capacity when not specified is 16")
	}
	if rb.Used() != 0 {
		t.Fatal("used at start should be 0")
	}
	rb.ExpandBy(16)
	if rb.Capacity() != 32 {
		t.Fatal("new capacity should be 32")
	}
	if rb.Remaining() != 31 {
		t.Fatal("there should be 31 slots available -one less than the total-")
	}
	if rb.Used() != 0 {
		t.Fatal("new used after expansion should still be 0")
	}

	for i := 0; i < 31; i++ {
		if rb.Full() {
			t.Fatal("it should not be full!")
		}
		rb.Push(i)
		if rb.Remaining() != 30-i {
			t.Fatal("there should be the correct amount of available slots")
		}
		if rb.Empty() {
			t.Fatal("it should not be empty!")
		}
	}
	if rb.Remaining() != 0 {
		t.Fatal("it should be full")
	}
	for i := 0; i < 31; i++ {
		if rb.Remaining() != i {
			t.Fatal("availability not as expected")
		}
		if rb.Pop() != i {
			t.Fatalf("did not pop the expected number")
		}
		if rb.Remaining() != i+1 {
			t.Fatal("availability should be expanding")
		}
	}
	if !rb.Empty() {
		t.Fatal("It should be empty!")
	}
}

func TestExpRingBufferSt_ExpandBy(t *testing.T) {
	rb := NewExpRingBuffer[int](20)
	rb.ExpandBy(30)
	if rb.Remaining() != 49 {
		t.Fatal("it should have been expanded")
	}
	if !rb.Empty() {
		t.Fatal("it is definitely empty!")
	}
	for i := 0; i < 40; i++ {
		rb.Push(i)
	}
	rb.ExpandBy(40)
	for i := 0; i < 40; i++ {
		if rb.Pop() != i {
			t.Fatal("it should have pop out the value after expanding")
		}
	}
	if !rb.Empty() {
		t.Fatal("should be empty by now")
	}
	if rb.Remaining() != rb.Capacity()-1 && rb.Capacity() != 89 {
		t.Fatal("capacity is not ok")
	}
}

func TestExpRingBufferSt_Advanced(t *testing.T) {
	write := 0
	read := 0

	rb := NewExpRingBuffer[int](32)
	for i := 0; i < 10000; i++ {

		if rb.Remaining() < 3 {
			for ii := 0; ii < 3; ii++ {
				if rb.Pop() != read {
					t.Fatal("not expected value!")
				}
				read++
			}
		}

		rb.Push(write)
		write++
		rb.Push(write)
		write++

		if rb.Pop() != read {
			t.Fatal("not expected value!")
		}
		read++

	}
}

// --------------------------------------------------------------------------------------------------------------------------------------

func BenchmarkExpRingBufferSt(b *testing.B) {
	rb := NewExpRingBuffer[int](32)
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			rb.Push(i)
		} else {
			rb.Pop()
		}
	}
}
