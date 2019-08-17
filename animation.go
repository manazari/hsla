package hsla

import (
	"image/color"
	"time"
)

//
type AnimatedColor struct {
	color.Color
	updateFunc func()
	paused     bool
}

//
func NewAnimatedColor(c color.Color, updateFunc func()) AnimatedColor {
	return AnimatedColor{c, updateFunc, false}
}

//
func (c *AnimatedColor) Start() {
	if c.paused == false {
		return
	}
	go func() {
		for range time.Tick(time.Millisecond * 60) {
			c.updateFunc()
			if c.IsPaused() {
				break
			}
		}
	}()

	c.paused = true
}

//
func (c *AnimatedColor) Stop() {
	c.paused = false
}

//
func (c AnimatedColor) IsPaused() bool {
	return c.paused
}
