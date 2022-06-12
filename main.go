package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/sergrom/BrutalManiaAutoplay/autoplayer"
	clrMap "github.com/sergrom/BrutalManiaAutoplay/color_map"
)

func main() {
	color.Blue("=== BrutalMania autoplay ===")
	time.Sleep(2 * time.Second)

	colorMap := clrMap.NewColorMap()
	if err := colorMap.LoadDir("./objects/"); err != nil {
		log.Fatalf(err.Error())
	}

	game := autoplayer.NewGame(colorMap)

	// Move mouse pointer to the game area
	game.MoveMouseToCenter()

	for {
		if game.IsOutOfArea() {
			// Break loop if mouse pointer got out of the game area
			break
		}

		switch {
		case game.StatusPlaying():
			game.DoPlay()
		case game.StatusGameOver():
			game.OpenBoxOrGoBack()
		case game.StatusInitial():
			game.ImproveWeaponsOrStartNewGame()
		}

		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()

	color.Blue("------------------------------------------------")
	color.Blue("Total games: %d", game.GamesCounter)
	color.Blue("Total time: %s", time.Since(game.Start).String())
	color.Blue("Total hits: %d", game.HitsCounter)
	color.Blue("------------------------------------------------")

	fmt.Println()
	color.Blue("Bye!")
}
