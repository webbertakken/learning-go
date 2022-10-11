package apptype

import (
	"fyne.io/fyne/v2"
	"image/color"
)

type BrushType int

type CanvasConfig struct {
	DrawingArea  fyne.Size
	CanvasOffset fyne.Position
	PxRows       int
	PxCols       int
	PxSize       int
}

type State struct {
	BrushColor     color.Color
	BrushType      BrushType
	SwatchSelected int
	FilePath       string
}

func (state *State) SetFilePath(filePath string) {
	state.FilePath = filePath
}
