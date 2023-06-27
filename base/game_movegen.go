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
		g.validRookCaptures(square, piece, movesIter)
	} else if piece.IsKnight() {
		g.validKnightMovesAndCaptures(square, piece, movesIter)
	} else if piece.IsBishop() {
		g.validBishopMoves(square, piece, movesIter)
		g.validBishopCaptures(square, piece, movesIter)
	} else if piece.IsQueen() {
		g.validQueenMoves(square, piece, movesIter)
		g.validQueenCaptures(square, piece, movesIter)
	} else if piece.IsKing() {
		g.validKingMovesAndCaptures(square, piece, movesIter)
	}
}

func (g *gameSt) validPawnMoves(square Square, piece Piece, iter MovesIterator) {
	if piece.Colour() == Black && square.Row() != Row2 { // so no promotion
		move := EncodeMove(square, square.S())
		if g.board.Get(move.To()) == Empty {
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
			if square.Row() == Row7 {
				move = EncodeMove(square, move.To().S())
				if !g.wouldCheckKing(piece.Colour(), move) {
					iter.Add(move)
				}
			}
		}
	} else if piece.Colour() == White && square.Row() != Row7 { // so no promotion
		move := EncodeMove(square, square.N())
		if g.board.Get(move.To()) == Empty {
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
			if square.Row() == Row2 {
				move = EncodeMove(square, move.To().N())
				if !g.wouldCheckKing(piece.Colour(), move) {
					iter.Add(move)
				}
			}
		}
	}
}

func (g *gameSt) validPawnCaptures(square Square, piece Piece, iter MovesIterator) {
	//XXX: en-passant here
	aEast, aWest := square.E(), square.W()
	if piece.Colour() == Black {
		aEast, aWest = aEast.S(), aWest.S()
	} else {
		aEast, aWest = aEast.N(), aWest.N()
	}
	attacks := g.squaresIterPool.Lease()
	attacks.Add(aWest).Add(aEast)
	g.validGenericCapture(square, attacks, piece, iter)
}

func (g *gameSt) validPawnPromotes(square Square, piece Piece, iter MovesIterator) {
	if piece.Colour() == Black {
		move := EncodeMove(square, square.S())
		if square.Row() == Row2 && g.board.Get(move.To()).IsEmpty() {
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(EncodeMovePromote(move.From(), move.To(), Queen))
				iter.Add(EncodeMovePromote(move.From(), move.To(), Rook))
				iter.Add(EncodeMovePromote(move.From(), move.To(), Bishop))
				iter.Add(EncodeMovePromote(move.From(), move.To(), Knight))
			}
		}
	} else {
		move := EncodeMove(square, square.N())
		if square.Row() == Row7 && g.board.Get(move.To()).IsEmpty() {
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(EncodeMovePromote(move.From(), move.To(), White+Queen))
				iter.Add(EncodeMovePromote(move.From(), move.To(), White+Rook))
				iter.Add(EncodeMovePromote(move.From(), move.To(), White+Bishop))
				iter.Add(EncodeMovePromote(move.From(), move.To(), White+Knight))

			}
		}
	}
}

func (g *gameSt) validRookMoves(square Square, piece Piece, iter MovesIterator) {
	b := g.Board()
	var move Move
	for s := square.N(); s != None && b.Get(s).IsEmpty(); s = s.N() {
		move = EncodeMove(square, s)
		if !g.wouldCheckKing(piece.Colour(), move) {
			iter.Add(move)
		}
	}
	for s := square.E(); s != None && b.Get(s).IsEmpty(); s = s.E() {
		move = EncodeMove(square, s)
		if !g.wouldCheckKing(piece.Colour(), move) {
			iter.Add(move)
		}
	}
	for s := square.S(); s != None && b.Get(s).IsEmpty(); s = s.S() {
		move = EncodeMove(square, s)
		if !g.wouldCheckKing(piece.Colour(), move) {
			iter.Add(move)
		}
	}
	for s := square.W(); s != None && b.Get(s).IsEmpty(); s = s.W() {
		move = EncodeMove(square, s)
		if !g.wouldCheckKing(piece.Colour(), move) {
			iter.Add(move)
		}
	}
}

func (g *gameSt) validRookCaptures(square Square, piece Piece, iter MovesIterator) {

}

func (g *gameSt) validKnightMovesAndCaptures(square Square, piece Piece, iter MovesIterator) {
	targets := g.squaresIterPool.Lease()
	defer targets.Return()
	targets.
		Add(square.N().N().W()).Add(square.N().N().E()).
		Add(square.E().E().N()).Add(square.E().E().S()).
		Add(square.S().S().E()).Add(square.S().S().W()).
		Add(square.W().W().S()).Add(square.W().W().N())
	g.validGenericMoveAndCapture(square, targets, piece, iter)
}

func (g *gameSt) validBishopMoves(square Square, piece Piece, iter MovesIterator) {

}

func (g *gameSt) validBishopCaptures(square Square, piece Piece, iter MovesIterator) {

}

func (g *gameSt) validQueenMoves(square Square, piece Piece, iter MovesIterator) {

}
func (g *gameSt) validQueenCaptures(square Square, piece Piece, iter MovesIterator) {

}

func (g *gameSt) validKingMovesAndCaptures(square Square, piece Piece, iter MovesIterator) {
	targets := g.squaresIterPool.Lease()
	defer targets.Return()
	targets.
		Add(square.N()).
		Add(square.N().E()).
		Add(square.E()).
		Add(square.E().S()).
		Add(square.S()).
		Add(square.S().W()).
		Add(square.W()).
		Add(square.W().N())
	g.validGenericMoveAndCapture(square, targets, piece, iter)
}

func (g *gameSt) validGenericMoveAndCapture(square Square, targets SquaresIterator, piece Piece, iter MovesIterator) {
	for targets.HasNext() {
		target := targets.Next()
		if target != None && (g.board.Get(target) == Empty || g.board.Get(target).Colour() == piece.Opponent()) {
			move := EncodeMove(square, target)
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
		}
	}
}

func (g *gameSt) validGenericCapture(square Square, targets SquaresIterator, piece Piece, iter MovesIterator) {
	for targets.HasNext() {
		target := targets.Next()
		if !target.IsNone() && !g.board.Get(target).IsEmpty() && g.board.Get(target).Colour() == piece.Opponent() {
			move := EncodeMove(square, target)
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
		}
	}
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
	kingPiece := kingColour.King()
	kingSquare := None
	for s := A1; s < H8 && kingSquare == None; s++ {
		if b.Get(s) == kingPiece {
			kingSquare = s
		}
	}

	// know, we check every opponents' piece for potential attack
	for s := A1; s < H8; s++ {
		piece := b.Get(s)
		if !piece.IsEmpty() && piece.Colour() != kingColour {
			if piece.CanAttack(b, EncodeMove(s, kingSquare)) {
				return true
			}
		}
	}

	return false
}
