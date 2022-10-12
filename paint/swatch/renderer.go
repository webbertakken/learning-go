package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

type Renderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

func (r *Renderer) MinSize() fyne.Size {
	return r.square.MinSize()
}

func (r *Renderer) Layout(size fyne.Size) {
	r.objects[0].Resize(size)
}

func (r *Renderer) Refresh() {
	r.Layout(fyne.NewSize(20, 20))
	r.square.FillColor = r.parent.Color
	if r.parent.Selected {
		r.square.StrokeWidth = 3
		r.square.StrokeColor = color.NRGBA{255, 255, 255, 255}
		r.objects[0] = &r.square
	} else {
		r.square.StrokeWidth = 0
		r.objects[0] = &r.square
	}
	canvas.Refresh(r.parent)
}

func (r *Renderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *Renderer) Destroy() {

}
