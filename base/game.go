package base

import "github.com/ComputerSaysYeah/RookMills/speed"
import . "github.com/ComputerSaysYeah/RookMills/api"

type gameSt struct {
	speed.Recyclable

	board          Board
	moveNo         int
	halfMoveNo     int
	nextPlayer     Piece
	enPassant      Square
	WK, WQ, bk, bq bool // castling

	boardPool     speed.Pool[Board]
	movesIterPool speed.Pool[MovesIterator]

	returner func(any)
}

func NewGame(boardPool speed.Pool[Board], movesIterPool speed.Pool[MovesIterator]) Game {
	board := boardPool.Lease()
	board.SetStartingPieces()
	return &gameSt{
		board:         board,
		moveNo:        1,
		halfMoveNo:    0,
		nextPlayer:    White,
		enPassant:     None,
		WK:            true,
		WQ:            true,
		bk:            true,
		bq:            true,
		boardPool:     boardPool,
		movesIterPool: movesIterPool,
		returner:      nil}
}

func (g *gameSt) Reset() {
	g.board.SetStartingPieces()
	g.moveNo = 1
	g.halfMoveNo = 0
	g.WK, g.WQ, g.bk, g.bq = true, true, true, true
	g.nextPlayer = White
}

func (g *gameSt) MoveNo() int {
	return g.moveNo
}

func (g *gameSt) HalfMoveNo() int {
	return g.halfMoveNo
}

func (g *gameSt) MoveNext() Piece {
	return g.nextPlayer
}

func (g *gameSt) EnPassant() Square {
	return g.enPassant
}

func (g *gameSt) Move(move Move) {

}

func (g *gameSt) Castling() (WK, WQ, bk, bq bool) {
	return g.WK, g.WQ, g.bk, g.bq
}

func (g *gameSt) SetReturnerFn(returner func(any)) {
	g.returner = returner
}

func (g *gameSt) Return() {
	g.returner(g)
}

func (g *gameSt) Board() Board {
	return g.board
}

func (g *gameSt) SetMoveNo(moveNo int) {
	g.moveNo = moveNo
}

func (g *gameSt) SetHalfMoveNo(halfMoveNo int) {
	g.halfMoveNo = halfMoveNo
}

func (g *gameSt) SetMoveNext(piece Piece) {
	g.nextPlayer = piece
}

func (g *gameSt) SetEnPassant(square Square) {
	g.enPassant = square
}

func (g *gameSt) SetCastling(WK, WQ, bk, bq bool) {
	g.WK = WK
	g.WQ = WQ
	g.bk = bk
	g.bq = bq
}
