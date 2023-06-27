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

func (s Square) Row() Square {
	return s & 0xf8
}

func (s Square) Col() Square {
	return s % 8
}

func (s Square) Min(o Square) Square {
	if s < o {
		return s
	}
	return o
}

func (s Square) Max(o Square) Square {
	if s > 0 {
		return s
	}
	return o
}

func (s Square) N() Square {
	if s == None || s.Row() == Row8 {
		return None
	}
	return s + OneRow
}

func (s Square) S() Square {
	if s == None || s.Row() == Row1 {
		return None
	}
	return s - OneRow
}

func (s Square) W() Square {
	if s == None || s.Col() == 0 {
		return None
	}
	return s - 1
}

func (s Square) E() Square {
	if s == None || s.Col() == 7 {
		return None
	}
	return s + 1
}

func (s Square) IsNone() bool {
	return s == None
}
