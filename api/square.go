package api

import (
	"fmt"
	"strconv"
	"strings"
)

func (s Square) String() string {
	if s.IsNone() {
		return "None"
	} else if s > 63 {
		return "!Invalid"
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
	if s > o {
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

func (s Square) NE() Square {
	if s == None || s.Row() == Row8 || s.Col() == ColH {
		return None
	}
	return s + OneRow + 1
}

func (s Square) EEN() Square {
	if s == None || s.Col() >= ColG || s.Row() == Row8 {
		return None
	}
	return s + 2 + OneRow
}

func (s Square) E() Square {
	if s == None || s.Col() == ColH {
		return None
	}
	return s + 1
}

func (s Square) EES() Square {
	if s == None || s.Col() >= ColG || s.Row() == Row1 {
		return None
	}
	return s + 2 - OneRow
}

func (s Square) SE() Square {
	if s == None || s.Col() == ColH || s.Row() == Row1 {
		return None
	}
	return s - OneRow + 1
}

func (s Square) SSE() Square {
	if s == None || s.Col() == ColH || s.Row() <= 2 {
		return None
	}
	return s - OneRow - OneRow + 1
}

func (s Square) S() Square {
	if s == None || s.Row() == Row1 {
		return None
	}
	return s - OneRow
}

func (s Square) SSW() Square {
	if s == None || s.Row() <= Row2 || s.Col() == ColA {
		return None
	}
	return s - OneRow - OneRow - 1
}

func (s Square) SW() Square {
	if s == None || s.Row() == Row1 || s.Col() == ColA {
		return None
	}
	return s - OneRow - 1
}

func (s Square) WWN() Square {
	if s == None || s.Col() <= ColB || s.Row() == Row8 {
		return None
	}
	return s + OneRow - 2
}

func (s Square) W() Square {
	if s == None || s.Col() == ColA {
		return None
	}
	return s - 1
}

func (s Square) WWS() Square {
	if s == None || s.Col() <= ColB || s.Row() == Row1 {
		return None
	}
	return s - OneRow - 2
}
func (s Square) NW() Square {
	if s == None || s.Col() == ColA || s.Row() == Row8 {
		return None
	}
	return s - 1 + OneRow
}

func (s Square) NNW() Square {
	if s == None || s.Col() == ColA || s.Row() >= Row7 {
		return None
	}
	return s - 1 + OneRow + OneRow
}

func (s Square) NNE() Square {
	if s == None || s.Col() == ColH || s.Row() >= Row7 {
		return None
	}
	return s + 1 + OneRow + OneRow
}

func (s Square) IsNone() bool {
	return s == None
}

func (s Square) IsValid() bool {
	return !s.IsNone()
}
