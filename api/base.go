package api

import (
	"github.com/ComputerSaysYeah/RookMills/speed"
)

type Piece uint8
type Square uint8
type Move uint16

const (
	Empty  Piece = 0
	Pawn   Piece = 1
	Knight Piece = 2
	Bishop Piece = 3
	Rook   Piece = 4
	Queen  Piece = 6
	King   Piece = 7

	Black Piece = 0 // i.e. Black Queen = 0 + 6 = 6  (0000_0110)
	White Piece = 8 // i.e. White Queen = 8 + 6 = 14 (0000_1110)

	ColA Square = 0
	ColB Square = 1
	ColC Square = 2
	ColD Square = 3
	ColE Square = 4
	ColF Square = 5
	ColG Square = 6
	ColH Square = 7

	Row1 Square = 0
	Row2 Square = 8
	Row3 Square = 2 * 8
	Row4 Square = 3 * 8
	Row5 Square = 4 * 8
	Row6 Square = 5 * 8
	Row7 Square = 6 * 8
	Row8 Square = 7 * 8

	Invalid Square = 255
	// i.e. Row6 + ColD = 6*8+3 = 51
)

type Board interface {
	speed.Recyclable
	Get(Square) Piece
	Set(Square, Piece)
	CopyFrom(*Board)
	Hash() uint64
	SetStartingPieces()
	String() string
}

type Game interface {
	speed.Recyclable
	MoveNo() int
	MoveNext() Piece
	ValidMoves() speed.Iterator[Move]
	Move(Move)
}
