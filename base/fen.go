package base

import (
	"errors"
	"fmt"
	. "github.com/ComputerSaysYeah/RookMills/api"
	"strconv"
	"strings"
)

func (g *gameSt) FromFEN(fen string) error {

	tokens := strings.Split(fen, " ")
	if len(tokens) != 6 {
		return errors.New("malformed FEN line not all tokens there")
	}
	rows := strings.Split(tokens[0], "/")
	if len(rows) != 8 {
		return errors.New("malformed FEN line, not eight lines")
	}

	g.Board().Reset()
	for r, row := range rows {
		c := 0
		for _, ch := range row {
			if ch >= '0' && ch <= '9' {
				c += int(ch) - '0'
			} else {
				g.Board().Set(Row8-Square(r*8)+Square(c), ParsePiece(ch))
				c++
			}
		}
	}

	if tokens[1] == "b" {
		g.SetMoveNext(Black)
	} else if tokens[1] == "w" {
		g.SetMoveNext(White)
	} else {
		return errors.New("non-parseable next player -2nd token- in FEN string")
	}

	WK, WQ, bk, bq := false, false, false, false
	for _, ch := range tokens[2] {
		switch ch {
		case 'K':
			WK = true
		case 'Q':
			WQ = true
		case 'k':
			bk = true
		case 'q':
			bq = true
		}
	}
	g.SetCastling(WK, WQ, bk, bq)

	g.SetEnPassant(ParseSquare(tokens[3]))

	if halfMove, err := strconv.Atoi(tokens[4]); err == nil {
		g.SetHalfMoveNo(halfMove)
	} else {
		return err
	}

	if moveNo, err := strconv.Atoi(tokens[5]); err == nil {
		g.SetMoveNo(moveNo)
	} else {
		return err
	}

	return nil
}

func (g *gameSt) ToFEN() string {

	b := g.Board()
	ans := ""

	for r := 7; r >= 0; r-- {
		empties := 0
		for c := Square(0); c < 8; c++ {
			piece := b.Get(Square(r)*OneRow + c)
			if piece.IsEmpty() {
				empties++
			} else {
				if empties > 0 {
					ans += fmt.Sprint(empties)
					empties = 0
				}
				ans += piece.String()
			}
		}
		if empties > 0 {
			ans += fmt.Sprint(empties)
		}
		if r != 0 {
			ans += "/"
		}
	}
	if g.MoveNext().Colour() == Black {
		ans += " b"
	} else {
		ans += " w"
	}
	ans += " "
	WK, WQ, bk, bq := g.Castling()
	if WK {
		ans += "K"
	}
	if WQ {
		ans += "Q"
	}
	if bk {
		ans += "k"
	}
	if bq {
		ans += "q"
	}
	if !WK && !WQ && !bk && !bq {
		ans += "-"
	}
	if !g.EnPassant().IsNone() {
		ans += " " + g.EnPassant().String()
	} else {
		ans += " -"
	}
	ans += " " + fmt.Sprint(g.HalfMoveNo())
	ans += " " + fmt.Sprint(g.MoveNo())

	return ans
}
