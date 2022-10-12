package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"takken.io/paint/appTypes"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch) *Swatch
}

func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	objects := []fyne.CanvasObject{square}
	return &Renderer{
		square:  *square,
		objects: objects,
		parent:  s,
	}
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func NewSwatch(state *appTypes.AppState, color color.Color, swatchIndex int, clickHandler func(s *Swatch) *Swatch) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		clickHandler: clickHandler,
		SwatchIndex:  swatchIndex,
	}

	swatch.ExtendBaseWidget(swatch)

	return swatch
}
