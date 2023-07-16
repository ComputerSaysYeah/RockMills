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
	containsMovesExactly(t, validMoves(game)) // CheckMate - no valid moves
}

func TestMoveGen_Interesting2(t *testing.T) {
	// https://lichess.org/editor/5k2/8/5K2/8/8/8/1p6/5Q2_w_-_-_0_1?color=white
	// https://twitter.com/ChessMood/status/1680321341785014272
	// tests promotion
	game := givenGame()
	_ = game.FromFEN("5k2/8/5K2/8/8/8/1p6/5Q2 w - - 0 1")
	containsMovesExactly(t, validMoves(game) /*Q horizontal*/, "F1A1", "F1B1", "F1C1", "F1D1", "F1E1", "F1G1", "F1H1",
		/*Q vert*/ "F1F2", "F1F3", "F1F4", "F1F5" /* Q diagonal */, "F1G2", "F1H3", "F1E2", "F1D3", "F1C4", "F1B5", "F1A6",
		/*king*/ "F6G6", "F6G5", "F6F5", "F6E5", "F6E6",
	)
	game.Move(ParseMove("F1C4"))
	containsMovesExactly(t, validMoves(game), "F8E8", "B2B1q", "B2B1r", "B2B1n", "B2B1b")
	game.Move(ParseMove("B2B1b"))                                                                     // promotes to Bishop instead of Queen just for testing
	containsMovesExactly(t, validMoves(game), "C4A4", "C4B4", "C4D4", "C4E4", "C4F4", "C4G4", "C4H4", // horizontal queen moves
		"C4C1", "C4C2", "C4C3", "C4C5", "C4C6", "C4C7", "C4C8", // vertical queen moves
		"C4D3", "C4E2", "C4F1", "C4B5", "C4A6" /**/, "C4A2", "C4B3", "C4D5", "C4E6", "C4F7", "C4G8", // diagonal queen moves
		"F6G5", "F6E5", "F6E6", // king moves
	)
	game.Move(ParseMove("C4C8"))
	containsMovesExactly(t, validMoves(game)) // CheckMate - no valid moves
}

func TestMoveGen_Interesting3(t *testing.T) {
	// https://lichess.org/editor/8/5N1n/2n2k1P/8/6PN/1B2R3/8/r2Q2KR_w_-_-_0_1?color=white
	// https://twitter.com/ChessMood/status/1666429801954156545
	// lots of Knights
	game := givenGame()
	_ = game.FromFEN("8/5N1n/2n2k1P/8/6PN/1B2R3/8/r2Q2KR w - - 0 1")
	containsMovesExactly(t, validMoves(game), "D1A1", "D1B1", "D1C1", "D1E1", "D1F1", // queen horizontal, can't move from Row1
		"G1G2", "G1H2", "G1F1", "G1F2", // king
		"H1H2", "H1H3", // rock H1
		"E3E8", "E3E7", "E3E6", "E3E5", "E3E4", "E3E2", "E3E1", "E3C3", "E3D3", "E3F3", "E3G3", "E3H3", // Rock E3
		"G4G5",                                                                 // pawn
		"H4G2", "H4F3", "H4G6", "H4F5", "F7H8", "F7D8", "F7D6", "F7E5", "F7G5", // Knights
		"B3A2", "B3C2", "B3C4", "B3D5", "B3E6", "B3A4", // Bishop
	)
	game.Move(ParseMove("D1A1"))
	containsMovesExactly(t, validMoves(game), "C6E5", "C6D4")
	game.Move(ParseMove("C6D4"))
	containsMovesExactly(t, validMoves(game), "A1B1", "A1C1", "A1D1", "A1E1", "A1F1", // queen horizontal
		"A1A2", "A1A3", "A1A4", "A1A5", "A1A6", "A1A7", "A1A8", // queen vertical
		"A1B2", "A1C3", "A1D4", // queen diagonal
		"G1G2", "G1H2", "G1F1", "G1F2", // king
		"H1H2", "H1H3", // rock H1
		"E3E8", "E3E7", "E3E6", "E3E5", "E3E4", "E3E2", "E3E1", "E3C3", "E3D3", "E3F3", "E3G3", "E3H3", // Rock E3
		"G4G5",                                                                 // pawn
		"H4G2", "H4F3", "H4G6", "H4F5", "F7H8", "F7D8", "F7D6", "F7E5", "F7G5", // Knights
		"B3A2", "B3C2", "B3C4", "B3D5", "B3E6", "B3A4", "B3D1", // Bishop
	)
	// lots of valid moves but the important one is A1D4
	game.Move(ParseMove("A1D4"))
	containsMovesExactly(t, validMoves(game)) // CheckMate - no valid moves
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
