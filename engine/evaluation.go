package engine

import (
	"github.com/endermn/Chlib"
)

func calcMobility(game *chess.Game) float32 {
	var mobilities [2]int
	firstColor := game.Position().Turn()

	mobilities[0] = len(game.ValidMoves())
	game.SetTurn(firstColor.Other())
	mobilities[1] = len(game.ValidMoves())

	return 0.1 * float32(mobilities[0]-mobilities[1])
}

func evaluate(game *chess.Game) float32 {
	var eval float32 = 0.0
	var mul int

	if game.Position().Turn() == chess.White {
		mul = 1
	} else {
		mul = -1
	}

	for _, piece := range game.Position().Board().SquareMap() {
		var colorValue int

		if piece.Color() == chess.White {
			colorValue = 1
		} else {
			colorValue = -1
		}

		switch piece.Type() {
		case chess.Pawn:
			eval += float32(1 * colorValue)
		case chess.Knight, chess.Bishop:
			eval += float32(3 * colorValue)
		case chess.Rook:
			eval += float32(5 * colorValue)
		case chess.Queen:
			eval += float32(9 * colorValue)
		}

	}
	eval += calcMobility(game) * float32(mul)

	return 0
}
