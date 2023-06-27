package base

import (
	. "github.com/ComputerSaysYeah/RookMills/api"
	"log"
	"testing"
)

func TestParseFENStr(t *testing.T) {
	g := givenGame()
	if err := g.FromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"); err != nil {
		log.Fatal(err)
	}

	b := NewBoardB()
	b.SetStartingPieces()
	if g.Board().Hash() != b.Hash() {
		log.Fatalf("it should have been equal to a brand new laid out board board, got: \n%v\n", g.Board().String())
	}

	if g.MoveNo() != 1 {
		log.Fatal("move should be 0")
	}
	if g.HalfMoveNo() != 0 {
		log.Fatal("half-move should be 0")
	}
	if g.MoveNext() != White {
		log.Fatal("next mover should be white")
	}
	if g.EnPassant() != None {
		log.Fatal("en passant should be none")
	}
	if WK, WQ, bk, bq := g.Castling(); !WK || !WQ || !bk || !bq {
		t.Fatal("castling not ok")
	}
}

func TestParseFENEnPassant(t *testing.T) {
	g := givenGame()
	if err := g.FromFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"); err != nil {
		log.Fatal(err)
	}

	if g.Board().Get(Row4+ColE) != White+Pawn {
		log.Fatal("should be a Pawn there")
	}
	if g.EnPassant() != Row3+ColE {
		log.Fatal("en Passant should be loaded")
	}
	if WK, WQ, bk, bq := g.Castling(); !WK || !WQ || !bk || !bq {
		t.Fatal("castling not ok")
	}
	if g.MoveNext() != Black {
		log.Fatal("black should be moving next")
	}
	if g.HalfMoveNo() != 0 {
		log.Fatal("half move is wrong")
	}
	if g.MoveNo() != 1 {
		log.Fatal("it should be move 1")
	}
}

func TestParseFENAdvanced(t *testing.T) {
	g := givenGame()
	if err := g.FromFEN("8/5k2/3p4/1p1Pp2p/pP2Pp1P/P4P1K/8/8 b - - 99 50"); err != nil {
		log.Fatal(err)
	}

	if g.Board().Get(Row3+ColH) != White+King {
		log.Fatal("should be a White King there")
	}
	if g.Board().Get(Row7+ColF) != Black+King {
		log.Fatal("should be a Black King there")
	}
	if g.EnPassant() != None {
		log.Fatal("should not be an en Passant")
	}
	if WK, WQ, bk, bq := g.Castling(); WK || WQ || bk || bq {
		t.Fatal("castling not ok")
	}
	if g.MoveNext() != Black {
		log.Fatal("black should be moving next")
	}
	if g.HalfMoveNo() != 99 {
		log.Fatal("half move is wrong")
	}
	if g.MoveNo() != 50 {
		log.Fatal("it should be move 50")
	}

}

func TestFENRoundTrip(t *testing.T) {
	g := givenGame()
	if err := g.FromFEN("r7/4r3/1qbk1n2/3p4/4P3/2NK1BQ1/3R4/7R b - - 0 1"); err != nil {
		log.Fatal(err)
	}
	if g.ToFEN() != "r7/4r3/1qbk1n2/3p4/4P3/2NK1BQ1/3R4/7R b - - 0 1" {
		t.Fatalf("we didn't get quite exactly \"r7/4r3/1qbk1n2/3p4/4P3/2NK1BQ1/3R4/7R b - - 0 1\", but: \"%v\"", g.ToFEN())
	}
}
