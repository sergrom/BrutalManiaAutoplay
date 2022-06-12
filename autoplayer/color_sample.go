package autoplayer

import (
	"image"

	"github.com/fatih/color"
	"github.com/kbinani/screenshot"
)

// ColorSample ...
type ColorSample struct {
	r, g, b int64
}

// TakeColorSample ...
func TakeColorSample(point image.Point, delta int) ColorSample {
	img, err := screenshot.CaptureRect(image.Rect(point.X-delta, point.Y-delta, point.X+delta, point.Y+delta))
	if err != nil {
		color.Red("TakingColorSample error: %s", err.Error())
		return ColorSample{}
	}

	var rSum, gSum, bSum, cnt int64
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			rSum, gSum, bSum = rSum+int64(r), gSum+int64(g), bSum+int64(b)
			cnt++
		}
	}

	return ColorSample{
		r: rSum / cnt,
		g: gSum / cnt,
		b: bSum / cnt,
	}
}

// IsRed ...
func (c ColorSample) IsRed() bool {
	return c.r > c.g+35000 && c.r > c.b+30000 && c.r > 40000 && c.r < 60000
}

// IsPurple ...
func (c ColorSample) IsPurple() bool {
	return c.b > c.g+15000 && c.r > c.g+15000 && c.b > 20000 && c.b < 40000
}

// IsYellow ..
func (c ColorSample) IsYellow() bool {
	return c.r > c.b+15000 && c.g > c.b+15000 && c.r > 35000 && c.r < 65000
}

// IsGreen ...
func (c ColorSample) IsGreen() bool {
	return c.g > c.r+10000 && c.g > c.b+10000 && c.g > 35000 && c.g < 60000
}

// IsWhite ...
func (c ColorSample) IsWhite() bool {
	return c.r > 40000 && c.g > 40000 && c.b > 40000
}

// IsBlack ..
func (c ColorSample) IsBlack() bool {
	return c.r < 12000 && c.g < 12000 && c.b < 12000
}

// IsBrown ...
func (c ColorSample) IsBrown() bool {
	return c.r > c.g+5000 && c.r > c.b+5000 && c.r > 10000 && c.r < 20000
}
