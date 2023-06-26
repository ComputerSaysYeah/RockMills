package api

import (
	"fmt"
	"testing"
)

func TestParsePiece(t *testing.T) {
	for _, ch := range "pPrRnNbBqQkK" {
		piece := ParsePiece(ch)
		if piece.String() != fmt.Sprintf("%c", ch) {
			t.Fatalf("'%c' did not round-trip into '%v' value %d\n", ch, piece.String(), piece)
		}
	}
}

func TestIsPiece(t *testing.T) {
	if !(White + Pawn).IsPawn() || !Pawn.IsPawn() {
		t.Fatal()
	}
}
