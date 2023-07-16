package base

import (
	. "github.com/ComputerSaysYeah/RookMills/api"
	"github.com/ComputerSaysYeah/RookMills/speed"
	"testing"
)

func Benchmark_GenMovesStart(b *testing.B) {
	game := givenGame()
	for i := 0; i < b.N; i++ {
		game.ValidMoves().Return()
	}
}

func Benchmark_GenMovesStart_D4(b *testing.B) {
	game := givenGame()
	game.Move(ParseMove("D2D4"))
	for i := 0; i < b.N; i++ {
		game.ValidMoves().Return()
	}
}

func Benchmark_GenMovesStart_D4_E5(b *testing.B) {
	game := givenGame()
	game.Move(ParseMove("D2D4"))
	game.Move(ParseMove("E7E5"))
	for i := 0; i < b.N; i++ {
		game.ValidMoves().Return()
	}
}

func Benchmark_GenMovesStart_D4_E5_DXE5(b *testing.B) {
	game := givenGame()
	game.Move(ParseMove("D2D4"))
	game.Move(ParseMove("E7E5"))
	game.Move(ParseMove("D4E5"))
	for i := 0; i < b.N; i++ {
		game.ValidMoves().Return()
	}
}

func Benchmark_GenMovesStart_D4_E5_DXE5_D6(b *testing.B) {
	game := givenGame()
	game.Move(ParseMove("D2D4"))
	game.Move(ParseMove("E7E5"))
	game.Move(ParseMove("D4E5"))
	game.Move(ParseMove("D7D6"))
	for i := 0; i < b.N; i++ {
		game.ValidMoves().Return()
	}
}

func Benchmark_GenMovesStart_D4_E5_DXE5_D6_E4(b *testing.B) {
	game := givenGame()
	game.Move(ParseMove("D2D4"))
	game.Move(ParseMove("E7E5"))
	game.Move(ParseMove("D4E5"))
	game.Move(ParseMove("D7D6"))
	game.Move(ParseMove("E2E4"))
	for i := 0; i < b.N; i++ {
		game.ValidMoves().Return()
	}
}

func Benchmark_GenMovesStart_D4_E5_DXE5_D6_E4_D5(b *testing.B) {
	// bishops open, queens open, plenty of potential moves
	game := givenGame()
	game.Move(ParseMove("D2D4"))
	game.Move(ParseMove("E7E5"))
	game.Move(ParseMove("D4E5"))
	game.Move(ParseMove("D7D6"))
	game.Move(ParseMove("E2E4"))
	game.Move(ParseMove("D6D5"))
	//println(game.Board().String())
	for i := 0; i < b.N; i++ {
		game.ValidMoves().Return()
	}
}

func Benchmark_GenMove_Interesting(b *testing.B) {
	game := givenGame()
	_ = game.FromFEN("8/5N1n/2n2k1P/8/6PN/1B2R3/8/r2Q2KR w - - 0 1")
	for i := 0; i < b.N; i++ {
		game.ValidMoves().Return()
	}
}

func Benchmark_GenMove_FromStart_Depth4(b *testing.B) {
	b.ReportAllocs()
	gamePool := speed.NewLeanPool(64, givenGame)
	game := gamePool.Lease()
	total := 0
	for i := 0; i < b.N; i++ {
		nodes := goDepth(game, gamePool, 4)
		b.ReportMetric(float64(nodes), "nodes/op")
		total += nodes
	}
	b.ReportMetric(float64(total), "nodes")
}

func goDepth(game Game, gamePool speed.Pool[Game], depth int) int {
	if depth == 0 {
		return 0
	}
	nodes := 0
	moves := game.ValidMoves()
	for moves.HasNext() {
		subGame := gamePool.Lease()
		subGame.CopyFrom(game)
		subGame.Move(moves.Next())
		nodes++
		nodes += goDepth(subGame, gamePool, depth-1)
		subGame.Return()
	}
	moves.Return()
	return nodes
}
