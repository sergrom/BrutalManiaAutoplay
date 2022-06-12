package autoplayer

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/kbinani/screenshot"
	"github.com/sergrom/BrutalManiaAutoplay/mouse"
	"image"
	"math"
	"time"
)

const (
	Velocity = 80.0 // px/sec

	centerX = (GameMinX + GameMaxX) / 2
	centerY = (GameMinY + GameMaxY) / 2

	oEnemy  = "enemy"
	oOrbiz  = "orbiz"
	oBorder = "border"

	oEnemyPercent  = 0.1
	oOrbizPercent  = 0.02
	oBorderPercent = 0.2
)

// Area ...
type Area struct {
	DX, DY   int
	Radius   int
	Distance int
}

// Object ...
type Object struct {
	Name    string
	Area    Area
	Percent float64
}

var orbizTarget = struct {
	isActual bool
	point    image.Point
	tUntil   time.Time
}{}

var areas = [][]Area{
	{Area{DX: 65, DY: 0, Distance: 65, Radius: 20}, Area{DX: 56, DY: 32, Distance: 65, Radius: 20}, Area{DX: 32, DY: 56, Distance: 65, Radius: 20}, Area{DX: 0, DY: 64, Distance: 65, Radius: 20}, Area{DX: -32, DY: 56, Distance: 65, Radius: 20}, Area{DX: -56, DY: 32, Distance: 65, Radius: 20}, Area{DX: -64, DY: 0, Distance: 65, Radius: 20}, Area{DX: -56, DY: -32, Distance: 65, Radius: 20}, Area{DX: -32, DY: -56, Distance: 65, Radius: 20}, Area{DX: 0, DY: -64, Distance: 65, Radius: 20}, Area{DX: 32, DY: -56, Distance: 65, Radius: 20}, Area{DX: 56, DY: -32, Distance: 65, Radius: 20}},
	{Area{DX: 105, DY: 0, Distance: 105, Radius: 20}, Area{DX: 99, DY: 32, Distance: 105, Radius: 20}, Area{DX: 84, DY: 61, Distance: 105, Radius: 20}, Area{DX: 61, DY: 85, Distance: 105, Radius: 20}, Area{DX: 32, DY: 99, Distance: 105, Radius: 20}, Area{DX: 0, DY: 104, Distance: 105, Radius: 20}, Area{DX: -32, DY: 99, Distance: 105, Radius: 20}, Area{DX: -61, DY: 84, Distance: 105, Radius: 20}, Area{DX: -85, DY: 61, Distance: 105, Radius: 20}, Area{DX: -99, DY: 32, Distance: 105, Radius: 20}, Area{DX: -104, DY: 0, Distance: 105, Radius: 20}, Area{DX: -99, DY: -32, Distance: 105, Radius: 20}, Area{DX: -84, DY: -62, Distance: 105, Radius: 20}, Area{DX: -61, DY: -85, Distance: 105, Radius: 20}, Area{DX: -31, DY: -100, Distance: 105, Radius: 20}, Area{DX: 0, DY: -104, Distance: 105, Radius: 20}, Area{DX: 32, DY: -99, Distance: 105, Radius: 20}, Area{DX: 62, DY: -84, Distance: 105, Radius: 20}, Area{DX: 85, DY: -61, Distance: 105, Radius: 20}, Area{DX: 100, DY: -31, Distance: 105, Radius: 20}},
	{Area{DX: 145, DY: 0, Distance: 145, Radius: 20}, Area{DX: 141, DY: 32, Distance: 145, Radius: 20}, Area{DX: 130, DY: 62, Distance: 145, Radius: 20}, Area{DX: 113, DY: 90, Distance: 145, Radius: 20}, Area{DX: 90, DY: 113, Distance: 145, Radius: 20}, Area{DX: 62, DY: 130, Distance: 145, Radius: 20}, Area{DX: 32, DY: 141, Distance: 145, Radius: 20}, Area{DX: 0, DY: 144, Distance: 145, Radius: 20}, Area{DX: -32, DY: 141, Distance: 145, Radius: 20}, Area{DX: -63, DY: 130, Distance: 145, Radius: 20}, Area{DX: -90, DY: 113, Distance: 145, Radius: 20}, Area{DX: -113, DY: 90, Distance: 145, Radius: 20}, Area{DX: -130, DY: 62, Distance: 145, Radius: 20}, Area{DX: -141, DY: 31, Distance: 145, Radius: 20}, Area{DX: -144, DY: 0, Distance: 145, Radius: 20}, Area{DX: -141, DY: -32, Distance: 145, Radius: 20}, Area{DX: -130, DY: -63, Distance: 145, Radius: 20}, Area{DX: -112, DY: -90, Distance: 145, Radius: 20}, Area{DX: -89, DY: -113, Distance: 145, Radius: 20}, Area{DX: -62, DY: -130, Distance: 145, Radius: 20}, Area{DX: -31, DY: -141, Distance: 145, Radius: 20}, Area{DX: 0, DY: -144, Distance: 145, Radius: 20}, Area{DX: 33, DY: -141, Distance: 145, Radius: 20}, Area{DX: 63, DY: -130, Distance: 145, Radius: 20}, Area{DX: 91, DY: -112, Distance: 145, Radius: 20}, Area{DX: 113, DY: -89, Distance: 145, Radius: 20}, Area{DX: 131, DY: -62, Distance: 145, Radius: 20}, Area{DX: 141, DY: -31, Distance: 145, Radius: 20}},
	{Area{DX: 185, DY: 0, Distance: 185, Radius: 20}, Area{DX: 182, DY: 32, Distance: 185, Radius: 20}, Area{DX: 173, DY: 63, Distance: 185, Radius: 20}, Area{DX: 160, DY: 92, Distance: 185, Radius: 20}, Area{DX: 141, DY: 119, Distance: 185, Radius: 20}, Area{DX: 118, DY: 141, Distance: 185, Radius: 20}, Area{DX: 92, DY: 160, Distance: 185, Radius: 20}, Area{DX: 63, DY: 173, Distance: 185, Radius: 20}, Area{DX: 31, DY: 182, Distance: 185, Radius: 20}, Area{DX: 0, DY: 184, Distance: 185, Radius: 20}, Area{DX: -32, DY: 182, Distance: 185, Radius: 20}, Area{DX: -63, DY: 173, Distance: 185, Radius: 20}, Area{DX: -92, DY: 160, Distance: 185, Radius: 20}, Area{DX: -119, DY: 141, Distance: 185, Radius: 20}, Area{DX: -142, DY: 118, Distance: 185, Radius: 20}, Area{DX: -160, DY: 92, Distance: 185, Radius: 20}, Area{DX: -174, DY: 62, Distance: 185, Radius: 20}, Area{DX: -182, DY: 31, Distance: 185, Radius: 20}, Area{DX: -184, DY: 0, Distance: 185, Radius: 20}, Area{DX: -182, DY: -32, Distance: 185, Radius: 20}, Area{DX: -173, DY: -63, Distance: 185, Radius: 20}, Area{DX: -159, DY: -93, Distance: 185, Radius: 20}, Area{DX: -141, DY: -119, Distance: 185, Radius: 20}, Area{DX: -118, DY: -142, Distance: 185, Radius: 20}, Area{DX: -91, DY: -160, Distance: 185, Radius: 20}, Area{DX: -62, DY: -174, Distance: 185, Radius: 20}, Area{DX: -31, DY: -182, Distance: 185, Radius: 20}, Area{DX: 0, DY: -184, Distance: 185, Radius: 20}, Area{DX: 33, DY: -182, Distance: 185, Radius: 20}, Area{DX: 64, DY: -173, Distance: 185, Radius: 20}, Area{DX: 93, DY: -159, Distance: 185, Radius: 20}, Area{DX: 119, DY: -141, Distance: 185, Radius: 20}, Area{DX: 142, DY: -118, Distance: 185, Radius: 20}, Area{DX: 160, DY: -91, Distance: 185, Radius: 20}, Area{DX: 174, DY: -62, Distance: 185, Radius: 20}, Area{DX: 182, DY: -30, Distance: 185, Radius: 20}},
	{Area{DX: 225, DY: 0, Distance: 225, Radius: 20}, Area{DX: 222, DY: 32, Distance: 225, Radius: 20}, Area{DX: 215, DY: 63, Distance: 225, Radius: 20}, Area{DX: 204, DY: 93, Distance: 225, Radius: 20}, Area{DX: 189, DY: 121, Distance: 225, Radius: 20}, Area{DX: 169, DY: 147, Distance: 225, Radius: 20}, Area{DX: 147, DY: 170, Distance: 225, Radius: 20}, Area{DX: 121, DY: 189, Distance: 225, Radius: 20}, Area{DX: -122, DY: 188, Distance: 225, Radius: 20}, Area{DX: -147, DY: 169, Distance: 225, Radius: 20}, Area{DX: -170, DY: 146, Distance: 225, Radius: 20}, Area{DX: -189, DY: 121, Distance: 225, Radius: 20}, Area{DX: -204, DY: 92, Distance: 225, Radius: 20}, Area{DX: -216, DY: 62, Distance: 225, Radius: 20}, Area{DX: -222, DY: 31, Distance: 225, Radius: 20}, Area{DX: -224, DY: 0, Distance: 225, Radius: 20}, Area{DX: -222, DY: -32, Distance: 225, Radius: 20}, Area{DX: -215, DY: -64, Distance: 225, Radius: 20}, Area{DX: -204, DY: -94, Distance: 225, Radius: 20}, Area{DX: -188, DY: -122, Distance: 225, Radius: 20}, Area{DX: -169, DY: -148, Distance: 225, Radius: 20}, Area{DX: -146, DY: -170, Distance: 225, Radius: 20}, Area{DX: -120, DY: -189, Distance: 225, Radius: 20}, Area{DX: 122, DY: -188, Distance: 225, Radius: 20}, Area{DX: 148, DY: -169, Distance: 225, Radius: 20}, Area{DX: 170, DY: -146, Distance: 225, Radius: 20}, Area{DX: 190, DY: -120, Distance: 225, Radius: 20}, Area{DX: 205, DY: -92, Distance: 225, Radius: 20}, Area{DX: 216, DY: -61, Distance: 225, Radius: 20}, Area{DX: 222, DY: -30, Distance: 225, Radius: 20}},
	{Area{DX: 265, DY: 0, Distance: 265, Radius: 20}, Area{DX: 263, DY: 31, Distance: 265, Radius: 20}, Area{DX: 257, DY: 63, Distance: 265, Radius: 20}, Area{DX: 247, DY: 94, Distance: 265, Radius: 20}, Area{DX: 234, DY: 123, Distance: 265, Radius: 20}, Area{DX: 217, DY: 150, Distance: 265, Radius: 20}, Area{DX: 198, DY: 175, Distance: 265, Radius: 20}, Area{DX: -198, DY: 175, Distance: 265, Radius: 20}, Area{DX: -218, DY: 149, Distance: 265, Radius: 20}, Area{DX: -234, DY: 122, Distance: 265, Radius: 20}, Area{DX: -248, DY: 93, Distance: 265, Radius: 20}, Area{DX: -257, DY: 62, Distance: 265, Radius: 20}, Area{DX: -263, DY: 31, Distance: 265, Radius: 20}, Area{DX: -264, DY: 0, Distance: 265, Radius: 20}, Area{DX: -262, DY: -32, Distance: 265, Radius: 20}, Area{DX: -257, DY: -64, Distance: 265, Radius: 20}, Area{DX: -247, DY: -94, Distance: 265, Radius: 20}, Area{DX: -234, DY: -124, Distance: 265, Radius: 20}, Area{DX: -217, DY: -151, Distance: 265, Radius: 20}, Area{DX: -197, DY: -176, Distance: 265, Radius: 20}, Area{DX: 199, DY: -174, Distance: 265, Radius: 20}, Area{DX: 219, DY: -149, Distance: 265, Radius: 20}, Area{DX: 235, DY: -121, Distance: 265, Radius: 20}, Area{DX: 248, DY: -92, Distance: 265, Radius: 20}, Area{DX: 257, DY: -61, Distance: 265, Radius: 20}, Area{DX: 263, DY: -30, Distance: 265, Radius: 20}},
}

// DoPlay ...
func (g *Game) DoPlay() {
	img, err := screenshot.CaptureRect(image.Rect(centerX-300, centerY-200, centerX+300, centerY+200))
	if err != nil {
		color.Red("Taking screenshot error: %s", err.Error())
	}

	orbises := make([]Object, 0)
	borders := make([]Object, 0)
	for _, loop := range areas {
		enemies := make([]Object, 0)
		orbizesLoop := make([]Object, 0)
		bordersLoop := make([]Object, 0)
		for _, area := range loop {
			for _, obj := range g.lookArea(img, area) {
				switch {
				case oEnemy == obj.Name && obj.Percent > oEnemyPercent:
					enemies = append(enemies, obj)
				case oOrbiz == obj.Name && obj.Percent > oOrbizPercent:
					orbizesLoop = append(orbizesLoop, obj)
				case oBorder == obj.Name && obj.Percent > oBorderPercent:
					bordersLoop = append(bordersLoop, obj)
				}
			}
		}

		if len(enemies) > 0 {
			msg(oEnemy, fmt.Sprintf("I see %d enemies.", len(enemies)))
			// go toward them or hit them
			if enemies[0].Area.Distance <= 150 {
				dir := getDirection(enemies[0].Area)
				mouse.Click(turnToHit(dir))
				mouse.Move(dir)
				msg("hit", "Hit!")
				g.HitsCounter++
				return
			}

			mouse.Move(getDirection(enemies[0].Area))
			orbizTarget.isActual = false
			return
		}

		if len(orbises) == 0 {
			orbises = append(orbises, orbizesLoop...)
		}

		if len(borders) == 0 {
			borders = append(borders, bordersLoop...)
		}
	}

	if len(orbises) > 0 {
		msg(oOrbiz, "I eat orbises.")
		if orbizTarget.isActual && orbizTarget.tUntil.After(time.Now()) {
			return
		}

		// go toward them
		var biggestOrbiz *Object
		for _, orb := range orbises {
			if biggestOrbiz == nil || biggestOrbiz.Percent < orb.Percent {
				biggestOrbiz = &orb
			}
		}

		tUntil := time.Now().Add(time.Duration(float64(biggestOrbiz.Area.Distance)/Velocity) * time.Second)
		if biggestOrbiz.Area.Distance < 75 {
			tUntil = tUntil.Add(500 * time.Millisecond)
		}

		orbizTarget = struct {
			isActual bool
			point    image.Point
			tUntil   time.Time
		}{isActual: true, point: getDirection(biggestOrbiz.Area), tUntil: tUntil}

		mouse.Move(orbizTarget.point)
		return
	}

	if len(borders) > 0 {
		// turn opposite to borders
		dxSum, dySum := 0, 0
		for _, bObj := range borders {
			dxSum += bObj.Area.DX
			dySum += bObj.Area.DY
		}

		dir := getDirection(Area{DX: -dxSum, DY: -dySum, Distance: int(math.Sqrt(float64(dxSum*dxSum + dySum*dySum)))})
		mouse.Move(dir)
	}
}

// lookArea ...
func (g Game) lookArea(img *image.RGBA, area Area) []Object {
	imgCenterX, imgCenterY := img.Bounds().Max.X/2, img.Bounds().Max.Y/2
	areaCenterX, areaCenterY := imgCenterX+area.DX, imgCenterY+area.DY

	cntAll := 0
	objectsMap := make(map[string]int, 0)
	for dy := -area.Radius; dy < area.Radius; dy += 2 {
		dx := int(math.Sqrt(float64(area.Radius*area.Radius) - float64(dy*dy)))
		for x := areaCenterX - dx; x < areaCenterX+dx; x += 2 {
			cntAll++
			rc, gc, bc, _ := img.At(x, areaCenterY+dy).RGBA()
			obj := g.ColorMap.Which(rc, gc, bc)
			if obj != "" {
				objectsMap[obj]++
			}
		}
	}

	objects := make([]Object, 0)
	for objName, cnt := range objectsMap {
		percent := float64(cnt) / float64(cntAll)
		objects = append(objects, Object{
			Name:    objName,
			Area:    area,
			Percent: percent,
		})
	}

	return objects
}

func getDirection(area Area) image.Point {
	k := 100.0 / float64(area.Distance)
	return image.Pt(centerX+int(float64(area.DX)*k), centerY+int(float64(area.DY)*k))
}

func turnToHit(point image.Point) image.Point {
	dx, dy := point.X-centerX, point.Y-centerY
	return image.Point{X: centerX + dy, Y: centerY - dx}
}

var msgKind string

func msg(kind, msg string) {
	if msgKind != kind {
		fmt.Print("\n")
	} else {
		fmt.Print("\r")
	}
	fmt.Print(">> " + msg)
	msgKind = kind
}
