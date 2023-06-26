package base

import (
	. "github.com/ComputerSaysYeah/RookMills/api"
)

func (g *gameSt) ValidMoves() MovesIterator {
	moves := g.movesIterPool.Lease()
	for square := Square(0); square < 64; square++ {
		piece := g.board.Get(square)
		if piece != Empty && piece.Colour() == g.nextPlayer {
			g.validMovesPieceIn(square, piece, moves)
		}
	}
	return moves
}

func (g *gameSt) validMovesPieceIn(square Square, piece Piece, movesIter MovesIterator) {
	if piece.IsPawn() {
		//XXX: DEAL with EnPassant
		g.validPawnMoves(square, piece, movesIter)
		g.validPawnCaptures(square, piece, movesIter)
		g.validPawnPromotes(square, piece, movesIter)
	} else if piece.IsRook() {
		g.validRookMoves(square, piece, movesIter)
		g.validRootCaptures(square, piece, movesIter)
	} else if piece.IsKnight() {
		g.validKnightMoves(square, piece, movesIter)
		g.validKnightCaptures(square, piece, movesIter)
	} else if piece.IsBishop() {
		g.validBishopMoves(square, piece, movesIter)
		g.validBishopCaptures(square, piece, movesIter)
	} else if piece.IsQueen() {
		g.validQueenMoves(square, piece, movesIter)
		g.validQueenCaptures(square, piece, movesIter)
	} else if piece.IsKing() {
		g.validKingMoves(square, piece, movesIter)
		g.validKingCaptures(square, piece, movesIter)
	}
}

func (g *gameSt) validPawnMoves(square Square, piece Piece, iter MovesIterator) {
	if piece.Colour() == Black {
		move := EncodeMove(square, square.S(), Empty)
		if g.board.Get(move.To()) == Empty {
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
			move = EncodeMove(square, move.To().S(), Empty)
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
		}
	} else {
		move := EncodeMove(square, square.N(), Empty)
		if g.board.Get(move.To()) == Empty {
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
			move = EncodeMove(square, move.To().N(), Empty)
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
		}
	}
}

func (g *gameSt) validPawnCaptures(square Square, piece Piece, iter MovesIterator) {
	if piece.Colour() == Black {
	}
}
func (g *gameSt) validPawnPromotes(square Square, piece Piece, iter MovesIterator) {
	if piece.Colour() == Black {
		move := EncodeMove(square, square.S(), Empty)
		if square.Row() == Row2 && g.board.Get(move.To()).IsEmpty() {
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(EncodeMove(move.From(), move.To(), Queen))
				iter.Add(EncodeMove(move.From(), move.To(), Rook))
				iter.Add(EncodeMove(move.From(), move.To(), Bishop))
				iter.Add(EncodeMove(move.From(), move.To(), Knight))
			}
		}
	} else {
		move := EncodeMove(square, square.N(), Empty)
		if square.Row() == Row8 && g.board.Get(move.To()).IsEmpty() {
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(EncodeMove(move.From(), move.To(), Queen))
				iter.Add(EncodeMove(move.From(), move.To(), Rook))
				iter.Add(EncodeMove(move.From(), move.To(), Bishop))
				iter.Add(EncodeMove(move.From(), move.To(), Knight))

			}
		}
	}
}

func (g *gameSt) validRookMoves(square Square, piece Piece, iter MovesIterator) {

}

func (g *gameSt) validRootCaptures(square Square, piece Piece, iter MovesIterator) {

}
func (g *gameSt) validKnightMoves(square Square, piece Piece, iter MovesIterator) {
	targets := []Square{
		square.N().N().W(), square.N().N().E(),
		square.E().E().N(), square.E().E().S(),
		square.S().S().E(), square.S().S().W(),
		square.W().W().S(), square.W().W().N(),
	}
	for _, target := range targets {
		if target != None && g.board.Get(target) == Empty {
			move := EncodeMove(square, target, Empty)
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
		}
	}
}

func (g *gameSt) validKnightCaptures(square Square, piece Piece, iter MovesIterator) {

}

func (g *gameSt) validBishopMoves(square Square, piece Piece, iter MovesIterator) {

}

func (g *gameSt) validBishopCaptures(square Square, piece Piece, iter MovesIterator) {

}

func (g *gameSt) validQueenMoves(square Square, piece Piece, iter MovesIterator) {

}
func (g *gameSt) validQueenCaptures(square Square, piece Piece, iter MovesIterator) {

}

func (g *gameSt) validKingMoves(square Square, piece Piece, iter MovesIterator) {

}
func (g *gameSt) validKingCaptures(square Square, piece Piece, iter MovesIterator) {

}

func (g *gameSt) wouldCheckKing(kingColour Piece, move Move) bool {

	b := g.boardPool.Lease()
	defer b.Return()
	//XXX: castling should be done here
	b.CopyFrom(g.board)
	b.Set(move.To(), b.Get(move.From()))
	b.Set(move.From(), Empty)

	// find the King
	kingSquare := Square(None)
	for s := A1; s < H8 && kingSquare == None; s++ {
		piece := b.Get(s)
		if piece.IsKing() && piece.Colour() == kingColour {
			kingSquare = s
		}
	}

	// know, we check every opponents' piece for potential attack
	for s := A1; s < H8; s++ {
		thisSquare := b.Get(s)
		if !thisSquare.IsEmpty() && thisSquare.Opponent() == kingColour {
			if thisSquare.IsValidAttack(b, EncodeMove(s, kingSquare, Empty)) {
				return true
			}
		}
	}

	return false
}
