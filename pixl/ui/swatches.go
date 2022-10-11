package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"image/color"
	"takken.io/pixl/swatch"
)

func buildSwatches(app *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		s := swatch.NewSwatch(app.State, initialColor, i, func(swatch *swatch.Swatch) *swatch.Swatch {
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}

			app.State.SwatchSelected = swatch.SwatchIndex
			app.State.BrushColor = swatch.Color

			return swatch
		})

		if i == 0 {
			s.Selected = true
			app.State.SwatchSelected = 0
			s.Refresh()
		}

		app.Swatches = append(app.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}

	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
