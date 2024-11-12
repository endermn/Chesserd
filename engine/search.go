package engine

import (
	"log"

	"github.com/endermn/Chlib"
)

var Inf float32 = 100000.0

func negamax(game *chess.Game, depth int, alpha float32, beta float32) float32 {
	if depth == 0 {
		return evaluate(game)
	}

	eval := -Inf
	moves := game.ValidMoves()
	game.Position()

	for _, move := range moves {
		pos := game.Clone()
		pos.Move(move)
		evalMove := -negamax(pos, depth-1, -beta, -alpha)

		eval = max(eval, evalMove)
		alpha = max(alpha, eval)

		if alpha >= beta {
			break
		}
	}
	return eval
}

func InitSearch(game *chess.Game, depth int) *chess.Move {
	moves := game.ValidMoves()
	eval := -Inf
	var bestMove *chess.Move

	for _, move := range moves {
		pos := game.Clone()
		pos.Move(move)
		evalMove := -negamax(pos, depth, -Inf, Inf)
		if evalMove > eval {
			eval = evalMove
			bestMove = move
		}
	}
	log.Printf("%v", eval)
	return bestMove
}
