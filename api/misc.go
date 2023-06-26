package api

func (m Move) To() Square {
	return Square(m & 0x3F)
}

func (m Move) From() Square {
	return Square((m & 0xFC0) >> 6)
}

func (m Move) Promote() Piece {
	return Piece((m & 0x7000) >> 12)
}

func EncodeMove(from, to Square, promote Piece) Move {
	return Move(to) | Move(from)<<6 | Move(promote)<<12
}

func (m Move) String() string {
	ans := m.From().String() + m.To().String()
	if m.Promote() != Empty {
		ans += m.Promote().String()
	}
	return ans
}

func ParseMove(move string) Move {
	if len(move) < 4 || len(move) > 5 {
		return EncodeMove(None, None, Empty)
	}
	from := ParseSquare(move[0:2])
	to := ParseSquare(move[2:4])
	piece := Empty
	if len(move) == 5 {
		piece = ParsePiece(rune(move[4]))
	}
	return EncodeMove(from, to, piece)
}
