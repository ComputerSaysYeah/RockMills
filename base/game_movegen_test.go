package base

import (
	. "github.com/ComputerSaysYeah/RookMills/api"
	"github.com/ComputerSaysYeah/RookMills/speed"
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
	_ = game.FromFEN("r7/4r3/1qbk1n2/3p4/4P3/2N1KBQ1/3R4/7R w - - 0 1")
	containsMovesExactly(t, validMoves(game), "d2d4", "e3e2", "e3f4", "e3d3") // only options is about escaping the check

	//game.Move(ParseMove("d2d4")) // r7/4r3/1qbk1n2/3p4/4P3/2NK1BQ1/3R4/7R b - - 0 1
	//containsMovesExactly(t, validMoves(game), "d6d7", "d6e6", "d6c5", "e7e5") // only options is about escaping the check
}

func TestMoveGen_RookKing(t *testing.T) {
	game := givenGame()
	_ = game.FromFEN("2r5/8/8/3k4/3K4/8/8/4R3 w - - 0 1")
	containsMovesExactly(t, validMoves(game), "d4d3", "d4e3", "d4d5") // d4d5=king eats king ... hmm... fair enough
	game.SetMoveNext(Black)
	containsMovesExactly(t, validMoves(game), "d5d6", "d5d4", "d5c6")
}

func TestMoveGen_Kings(t *testing.T) {
	game := givenGame()
	_ = game.FromFEN("8/8/3k4/8/3K4/8/8/8 w - - 0 1")
	containsMovesExactly(t, validMoves(game), "d4e4", "d4e3", "d4d3", "d4c3", "d4c4")
	game.SetMoveNext(Black)
	containsMovesExactly(t, validMoves(game), "d6d7", "d6e7", "d6e6", "d6c6", "d6c7")
}

func TestMoveGen_Pawns(t *testing.T) {
	game := givenGame()
	_ = game.FromFEN("7k/8/8/3p4/4P3/8/8/K7 w - - 0 1")
	containsMovesExactly(t, validMoves(game) /* king */, "a1a2", "a1b2", "a1b1" /* pawn */, "e4e5", "e4d5")
	game.SetMoveNext(Black)
	containsMovesExactly(t, validMoves(game) /* king */, "h8h7", "h8g7", "h8g8" /* pawn */, "d5e4", "d5d4")
}

func TestMoveGen_PawnsPromotion(t *testing.T) {
	game := givenGame()
	_ = game.FromFEN("7k/4P3/8/8/8/8/3p4/K7 w - - 0 1")
	containsMovesExactly(t, validMoves(game) /* king */, "A1A2", "A1B2", "A1B1" /* pawn */, "E7E8Q", "E7E8B", "E7E8N", "E7E8R")
	game.SetMoveNext(Black)
	containsMovesExactly(t, validMoves(game) /* king */, "H8H7", "H8G7", "H8G8" /* pawn */, "D2D1q", "D2D1b", "D2D1n", "D2D1r")
}

// ---------------------------------------------------------------------------------------------------------------------------------------

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
