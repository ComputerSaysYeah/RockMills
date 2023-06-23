package base

import (
	. "github.com/ComputerSaysYeah/RookMills/api"
	"github.com/ComputerSaysYeah/RookMills/speed"
	"testing"
)

func TestGame_Start(t *testing.T) {
	bp := speed.NewLeanPool(16, NewBoardB)
	g := NewGame(bp)

	if g.MoveNo() != 1 {
		t.Fatal()
	}
	if g.MoveNext() != White {
		t.Fatal()
	}
}
