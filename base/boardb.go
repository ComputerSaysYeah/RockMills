package base

import (
	"fmt"
	. "github.com/ComputerSaysYeah/RookMills/api"
	"hash/crc64"
	"unsafe"
)

var CRC64IsoTable *crc64.Table
var pieces = map[Piece]rune{}

func init() {
	CRC64IsoTable = crc64.MakeTable(crc64.ISO)
	pieces[White+Rook] = '♖'
	pieces[White+Knight] = '♘'
	pieces[White+Bishop] = '♗'
	pieces[White+Queen] = '♕'
	pieces[White+King] = '♔'
	pieces[White+Pawn] = '♙'
	pieces[Black+Rook] = '♜'
	pieces[Black+Knight] = '♞'
	pieces[Black+Bishop] = '♝'
	pieces[Black+Queen] = '♛'
	pieces[Black+King] = '♚'
	pieces[Black+Pawn] = '♟'
	pieces[Empty] = ' '
}

type BoardB struct {
	squares  [64]Piece
	returner func(any)
}

func NewBoardB() Board {
	return &BoardB{squares: [64]Piece{}, returner: nil}
}

func (b *BoardB) Get(s Square) Piece {
	return b.squares[s]
}
func (b *BoardB) Set(s Square, p Piece) {
	b.squares[s] = p
}

func (b *BoardB) Hash() uint64 {
	pb := &b.squares[0]
	up := unsafe.Pointer(pb)
	pi := (*[64]byte)(up)
	return crc64.Checksum(pi[:], CRC64IsoTable)
	//BenchmarkBoardB_Hash-8   	31910734	        36.79 ns/op

	//hash := uint64(17)
	//for i := 0; i < len(b.squares); i++ {
	//	hash = hash<<5 + hash
	//	hash += uint64(b.squares[i])
	//}
	//return hash
	//BenchmarkBoardB_Hash-8   	25488656	        44.49 ns/op

}

func (b *BoardB) SetStartingPieces() {
	b.Reset()

	b.Set(Row1+ColA, White+Rook)
	b.Set(Row1+ColB, White+Knight)
	b.Set(Row1+ColC, White+Bishop)
	b.Set(Row1+ColD, White+Queen)
	b.Set(Row1+ColE, White+King)
	b.Set(Row1+ColF, White+Bishop)
	b.Set(Row1+ColG, White+Knight)
	b.Set(Row1+ColH, White+Rook)

	b.Set(Row8+ColA, Black+Rook)
	b.Set(Row8+ColB, Black+Knight)
	b.Set(Row8+ColC, Black+Bishop)
	b.Set(Row8+ColD, Black+Queen)
	b.Set(Row8+ColE, Black+King)
	b.Set(Row8+ColF, Black+Bishop)
	b.Set(Row8+ColG, Black+Knight)
	b.Set(Row8+ColH, Black+Rook)

	for col := ColA; col <= ColH; col++ {
		b.Set(Row2+col, White+Pawn)
		b.Set(Row7+col, Black+Pawn)
	}
}

func (b *BoardB) String() string {
	res := ""

	for row := Row8; row <= Row8; row -= 8 {
		for col := ColA; col <= ColH; col++ {
			if b.Get(row+col) == Empty {
				ofs := (row>>3)%2 + 1
				if (row+col+ofs)%2 == 0 {
					res += "\u2591"
				} else {
					res += " "
				}
			} else {
				res += fmt.Sprintf("%c", pieces[b.Get(row+col)])
			}
		}
		res += "\n"
	}
	return res
}

func (b *BoardB) CopyFrom(o *Board) {
	copy(b.squares[:], (*o).(*BoardB).squares[:])
}

func (b *BoardB) Reset() {
	b.squares[0] = Empty
	b.squares[1] = Empty
	copy(b.squares[2:], b.squares[0:2])
	copy(b.squares[4:], b.squares[0:4])
	copy(b.squares[8:], b.squares[0:8])
	copy(b.squares[16:], b.squares[0:16])
	copy(b.squares[32:], b.squares[0:32])
	//for i := 0; i < len(b.squares); i++ {
	//	b.squares[i] = Empty
	//}
	// 13.66ns vs 18.49ms doing a loop (about 30% faster)
}

func (b *BoardB) Return() {
	b.returner(b)
}

func (b *BoardB) SetReturnerFn(returner func(any)) {
	b.returner = returner
}
