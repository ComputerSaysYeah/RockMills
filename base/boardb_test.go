package base

import (
	. "github.com/ComputerSaysYeah/RookMills/api"
	"log"
	"testing"
)

func TestBoardB_Reset(t *testing.T) {
	bb := NewBoardB()
	for i := 0; i < 64; i++ {
		bb.Set(Square(i), Rook)
	}
	bb.Reset()
	for i := 0; i < 64; i++ {
		if bb.Get(Square(i)) != Empty {
			t.Fail()
		}
	}
}

func TestBoardB_Hash(t *testing.T) {
	bb := NewBoardB()
	for i := 0; i < 64; i++ {
		before := bb.Hash()
		bb.Set(Square(i), Rook)
		after := bb.Hash()
		if before == after {
			t.Fail()
		}
	}
}

func TestBoardB_CopyFrom(t *testing.T) {
	a := NewBoardB()
	a.Set(Row2+ColD, Rook)

	b := NewBoardB()
	b.CopyFrom(&a)

	if a.Hash() != b.Hash() {
		log.Fatal()
	}
	if b.Get(Row2+ColD) != Rook {
		log.Fatal()
	}
}

func TestBoardB_String(t *testing.T) {
	a := NewBoardB()
	a.SetStartingPieces()
	log.Printf("\n%v", a)
}

// --------------------------------------------------------------------------------------------------------------------------------------

func BenchmarkBoardBReset(b *testing.B) {
	bb := NewBoardB()
	for i := 0; i < b.N; i++ {
		bb.Reset()
	}
}

func BenchmarkBoardB_Hash(b *testing.B) {
	bb := NewBoardB()
	for i := 0; i < b.N; i++ {
		bb.Hash()
	}
}
