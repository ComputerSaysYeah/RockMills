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
	if piece.Colour() == Black {
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
	} else {
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
	attacks := []Square{square.E(), square.W()}
	for i := 0; i < len(attacks); i++ {
		if piece.Colour() == Black {
			attacks[i] = attacks[i].S()
		} else {
			attacks[i] = attacks[i].N()
		}
	}
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
		if square.Row() == Row8 && g.board.Get(move.To()).IsEmpty() {
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(EncodeMovePromote(move.From(), move.To(), Queen))
				iter.Add(EncodeMovePromote(move.From(), move.To(), Rook))
				iter.Add(EncodeMovePromote(move.From(), move.To(), Bishop))
				iter.Add(EncodeMovePromote(move.From(), move.To(), Knight))

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
	g.validGenericMoveAndCapture(square,
		[]Square{
			square.N().N().W(), square.N().N().E(),
			square.E().E().N(), square.E().E().S(),
			square.S().S().E(), square.S().S().W(),
			square.W().W().S(), square.W().W().N(),
		},
		piece, iter)
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
	g.validGenericMoveAndCapture(square,
		[]Square{square.N(), square.N().E(), square.E(), square.E().S(), square.S(), square.S().W(), square.W(), square.W().N()},
		piece, iter)
}

func (g *gameSt) validGenericMoveAndCapture(square Square, targets []Square, piece Piece, iter MovesIterator) {
	for _, target := range targets {
		if target != None && (g.board.Get(target) == Empty || g.board.Get(target).Colour() == piece.Opponent()) {
			move := EncodeMove(square, target)
			if !g.wouldCheckKing(piece.Colour(), move) {
				iter.Add(move)
			}
		}
	}
}

func (g *gameSt) validGenericCapture(square Square, targets []Square, piece Piece, iter MovesIterator) {
	for _, target := range targets {
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
	kingPiece := kingColour + King
	kingSquare := None
	for s := A1; s < H8 && kingSquare == None; s++ {
		if b.Get(s) == kingPiece {
			kingSquare = s
		}
	}

	// know, we check every opponents' piece for potential attack
	for s := A1; s < H8; s++ {
		thisSquare := b.Get(s)
		if !thisSquare.IsEmpty() && thisSquare.Colour() != kingColour {
			if thisSquare.IsAttack(b, EncodeMove(s, kingSquare)) {
				return true
			}
		}
	}

	return false
}
