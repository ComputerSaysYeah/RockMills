package api

import (
	"testing"
)

func TestSquareStrings(t *testing.T) {
	if ParseSquare("h4") != ColH+Row4 {
		t.Fatal()
	}
	if (ColD + Row6).String() != "D6" {
		t.Fatalf("'%v' is not 'D6'\n", (ColD + Row6).String())
	}
	for i := 0; i < 64; i++ {
		square := Square(i)
		if ParseSquare(square.String()) != square {
			t.Fatal("it won't round-trip ParseSquare() String() ")
		}
	}
}
