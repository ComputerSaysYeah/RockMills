package util

import (
	"errors"
	. "github.com/ComputerSaysYeah/RookMills/api"
	"strconv"
	"strings"
)

func ParseFEN(game Game, fen string) error {

	tokens := strings.Split(fen, " ")
	if len(tokens) != 6 {
		return errors.New("malformed FEN line not all tokens there")
	}
	rows := strings.Split(tokens[0], "/")
	if len(rows) != 8 {
		return errors.New("malformed FEN line, not eight lines")
	}

	game.Board().Reset()
	for r, row := range rows {
		c := 0
		for _, ch := range row {
			if ch >= '0' && ch <= '9' {
				c += int(ch) - '0'
			} else {
				game.Board().Set(Row8-Square(r*8)+Square(c), ParsePiece(ch))
				c++
			}
		}
	}

	if tokens[1] == "b" {
		game.SetMoveNext(Black)
	} else if tokens[1] == "w" {
		game.SetMoveNext(White)
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
	game.SetCastling(WK, WQ, bk, bq)

	game.SetEnPassant(ParseSquare(tokens[3]))

	if halfMove, err := strconv.Atoi(tokens[4]); err == nil {
		game.SetHalfMoveNo(halfMove)
	} else {
		return err
	}

	if moveNo, err := strconv.Atoi(tokens[5]); err == nil {
		game.SetMoveNo(moveNo)
	} else {
		return err
	}

	return nil
}
