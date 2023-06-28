package api

import "testing"

func TestMove(t *testing.T) {
	for fromRow := Row8; fromRow <= Row8; fromRow -= 8 {
		for fromCol := ColA; fromCol <= ColH; fromCol++ {
			for toRow := Row8; toRow <= Row8; toRow -= 8 {
				for toCol := ColA; toCol <= ColH; toCol++ {
					for piece := Empty; piece <= King; piece++ {
						for color := Black; color <= White; color += White {
							pieceWithColour := piece + color
							if piece.IsEmpty() {
								pieceWithColour = piece
							}
							move := EncodeMovePromote(fromRow+fromCol, toRow+toCol, pieceWithColour)

							if move.From() != fromRow+fromCol {
								t.Fatal()
							}
							if move.To() != toRow+toCol {
								t.Fatal()
							}
							if !piece.IsEmpty() && move.Promote() != pieceWithColour {
								t.Fatalf("piece: '%v' pieceWithColour: '%v'", piece.String(), pieceWithColour.String())
							}
							if piece.IsEmpty() && !move.Promote().IsEmpty() {
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
							if !piece.IsEmpty() && s[4:5] != pieceWithColour.String() {
								t.Fatal()
							}
							if move != ParseMove(move.String()) { //ParseMove
								t.Fatalf("'%v' does not parses back correctly, we got '%v'", move.String(), ParseMove(move.String()).String())
							}
						}
					}

				}
			}
		}
	}
}
