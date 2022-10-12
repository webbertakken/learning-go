package brush

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"image/color"
	"takken.io/paint/appTypes"
)

func Cursor(config appTypes.CanvasConfig, brushType appTypes.BrushType, event *desktop.MouseEvent, x int, y int) []fyne.CanvasObject {
	var objects []fyne.CanvasObject

	switch brushType {
	case Pixel:
		pixelRatio := float32(config.PixelRatio)
		xOrigin := float32(x)*pixelRatio + config.CanvasOffset.X
		yOrigin := float32(y)*pixelRatio + config.CanvasOffset.Y

		cursorColor := color.NRGBA{R: 80, G: 80, B: 80, A: 255}

		left := canvas.NewLine(cursorColor)
		left.StrokeWidth = 3
		left.Position1 = fyne.NewPos(xOrigin, yOrigin)
		left.Position2 = fyne.NewPos(xOrigin, yOrigin+pixelRatio)

		top := canvas.NewLine(cursorColor)
		top.StrokeWidth = 3
		top.Position1 = fyne.NewPos(xOrigin, yOrigin)
		top.Position2 = fyne.NewPos(xOrigin+pixelRatio, yOrigin)

		right := canvas.NewLine(cursorColor)
		right.StrokeWidth = 3
		right.Position1 = fyne.NewPos(xOrigin+pixelRatio, yOrigin)
		right.Position2 = fyne.NewPos(xOrigin+pixelRatio, yOrigin+pixelRatio)

		bottom := canvas.NewLine(cursorColor)
		bottom.StrokeWidth = 3
		bottom.Position1 = fyne.NewPos(xOrigin, yOrigin+pixelRatio)
		bottom.Position2 = fyne.NewPos(xOrigin+pixelRatio, yOrigin+pixelRatio)

		objects = append(objects, left, top, right, bottom)
	}

	return objects
}
