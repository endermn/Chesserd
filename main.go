package main

import (
	 _"fmt"
	"log"
	"sync"
	_ "time"

	"github.com/endermn/Chlib"
	"github.com/endermn/bot/engine"
)

func run(botTurn chess.Color, depth int) string {
	// fen, err := chess.FEN("4k3/R7/1R6/8/8/8/8/4K3 w - - 0 1")
	// if err != nil {
	// 	panic("Failed to parse fen string")
	// }
	game := chess.NewGame()

	for game.Outcome() == chess.NoOutcome {
		if game.Position().Turn() == chess.White {
			depth = 2
		} else {
			depth = 3
		}
		move := engine.InitSearch(game, depth)
		game.Move(move)
		log.Printf("%v", move)
	log.Printf(game.Position().Board().Draw())
	}
	// for game.Outcome() == chess.NoOutcome {
	// 	if game.Position().Turn() == botTurn {
	// 		start := time.Now()
	//
	// 		move := engine.InitSearch(game, depth)
	// 		log.Printf("%v", move)
	// 		game.Move(move)
	//
	// 		end := time.Now()
	// 		log.Printf("%v", end.Sub(start))
	// 	} else {
	// 		log.Printf("Your turn")
	//
	// 		var move string
	// 		fmt.Scanf("%s", &move)
	// 		game.MoveStr(move)
	// 	}
	// 	log.Printf(game.Position().Board().Draw())
	// }
	return game.Outcome().String()
}

func main() {
	var wg sync.WaitGroup
	outcomes := []string{}

	for range 1 {
		wg.Add(1)
		go func() {
			outcomes = append(outcomes, run(chess.White, 2))
			wg.Done()
		}()
	}
	wg.Wait()
	log.Printf("%v", outcomes)
}
