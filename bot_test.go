// bot_test.go
package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"
	"testing"

	"github.com/endermn/Chlib"
)

// getBotMove runs the specified bot binary with the given FEN and depth,
// and returns the move generated by that bot.
func getBotMove(botPath, fen string, depth int) (string, error) {
	cmd := exec.Command(botPath, "-fen", fen, "-depth", fmt.Sprint(depth))

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func runGame(t *testing.T, wPath, bPath string, depth1, depth2 int) string {
	game := chess.NewGame()

	for game.Outcome() == chess.NoOutcome {
		// Determine which bot's turn it is and get the appropriate bot path and depth
		var botPath string
		var depth int
		if game.Position().Turn() == chess.White {
			botPath = wPath
			depth = depth1
		} else {
			botPath = bPath
			depth = depth2
		}

		// Get the current board state in FEN format
		fen := game.Position().String()

		// Run the bot to get its move
		notation := chess.AlgebraicNotation{}
		moveStr, err := getBotMove(botPath, fen, depth)
		if err != nil {
			t.Fatalf("Failed to get move string from bot (%s): %v", botPath, err)
		}

		move, err := notation.Decode(game.Position(), moveStr)
		if err != nil {
			t.Fatalf("Failed to get move from bot (%s): %v", botPath, err)
		}
		game.Move(move)
	}

	return game.Outcome().String()
}

func TestBotVersions(t *testing.T) {
	botCurrent := "./bin/bot_current"
	botStable := "./bin/bot_stable"

	// Run the test match between current and stable bot versions
	var wg sync.WaitGroup
	newWins, oldWins, draws := 0, 0, 0

	for i := range 10 {
		wg.Add(1)
		go t.Run("Current vs Stable", func(t *testing.T) {
			var outcome string
			if i%2 == 0 {
				outcome = runGame(t, botCurrent, botStable, 2, 2) // depths can vary
			} else {
				outcome = runGame(t, botStable, botCurrent, 2, 2) // depths can vary
			}
			if outcome == "" {
				t.Error("Expected a game outcome, but got an empty string")
			}

			switch outcome {
			case string(chess.BlackWon):
				if i%2 != 0 {
					newWins++
				} else {
					oldWins++
				}
			case string(chess.WhiteWon):
				if i%2 == 0 {
					newWins++
				} else {
					oldWins++
				}
			default:
				draws++
			}

			wg.Done()
		})
	}
	wg.Wait()
	t.Logf("NewWins: %d, OldWins: %d, Draws: %d", newWins, oldWins, draws)
	t.Logf("The new changes are better with: %f percent", float32(newWins)/float32(oldWins)*100-100)
}
