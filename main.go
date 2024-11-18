package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/endermn/Chlib"
	"github.com/endermn/bot/engine"
)

func run(botTurn chess.Color, depth int) string {
	// fen, err := chess.FEN("4k3/R7/1R6/8/8/8/8/4K3 w - - 0 1")
	// if err != nil {
	// 	panic("Failed to parse fen string")
	// }
	game := chess.NewGame()
	playSelf := true

	for game.Outcome() == chess.NoOutcome {
		if playSelf {
			switch game.Position().Turn() {
			case chess.White:
				depth = 2
			case chess.Black:
				depth = 3
			}
			move := engine.InitSearch(game, depth)
			game.Move(move)
			log.Printf("%v", move)
			continue
		}

		if game.Position().Turn() == botTurn {
			start := time.Now()

			move := engine.InitSearch(game, depth)
			log.Printf("%v", move)
			game.Move(move)

			end := time.Now()
			log.Printf("%v", end.Sub(start))
		} else {
			log.Printf("Your turn")

			var move string
			fmt.Scanf("%s", &move)
			game.MoveStr(move)
		}
		log.Printf(game.Position().Board().Draw())
	}
	return game.Outcome().String()
}

func main() {
	fenStr := flag.String("fen", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", "FEN string representing the current board state")
	depth := flag.Int("depth", 2, "Search depth for bot move")
	flag.Parse()

	if *fenStr == "" {
		log.Fatal("No FEN string provided")
	}

	fen, err := chess.FEN(*fenStr)
	if err != nil {
		log.Fatalf("Invalid FEN: %v", err)
	}

	game := chess.NewGame(fen)

	move := engine.InitSearch(game, *depth)
	// notation := chess.AlgebraicNotation{}
	fmt.Print(move)
}
