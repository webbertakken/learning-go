package canvas

import (
	"image"
	"image/color"
)

func (c *Canvas) LoadImage(img image.Image, format string) {
	dimensions := img.Bounds()

	c.CanvasConfig.Cols = dimensions.Dx()
	c.CanvasConfig.Rows = dimensions.Dy()

	c.PixelData = img
	c.Format = format
	c.shouldReloadImage = true
	c.Refresh()
}

func (c *Canvas) NewDrawing(cols, rows int) {
	c.State.SetFilePath("")
	c.Cols = cols
	c.Rows = rows

	pixelData := NewBlankImage(cols, rows, color.NRGBA{128, 128, 128, 255})
	c.LoadImage(pixelData, "png")
}
