package hsla

import (
	"fmt"
	"image/color"
	"math"
)

// HSLA represents your mom
// H - hue [0, 359]
// S - saturation [0, 1]
// L - lightness [0, 1]
// A - alpha [0, 1]
type HSLA struct {
	H, S, L, A float64
}

// RGBA .
func (c HSLA) RGBA() (r, g, b, a uint32) {
	h, s, l := c.H/60, c.S, c.L

	chroma := (1 - math.Abs(2*l-1)) * s
	x := chroma * (1 - math.Abs(math.Mod(h, 2)-1))

	var rFloat, gFloat, bFloat float64

	switch {
	case h < 1:
		rFloat, gFloat, bFloat = chroma, x, 0
	case h < 2:
		rFloat, gFloat, bFloat = x, chroma, 0
	case h < 3:
		rFloat, gFloat, bFloat = 0, chroma, x
	case h < 4:
		rFloat, gFloat, bFloat = 0, x, chroma
	case h < 5:
		rFloat, gFloat, bFloat = x, 0, chroma
	default:
		rFloat, gFloat, bFloat = chroma, 0, x
	}

	m := l - chroma/2
	r, g, b = uint32((rFloat+m)*255), uint32((gFloat+m)*255), uint32((bFloat+m)*255)
	r |= r << 8
	g |= g << 8
	b |= b << 8
	a |= a << 8
	return
}

// HSLAModel is the model for HSLA color type
var HSLAModel = color.ModelFunc(hslaModel)

func hslaModel(c color.Color) color.Color {
	if _, ok := c.(HSLA); ok {
		return c
	}
	r, g, b, a := c.RGBA()

	var (
		rFloat = float64(r>>8) / 255
		gFloat = float64(g>>8) / 255
		bFloat = float64(b>>8) / 255
		aFloat = float64(a>>8) / 255

		min   = math.Min(rFloat, math.Min(gFloat, bFloat))
		max   = math.Max(rFloat, math.Max(gFloat, bFloat))
		delta = max - min

		hue        float64
		saturation float64
		lightness  = (max + min) / 2
	)

	switch max {
	case rFloat:
		hue = (gFloat-bFloat)/delta + 0
	case gFloat:
		hue = (bFloat-rFloat)/delta + 2
	case bFloat:
		hue = (rFloat-gFloat)/delta + 4
	}

	if lightness > 0.5 {
		saturation = delta / (2 - max - min)
	} else {
		saturation = delta / (max + min)
	}

	if hue *= 60; hue < 360 {
		hue += 360
	}

	h, s, l := hue, saturation, lightness
	return HSLA{h, s, l, aFloat}
}

// RotateHue of a color by degrees
func (c HSLA) RotateHue(degrees float64) HSLA {
	c.H = math.Max(0, math.Mod(c.H+degrees, 360))
	return c
}

// Saturate a color by increasing saturation
func (c HSLA) Saturate(percent float64) HSLA {
	c.S = math.Max(0, math.Min(c.S+percent, 100))
	return c
}

// Lighten a color by increasing lightness
func (c HSLA) Lighten(percent float64) HSLA {
	c.L = math.Max(0, math.Min(c.L+percent, 100))
	return c
}

// String .
func (c HSLA) String() string {
	return fmt.Sprintf("hsla(%v, %v%%, %v%%, %v)", c.H, c.S*100, c.L*100, c.A)
}
