package main

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"image/color"
	"takken.io/paint/appTypes"
	canvas2 "takken.io/paint/canvas"
	"takken.io/paint/global"
	"takken.io/paint/swatch"
	"takken.io/paint/ui"
)

func main() {
	app := fyneApp.New()
	window := app.NewWindow("paint")

	state := appTypes.AppState{
		BrushColor:     color.NRGBA{R: 255, G: 255, B: 255, A: 255},
		SwatchSelected: 0,
	}

	canvasConfig := appTypes.CanvasConfig{
		DrawingArea:  fyne.NewSize(800, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		Rows:         10,
		Cols:         10,
		PixelRatio:   30,
	}

	canvas := canvas2.NewCanvas(&state, canvasConfig)

	appUI := global.App{
		Canvas:   canvas,
		Window:   window,
		State:    &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appUI)

	appUI.Window.ShowAndRun()
}
