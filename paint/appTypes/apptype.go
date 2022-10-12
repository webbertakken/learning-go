package appTypes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"image/color"
)

type BrushType int

type BrushableState struct {
	IsBrushing bool
}

type FileState struct {
	FilePath string
}

type CanvasConfig struct {
	DrawingArea  fyne.Size
	CanvasOffset fyne.Position
	Rows         int
	Cols         int
	PixelRatio   int
}

type AppState struct {
	BrushColor     color.Color
	BrushType      BrushType
	SwatchSelected int
	//Canvasses      []*CanvasState
}

type CanvasState struct {
	BrushableState
	File FileState
}

//func (appState *AppState) NewTab(canvasState *CanvasState) {
//	appState.Canvasses = append(appState.Canvasses, canvasState)
//}

func (state *CanvasState) SetFilePath(filePath string) {
	state.File.FilePath = filePath
}

type Brushable interface {
	SetColor(c color.Color, x, y int)
	MouseToCanvasXY(event *desktop.MouseEvent) (*int, *int)
}
