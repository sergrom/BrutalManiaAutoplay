package autoplayer

import (
	clrMap "github.com/sergrom/BrutalManiaAutoplay/color_map"
	"github.com/sergrom/BrutalManiaAutoplay/mouse"

	//"github.com/sergrom/BrutalManiaAutoplay"
	"image"
	"time"

	"github.com/fatih/color"
)

const (
	GameMinX int = 1920
	GameMaxX int = 2870
	GameMinY int = 140
	GameMaxY int = 630
)

var (
	// play screen points
	playScullPoint = image.Point{X: GameMinX + 935, Y: GameMinY + 30}

	// initial screen points
	initBPoint         = image.Point{X: GameMinX + 316, Y: GameMinY + 40}
	initMPoint         = image.Point{X: GameMinX + 523, Y: GameMinY + 45}
	initUpgradesPoint  = image.Point{X: GameMinX + 413, Y: GameMinY + 392}
	initPlayPoint      = image.Point{X: GameMinX + 552, Y: GameMinY + 391}
	initCupPoint       = image.Point{X: GameMinX + 369, Y: GameMinY + 390}
	initAvailablePoint = image.Point{X: GameMinX + 780, Y: GameMinY + 385}

	// achievements screen points
	achOpenPoint = image.Point{X: GameMinX + 627, Y: GameMinY + 131}
	achBackPoint = image.Point{X: GameMinX + 472, Y: GameMinY + 418}

	// game over screen points
	goOpenBoxPoint = image.Point{X: GameMinX + 481, Y: GameMinY + 245}
)

type Game struct {
	Start        time.Time
	ColorMap     *clrMap.ColorMap
	GamesCounter int
	HitsCounter  int
}

// NewGame ...
func NewGame(colorMap *clrMap.ColorMap) Game {
	g := Game{
		Start:    time.Now(),
		ColorMap: colorMap,
	}

	return g
}

// MoveMouseToCenter ...
func (g *Game) MoveMouseToCenter() {
	mouse.Move(image.Pt((GameMaxX+GameMinX)/2, (GameMaxY+GameMinY)/2))
}

// IsOutOfArea ...
func (g *Game) IsOutOfArea() bool {
	x, y, err := mouse.GetPointerCoordinates()
	if err != nil {
		color.Red("getMousePointerCoordinates error: %s", err.Error())
		return true
	}

	if x < GameMinX || x > GameMaxX || y < GameMinY || y > GameMaxY {
		return true
	}
	return false
}

// StatusInitial ...
func (g *Game) StatusInitial() bool {
	bPoint := TakeColorSample(initBPoint, 3)
	mPoint := TakeColorSample(initMPoint, 3)
	upgradePoint := TakeColorSample(initUpgradesPoint, 3)
	return bPoint.IsRed() && mPoint.IsYellow() && upgradePoint.IsYellow()
}

// StatusGameOver ...
func (g *Game) StatusGameOver() bool {
	scullPoint := TakeColorSample(playScullPoint, 5)
	avPoint := TakeColorSample(initAvailablePoint, 5)
	upgradePoint := TakeColorSample(initUpgradesPoint, 5)
	return scullPoint.IsBrown() && avPoint.IsBrown() && upgradePoint.IsBrown()
}

// OpenBoxOrGoBack ...
func (g *Game) OpenBoxOrGoBack() {
	mouse.Click(goOpenBoxPoint)
}

// StatusPlaying ...
func (g *Game) StatusPlaying() bool {
	return TakeColorSample(playScullPoint, 3).IsWhite()
}

func (g *Game) ImproveWeaponsOrStartNewGame() {
	if g.isImprovingWeaponsAvailable() {
		mouse.Click(initAvailablePoint) // improve weapons
	} else {
		mouse.Click(initCupPoint) // click  "cup" button
		time.Sleep(100 * time.Millisecond)

		// while there is any achievements - open them
		for g.isAchievementAvailable() {
			mouse.Click(achOpenPoint)
			time.Sleep(100 * time.Millisecond)
		}

		mouse.Click(achBackPoint) // click "Back" button
		time.Sleep(3 * time.Second)

		mouse.Click(initPlayPoint) // click "Play" button
		mouse.Move(image.Pt((GameMaxX+GameMinX)/2, (GameMaxY+GameMinY)/2))

		color.Green("========== NEW GAME ===========")
		g.GamesCounter++
	}
}

func (g *Game) isImprovingWeaponsAvailable() bool {
	return TakeColorSample(initAvailablePoint, 3).IsGreen()
}

func (g *Game) isAchievementAvailable() bool {
	return TakeColorSample(achOpenPoint, 3).IsGreen()
}
