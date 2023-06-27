package base

import (
	. "github.com/ComputerSaysYeah/RookMills/api"
	"testing"
)

func TestGame_Start(t *testing.T) {
	g := givenGame()

	if g.MoveNo() != 1 {
		t.Fatal()
	}
	if g.MoveNext() != White {
		t.Fatal()
	}
}
