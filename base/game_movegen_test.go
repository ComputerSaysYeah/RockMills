package base

import (
	"fmt"
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
func TestMoveGen_PawnsPromotion_NoKings(t *testing.T) {
	game := givenGame()
	_ = game.FromFEN("8/4P3/8/8/8/8/3p4/8 w - - 0 1")
	containsMovesExactly(t, validMoves(game), "E7E8Q", "E7E8B", "E7E8N", "E7E8R")
	game.SetMoveNext(Black)
	containsMovesExactly(t, validMoves(game), "D2D1q", "D2D1b", "D2D1n", "D2D1r")
}

func TestMoveGen_CheckMate(t *testing.T) {
	game := givenGame()
	_ = game.FromFEN("rk5Q/pppp4/8/8/8/8/8/8 w - - 0 1")
	containsMovesExactly(t, validMoves(game), "H8B8",
		"H8H7", "H8H6", "H8H5", "H8H4", "H8H3", "H8H2", "H8H1",
		"H8G8", "H8F8", "H8E8", "H8D8", "H8C8",
		"H8G7", "H8F6", "H8E5", "H8D4", "H8C3", "H8B2", "H8A1")
	game.SetMoveNext(Black)
	containsMovesExactly(t, validMoves(game))
}

func TestMoveGen_Interesting(t *testing.T) {
	game := givenGame()
	// https://twitter.com/ChessMood/status/1678871792683741184a
	_ = game.FromFEN("8/8/6K1/8/5p1p/5p1k/5P1P/6RR w - - 0 1")
	containsMovesExactly(t, validMoves(game), "G6H7", "G6H6", "G6F6", "G6F7", "G6H5", "G6F5", "G6G7", "G6G5", // king moves
		"G1A1", "G1B1", "G1C1", "G1D1", "G1E1", "G1F1", "G1G2", "G1G3", "G1G4", "G1G5") // rook moves
	game.Move(ParseMove("G1G3"))
	containsMovesExactly(t, validMoves(game), "F4G3", "H4G3")
	game.Move(ParseMove("F4G3"))
	containsMovesExactly(t, validMoves(game), "G6H7", "G6H6", "G6F6", "G6F7", "G6H5", "G6F5", "G6G7", "G6G5", // king moves,
		"H1G1", "H1F1", "H1E1", "H1D1", "H1C1", "H1B1", "H1A1", // rook moves
		"H2G3", "F2G3", // Pawns attacks
	)
	game.Move(ParseMove("H2G3"))
	containsMovesExactly(t, validMoves(game), "H3G2", "H3G4")
	game.Move(ParseMove("H3G2"))                                                                              // escaped
	containsMovesExactly(t, validMoves(game), "G6H7", "G6H6", "G6F6", "G6F7", "G6H5", "G6F5", "G6G7", "G6G5", // king moves,
		"H1G1", "H1F1", "H1E1", "H1D1", "H1C1", "H1B1", "H1A1", "H1H2", "H1H3", "H1H4", // rook moves
		"G3G4", "G3H4", // pawn
	)
}

func TestMoveGen_Interesting_MateIn2(t *testing.T) {
	game := givenGame()
	// https://twitter.com/ChessMood/status/1678871792683741184a
	_ = game.FromFEN("8/8/6K1/8/5p1p/5p1k/5P1P/6RR w - - 0 1")
	game.Move(ParseMove("G1G4"))
	containsMovesExactly(t, validMoves(game), "H3G4")
	game.Move(ParseMove("H3G4"))
	containsMovesExactly(t, validMoves(game), "G6H7", "G6H6", "G6F6", "G6F7", "G6G7", // king moves,
		"H1G1", "H1F1", "H1E1", "H1D1", "H1C1", "H1B1", "H1A1", // rook moves
		"H2H3", // Pawns
	)
	game.Move(ParseMove("H2H3"))
	containsMovesExactly(t, validMoves(game)) // no valid moves
}

// ---------------------------------------------------------------------------------------------------------------------------------------

func givenGame() Game {
	return NewGame(
		speed.NewLeanPool(16, NewBoardB),
		speed.NewLeanPool(16, NewMovesIterator),
		speed.NewLeanPool(32, NewSquaresIterator))
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
		if !move.IsValid() {
			t.Fatalf("the move %v provided in the test is not valid, check it\n", expect)
		}
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

func printValidMoves(game Game) {
	fmt.Println(keys(validMoves(game)))
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
	for k := range moves {
		ans = append(ans, k.String())
	}
	return ans
}
