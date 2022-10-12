package global

import (
	"fyne.io/fyne/v2"
	"takken.io/paint/appTypes"
	"takken.io/paint/canvas"
	"takken.io/paint/swatch"
)

type App struct {
	Canvas   *canvas.Canvas
	Window   fyne.Window
	State    *appTypes.AppState
	Swatches []*swatch.Swatch
}
