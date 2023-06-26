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

func (p Piece) IsValidAttack(b Board, move Move) bool {
	if p.IsPawn() {
		if p.Colour() == Black {
			to := move.From().S().W()
			if to == move.To() && b.Get(to) == White {
				return true
			}
			to = move.From().S().E()
			if to == move.To() && b.Get(to) == White {
				return true
			}
		} else {
			to := move.From().N().W()
			if to == move.To() && b.Get(to) == Black {
				return true
			}
			to = move.From().N().E()
			if to == move.To() && b.Get(to) == Black {
				return true
			}
		}
	} else if p.IsRook() {
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
	return false
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
