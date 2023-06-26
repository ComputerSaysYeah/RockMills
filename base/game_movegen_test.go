package base

import (
	. "github.com/ComputerSaysYeah/RookMills/api"
	"github.com/ComputerSaysYeah/RookMills/speed"
	"testing"
)

func TestMoveGen_Start(t *testing.T) {
	game := NewGame(
		speed.NewLeanPool(16, NewBoardB),
		speed.NewLeanPool(16, NewMovesIterator))
	movesIter := game.ValidMoves()
	moves := movesIterToMap(movesIter)
	movesIter.Return()

	containsMoves(t, moves, "a2a4", "b2b4", "c2c4", "d2d4", "e2e4", "f2f4", "g2g4", "h2h4", "b1a3", "b1c3", "g1f3", "g1h3")
}

func containsMoves(t *testing.T, moves map[Move]bool, expected ...string) {

	for _, expect := range expected {
		move := ParseMove(expect)
		if !moves[move] {
			t.Fatalf("The move %s is not contained in %v", move.String(), keys(moves))
		}

	}

}

func movesIterToMap(moves speed.Iterator[Move]) map[Move]bool {
	ans := map[Move]bool{}
	for moves.HasNext() {
		ans[moves.Next()] = true
	}
	return ans
}

func keys(moves map[Move]bool) []string {
	ans := []string{}
	for k, _ := range moves {
		ans = append(ans, k.String())
	}
	return ans
}
