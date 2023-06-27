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

func TestMove(t *testing.T) {
	for fromRow := Row8; fromRow <= Row8; fromRow -= 8 {
		for fromCol := ColA; fromCol <= ColH; fromCol++ {
			for toRow := Row8; toRow <= Row8; toRow -= 8 {
				for toCol := ColA; toCol <= ColH; toCol++ {
					for piece := Empty; piece <= King; piece++ {
						move := EncodeMovePromote(fromRow+fromCol, toRow+toCol, piece)
						if move.From() != fromRow+fromCol {
							t.Fatal()
						}
						if move.To() != toRow+toCol {
							t.Fatal()
						}
						if move.Promote() != piece {
							t.Fatal()
						}
						s := move.String()
						//log.Println((fromRow + fromCol).String(), (toRow + toCol).String(), piece.String(), s)
						if s[0:2] != (fromRow + fromCol).String() {
							t.Fatal()
						}
						if s[2:4] != (toRow + toCol).String() {
							t.Fatal()
						}
						if piece != Empty && s[4:5] != piece.String() {
							t.Fatal()
						}
						if move != ParseMove(move.String()) { //ParseMove
							t.Fatal()
						}
					}
				}
			}
		}
	}
}
