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
		g.validRookMovesAndCapture(square, piece, movesIter)
	} else if piece.IsKnight() {
		g.validKnightMovesAndCaptures(square, piece, movesIter)
	} else if piece.IsBishop() {
		g.validBishopMovesAndCaptures(square, piece, movesIter)
	} else if piece.IsQueen() {
		g.validRookMovesAndCapture(square, piece, movesIter)
		g.validBishopMovesAndCaptures(square, piece, movesIter)
	} else if piece.IsKing() {
		g.validKingMovesAndCaptures(square, piece, movesIter)
	}
}

func (g *gameSt) validPawnMoves(square Square, piece Piece, iter MovesIterator) {
	if piece.Colour() == Black {
		if square.Row() != Row2 { // so no promotion
			move := EncodeMove(square, square.S())
			if g.board.Get(move.To()).IsEmpty() {
				if !g.wouldCheckKing(piece.Colour(), move) {
					iter.Add(move)
				}
				if square.Row() == Row7 {
					move = EncodeMove(square, move.To().S())
					if g.board.Get(move.To()).IsEmpty() && !g.wouldCheckKing(piece.Colour(), move) {
						iter.Add(move)
					}
				}
			}
		}
		// white
	} else if square.Row() != Row7 { // so no promotion
		move := EncodeMove(square, square.N())
		if g.board.Get(move.To()).IsEmpty() {
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
			if square.Row() == Row2 {
				move = EncodeMove(square, move.To().N())
				if g.board.Get(move.To()).IsEmpty() && !g.wouldCheckKing(piece.Colour(), move) {
					iter.Add(move)
				}
			}
		}
	}
}

func (g *gameSt) validPawnCaptures(square Square, piece Piece, iter MovesIterator) {
	//XXX: en-passant here
	if piece.Colour() == Black {
		g.validCapture(square, square.SE(), piece, iter)
		g.validCapture(square, square.SW(), piece, iter)
	} else {
		g.validCapture(square, square.NE(), piece, iter)
		g.validCapture(square, square.NW(), piece, iter)
	}
}

func (g *gameSt) validPawnPromotes(square Square, piece Piece, iter MovesIterator) {
	if piece.Colour() == Black {
		if square.Row() == Row2 {
			move := EncodeMove(square, square.S())
			if g.board.Get(move.To()).IsEmpty() && !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(EncodeMovePromote(move.From(), move.To(), Queen))
				iter.Add(EncodeMovePromote(move.From(), move.To(), Rook))
				iter.Add(EncodeMovePromote(move.From(), move.To(), Bishop))
				iter.Add(EncodeMovePromote(move.From(), move.To(), Knight))
			}
		}
	} else {
		if square.Row() == Row7 {
			move := EncodeMove(square, square.N())
			if g.board.Get(move.To()).IsEmpty() && !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(EncodeMovePromote(move.From(), move.To(), White+Queen))
				iter.Add(EncodeMovePromote(move.From(), move.To(), White+Rook))
				iter.Add(EncodeMovePromote(move.From(), move.To(), White+Bishop))
				iter.Add(EncodeMovePromote(move.From(), move.To(), White+Knight))
			}
		}
	}
}

func (g *gameSt) validRookMovesAndCapture(square Square, piece Piece, iter MovesIterator) {
	b := g.Board()
	var s Square
	for s = square.N(); s.IsValid() && b.Get(s).IsEmpty(); s = s.N() {
		g.validMoveCapture(square, s, piece, iter)
	}
	if s.IsValid() && b.Get(s).Colour() == piece.Opponent() {
		g.validMoveCapture(square, s, piece, iter)
	}
	for s = square.E(); s.IsValid() && b.Get(s).IsEmpty(); s = s.E() {
		g.validMoveCapture(square, s, piece, iter)
	}
	if s.IsValid() && b.Get(s).Colour() == piece.Opponent() {
		g.validMoveCapture(square, s, piece, iter)
	}
	for s = square.S(); s.IsValid() && b.Get(s).IsEmpty(); s = s.S() {
		g.validMoveCapture(square, s, piece, iter)
	}
	if s.IsValid() && b.Get(s).Colour() == piece.Opponent() {
		g.validMoveCapture(square, s, piece, iter)
	}
	for s = square.W(); s.IsValid() && b.Get(s).IsEmpty(); s = s.W() {
		g.validMoveCapture(square, s, piece, iter)
	}
	if s.IsValid() && b.Get(s).Colour() == piece.Opponent() {
		g.validMoveCapture(square, s, piece, iter)
	}
}

func (g *gameSt) validKnightMovesAndCaptures(square Square, piece Piece, iter MovesIterator) {
	g.validMoveCapture(square, square.NNW(), piece, iter)
	g.validMoveCapture(square, square.NNE(), piece, iter)
	g.validMoveCapture(square, square.EEN(), piece, iter)
	g.validMoveCapture(square, square.EES(), piece, iter)
	g.validMoveCapture(square, square.SSW(), piece, iter)
	g.validMoveCapture(square, square.SSE(), piece, iter)
	g.validMoveCapture(square, square.WWN(), piece, iter)
	g.validMoveCapture(square, square.WWS(), piece, iter)
}

func (g *gameSt) validBishopMovesAndCaptures(square Square, piece Piece, iter MovesIterator) {
	b := g.Board()
	var s Square
	for s = square.NE(); s.IsValid() && b.Get(s).IsEmpty(); s = s.NE() {
		g.validMoveCapture(square, s, piece, iter)
	}
	if s.IsValid() && b.Get(s).Colour() == piece.Opponent() {
		g.validMoveCapture(square, s, piece, iter)
	}
	for s = square.SE(); s.IsValid() && b.Get(s).IsEmpty(); s = s.SE() {
		g.validMoveCapture(square, s, piece, iter)
	}
	if s.IsValid() && b.Get(s).Colour() == piece.Opponent() {
		g.validMoveCapture(square, s, piece, iter)
	}
	for s = square.SW(); s.IsValid() && b.Get(s).IsEmpty(); s = s.SW() {
		g.validMoveCapture(square, s, piece, iter)
	}
	if s.IsValid() && b.Get(s).Colour() == piece.Opponent() {
		g.validMoveCapture(square, s, piece, iter)
	}
	for s = square.NW(); s.IsValid() && b.Get(s).IsEmpty(); s = s.NW() {
		g.validMoveCapture(square, s, piece, iter)
	}
	if s.IsValid() && b.Get(s).Colour() == piece.Opponent() {
		g.validMoveCapture(square, s, piece, iter)
	}
}

func (g *gameSt) validKingMovesAndCaptures(square Square, piece Piece, iter MovesIterator) {
	g.validMoveCapture(square, square.N(), piece, iter)
	g.validMoveCapture(square, square.NE(), piece, iter)
	g.validMoveCapture(square, square.E(), piece, iter)
	g.validMoveCapture(square, square.SE(), piece, iter)
	g.validMoveCapture(square, square.S(), piece, iter)
	g.validMoveCapture(square, square.SW(), piece, iter)
	g.validMoveCapture(square, square.W(), piece, iter)
	g.validMoveCapture(square, square.NW(), piece, iter)
}

func (g *gameSt) validMoveCapture(from Square, target Square, piece Piece, iter MovesIterator) {
	if target.IsValid() && (g.board.Get(target) == Empty || g.board.Get(target).Colour() == piece.Opponent()) {
		move := EncodeMove(from, target)
		if !g.wouldCheckKing(piece.Colour(), move) {
			iter.Add(move)
		}
	}
}

func (g *gameSt) validCapture(square Square, target Square, piece Piece, iter MovesIterator) {
	if !target.IsNone() && !g.board.Get(target).IsEmpty() && g.board.Get(target).Colour() == piece.Opponent() {
		move := EncodeMove(square, target)
		if !g.wouldCheckKing(piece.Colour(), move) {
			iter.Add(move)
		}
	}
}

func (g *gameSt) wouldCheckKing(kingColour Piece, move Move) bool {

	b := g.board

	origFromPiece := b.Get(move.From())
	origToPiece := b.Get(move.To())
	defer func() {
		b.Set(move.From(), origFromPiece)
		b.Set(move.To(), origToPiece)
	}()

	if move.Promote().IsEmpty() {
		b.Set(move.To(), b.Get(move.From()))
	} else {
		b.Set(move.To(), move.Promote())
	}
	b.Set(move.From(), Empty)

	kingSquare := b.KingSquare(kingColour)
	if kingSquare.IsNone() {
		return false
	}

	// know, we check every opponents' piece for potential attack
	for s := A1; s <= H8; s++ {
		piece := b.Get(s)
		if !piece.IsEmpty() && piece.Colour() != kingColour {
			if piece.CanAttack(b, EncodeMove(s, kingSquare)) {
				return true
			}
		}
	}

	return false
}
