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
		return Invalid
	}
	return Square((ch - 'a') + ((st[1] - '1') * 8))
}

func (p Piece) String() string {
	switch p {
	case Black + Pawn, White + Pawn:
		return "p"
	case Black + Rook, White + Rook:
		return "r"
	case Black + Knight, White + Knight:
		return "k"
	case Black + Bishop, White + Bishop:
		return "b"
	case Black + Queen, White + Queen:
		return "q"
	case Black + King, White + King:
		return "k"
	default:
		return "?"
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
