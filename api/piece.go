package api

func (p Piece) String() string {
	switch p {
	case Black + Pawn:
		return "p"
	case White + Pawn:
		return "P"
	case Black + Rook:
		return "r"
	case White + Rook:
		return "R"
	case Black + Knight:
		return "n"
	case White + Knight:
		return "N"
	case Black + Bishop:
		return "b"
	case White + Bishop:
		return "B"
	case Black + Queen:
		return "q"
	case White + Queen:
		return "Q"
	case Black + King:
		return "k"
	case White + King:
		return "K"
	default:
		return "?"
	}
}

func (p Piece) Colour() Piece {
	if p >= 8 {
		return White
	} else {
		return Black
	}
}

func (p Piece) Opponent() Piece {
	if p.Colour() == White {
		return Black
	} else {
		return White
	}
}

func (p Piece) IsPawn() bool {
	return p&7 == Pawn
}

func (p Piece) IsRook() bool {
	return p&7 == Rook
}

func (p Piece) IsKnight() bool {
	return p&7 == Knight
}

func (p Piece) IsBishop() bool {
	return p&7 == Bishop
}

func (p Piece) IsQueen() bool {
	return p&7 == Queen
}

func (p Piece) IsKing() bool {
	return p&7 == King
}

func (p Piece) IsEmpty() bool {
	return p == Empty
}

func (p Piece) Pawn() Piece {
	return p.Colour() + Pawn
}

func (p Piece) Rook() Piece {
	return p.Colour() + Rook
}

func (p Piece) Knight() Piece {
	return p.Colour() + Knight
}

func (p Piece) Bishop() Piece {
	return p.Colour() + Bishop
}

func (p Piece) Queen() Piece {
	return p.Colour() + Queen
}

func (p Piece) King() Piece {
	return p.Colour() + King
}

// CanAttack receives a board, and a Move and checks if the piece in the Move.From can attack Move.To i.e. if a Rook has no pieces in between, it does not
// check the Move.To is empty or full, just answers if the Piece located at the Board's Move.From can in this board reach and attack Move.To
func (p Piece) CanAttack(b Board, move Move) bool {
	if p.IsPawn() {
		return p.canPawnAttack(b, move)
	} else if p.IsRook() {
		return p.canRookAttack(b, move)
	} else if p.IsQueen() {
		return p.canQueenAttack(b, move)
	} else if p.IsKing() {
		return p.canKingAttack(b, move)
	}
	return false
}

func (p Piece) canPawnAttack(_ Board, move Move) bool {
	if p.Colour() == Black {
		return move.From().S().W() == move.To() || move.From().S().E() == move.To()
	} else {
		return move.From().N().W() == move.To() || move.From().N().E() == move.To()
	}
}

func (p Piece) canRookAttack(b Board, move Move) bool {
	if move.From().Row() != move.To().Row() && move.From().Col() != move.To().Col() {
		return false
	}
	j := move.From().Min(move.To())
	k := move.From().Max(move.To())
	if move.From().Row() == move.To().Row() {
		for i := j + 1; i < k; i++ {
			if b.Get(i) != Empty {
				return false
			}
		}
	} else {
		for i := j + OneRow; i < k; i += OneRow {
			if b.Get(i) != Empty {
				return false
			}
		}
	}
	return true
}

func (p Piece) canQueenAttack(b Board, move Move) bool {
	// quick check, not same colour? out
	if b.Get(move.From()).Colour() != b.Get(move.To()).Opponent() {
		return false
	}
	// horizontal - vertical, like Rooks
	if move.From().Row() == move.To().Row() || move.From().Col() == move.To().Col() {
		j := move.From().Min(move.To())
		k := move.From().Max(move.To())
		if move.From().Row() == move.To().Row() {
			for i := j + 1; i < k; i++ {
				if b.Get(i) != Empty {
					return false
				}
			}
		} else {
			for i := j + OneRow; i < k; i += OneRow {
				if b.Get(i) != Empty {
					return false
				}
			}
		}
		return true
	}
	// diagonals, like bishops
	// fast check if abs(delta rows) != abs(delta cols) => no
	if Abs8(int8(move.From().Col())-int8(move.To().Col())) != Abs8((int8(move.From().Row())-int8(move.To().Row()))/8) {
		return false
	}
	dRow := int8(OneRow)
	if move.From().Row() > move.To().Row() {
		dRow = -dRow
	}
	dCol := int8(1)
	if move.From().Col() > move.To().Col() {
		dCol = -dCol
	}
	delta := dRow + dCol
	for square := Square(int8(move.From()) + delta); square != move.To(); square = Square(int8(square) + delta) {
		if b.Get(square) != Empty {
			return false
		}
	}

	return true
}

func (p Piece) canKingAttack(b Board, move Move) bool {
	if b.Get(move.From()).Colour() != b.Get(move.To()).Opponent() {
		return false
	}
	return Abs8(int8(move.From().Row())-int8(move.To().Row())) <= 8 && Abs8(int8(move.From().Col())-int8(move.To().Col())) <= 1
}

func ParsePiece(ch rune) Piece {
	switch ch {
	case 'p':
		return Black + Pawn
	case 'P':
		return White + Pawn
	case 'r':
		return Black + Rook
	case 'R':
		return White + Rook
	case 'n':
		return Black + Knight
	case 'N':
		return White + Knight
	case 'b':
		return Black + Bishop
	case 'B':
		return White + Bishop
	case 'q':
		return Black + Queen
	case 'Q':
		return White + Queen
	case 'k':
		return Black + King
	case 'K':
		return White + King
	default:
		return Empty
	}
}
