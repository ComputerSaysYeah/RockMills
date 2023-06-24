package api

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (s Square) String() string {
	if s > 63 {
		log.Fatal("square out of range")
	}
	return fmt.Sprintf("%c", rune(65+uint8(s)%8)) + strconv.Itoa((int(s)/8)+1)
}

func ParseSquare(st string) Square {
	ch := strings.ToLower(st)[0]
	if ch < 'a' || ch > 'h' || st[1] < '1' || st[1] > '8' {
		return None
	}
	return Square((ch - 'a') + ((st[1] - '1') * 8))
}

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
