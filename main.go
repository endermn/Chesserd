package main

import (
	"log"
	"time"

	"github.com/endermn/Chlib"
	"github.com/endermn/bot/engine"
)

func main() {
	// fenStr, _ := chess.FEN("4k3/R7/1R6/8/8/8/8/4K3 w - - 0 1")
	game := chess.NewGame()
	start := time.Now()
	log.Printf("%v", engine.InitSearch(game, 4))
	end := time.Now()
	log.Printf("%v", end.Sub(start))
	log.Printf(game.Position().Board().Draw())
}
