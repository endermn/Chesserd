package engine

import (
	"sync"

	"github.com/endermn/Chlib"
)

var Inf float32 = 100000.0

func negamax(game *chess.Game, depth int, alpha float32, beta float32) float32 {
	if depth == 0 {
		return evaluate(game)
	}

	eval := -Inf
	moves := game.ValidMoves()

	for _, move := range moves {
		pos := game.Clone()
		pos.Move(move)
		evalMove := -negamax(pos, depth-1, -beta, -alpha)
		eval = max(eval, evalMove)
		alpha = max(alpha, eval)

		if alpha >= beta {
			break // Beta cutoff
		}
	}

	return eval
}

// InitSearch starts the concurrent Negamax search on the first level of moves.
func InitSearch(game *chess.Game, depth int) *chess.Move {
	moves := game.ValidMoves()
	eval := -Inf
	bestMove := moves[0]
	var wg sync.WaitGroup
	resultChan := make(chan struct {
		move *chess.Move
		eval float32
	}, len(moves))

	for _, move := range moves {
		wg.Add(1)
		go func(move *chess.Move) {
			defer wg.Done()
			pos := game.Clone()
			pos.Move(move)

			evalMove := -negamax(pos, depth-1, -Inf, Inf)

			resultChan <- struct {
				move *chess.Move
				eval float32
			}{move, evalMove}
		}(move)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		if result.eval > eval {
			eval = result.eval
			bestMove = result.move
		}
	}

	return bestMove
}
