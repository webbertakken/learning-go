package canvas

import "fyne.io/fyne/v2"

func (c *Canvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	c.CanvasOffset.X += xDiff
	c.CanvasOffset.Y += yDiff

	c.Refresh()
}

func (c *Canvas) scale(direction int) {
	switch {
	case direction > 0:
		c.PixelRatio += 1
	case direction < 0:
		if c.PixelRatio >= 2 {
			c.PixelRatio -= 1
		}
	default:
		c.PixelRatio = 10
	}
}
