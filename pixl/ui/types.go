package ui

import (
	"fyne.io/fyne/v2"
	"takken.io/pixl/apptype"
	"takken.io/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
