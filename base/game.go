package base

import "github.com/ComputerSaysYeah/RookMills/speed"
import . "github.com/ComputerSaysYeah/RookMills/api"

type gameSt struct {
	board      Board
	moveNo     int
	nextPlayer Piece

	boardPool speed.Pool[Board]
	returner  func(any)
}

func NewGame(boardPool speed.Pool[Board]) Game {
	board := boardPool.Lease()
	board.SetStartingPieces()
	return &gameSt{
		board:      board,
		moveNo:     1,
		nextPlayer: White,
		boardPool:  boardPool,
		returner:   nil}
}

func (g *gameSt) Reset() {
	g.board.SetStartingPieces()
	g.moveNo = 1
	g.nextPlayer = White
}

func (g *gameSt) MoveNo() int {
	return g.moveNo
}

func (g *gameSt) MoveNext() Piece {
	return g.nextPlayer
}

func (g *gameSt) ValidMoves() speed.Iterator[Move] {
	return nil
}

func (g *gameSt) Move(move Move) {

}

func (g *gameSt) SetReturnerFn(returner func(any)) {
	g.returner = returner
}

func (g *gameSt) Return() {
	g.returner(g)
}
