package base

import (
	. "github.com/ComputerSaysYeah/RookMills/api"
	"github.com/ComputerSaysYeah/RookMills/speed"
	"github.com/ComputerSaysYeah/RookMills/util"
	"testing"
)

func TestMoveGen_Start(t *testing.T) {
	game := givenGame()
	containsMovesExactly(t, validMoves(game),
		"b1a3", "b1c3", "g1f3", "g1h3", // knights
		"a2a3", "b2b3", "c2c3", "d2d3", "e2e3", "f2f3", "g2g3", "h2h3", //pawns 1 move
		"a2a4", "b2b4", "c2c4", "d2d4", "e2e4", "f2f4", "g2g4", "h2h4") //pawns 2 moves
}

func TestMoveGen_2ndMove(t *testing.T) {
	game := givenGame()
	game.Move(ParseMove("d2d4"))
	containsMovesExactly(t, validMoves(game), "b8a6", "b8c6", "g8f6", "g8h6", // knights
		"a7a6", "b7b6", "c7c6", "d7d6", "e7e6", "f7f6", "g7g6", "h7h6", // pawns 1 move
		"a7a5", "b7b5", "c7c5", "d7d5", "e7e5", "f7f5", "g7g5", "h7h5") // pawns 1 move
}

func TestMoveGen_Advanced(t *testing.T) {
	game := givenGame() // https://lichess.org/editor/r7/4r3/1qbk1n2/3p4/4P3/2N1KBQ1/3R4/7R_w_-_-_0_1?color=white
	_ = util.ParseFEN(game, "r7/4r3/1qbk1n2/3p4/4P3/2N1KBQ1/3R4/7R w - - 0 1")
	containsMovesExactly(t, validMoves(game), "d2d4", "e3e2", "e3f4", "e3d3") // only options is about escaping the check
}

func givenGame() Game {
	return NewGame(
		speed.NewLeanPool(16, NewBoardB),
		speed.NewLeanPool(16, NewMovesIterator))
}

// ---------------------------------------------------------------------------------------------------------------------------------------

func containsMovesExactly(t *testing.T, moves map[Move]bool, expected ...string) {
	containsMoves(t, moves, expected...)
	if len(moves) != len(expected) {
		for _, expect := range expected {
			delete(moves, ParseMove(expect))
		}
		nonMatched := keys(moves)
		t.Fatalf("Some moves have not been matched: %v", nonMatched)
	}
}

func containsMoves(t *testing.T, moves map[Move]bool, expected ...string) {
	for _, expect := range expected {
		move := ParseMove(expect)
		if !moves[move] {
			t.Fatalf("The move %s is not contained in %v", move.String(), keys(moves))
		}

	}
}

// ---------------------------------------------------------------------------------------------------------------------------------------

func validMoves(game Game) map[Move]bool {
	movesIter := game.ValidMoves()
	defer movesIter.Return()
	return movesIterToMap(movesIter)
}

func movesIterToMap(moves speed.Iterator[Move]) map[Move]bool {
	ans := map[Move]bool{}
	for moves.HasNext() {
		ans[moves.Next()] = true
	}
	return ans
}

func keys(moves map[Move]bool) []string {
	var ans []string
	for k, _ := range moves {
		ans = append(ans, k.String())
	}
	return ans
}
