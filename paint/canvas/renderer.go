package canvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Renderer struct {
	Canvas *Canvas
	Image  *canvas.Image
	Border []canvas.Line
	Cursor []fyne.CanvasObject
}

func (r *Renderer) SetCursor(objects []fyne.CanvasObject) {
	r.Cursor = objects
}

func (r *Renderer) MinSize() fyne.Size {
	return r.Canvas.DrawingArea
}

func (r *Renderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(r.Border); i++ {
		objects = append(objects, &r.Border[i])
	}
	objects = append(objects, r.Image)
	objects = append(objects, r.Cursor...)

	return objects
}

func (r *Renderer) Destroy() {}

func (r *Renderer) Refresh() {
	if r.Canvas.shouldReloadImage {
		r.Image = canvas.NewImageFromImage(r.Canvas.PixelData)
		r.Image.ScaleMode = canvas.ImageScalePixels
		r.Image.FillMode = canvas.ImageFillContain
		r.Canvas.shouldReloadImage = false
	}
	r.Layout(r.Canvas.Size())
	canvas.Refresh(r.Image)
}

func (r *Renderer) Layout(size fyne.Size) {
	r.LayoutCanvas(size)
	r.LayoutBorder(size)
}

func (r *Renderer) LayoutCanvas(size fyne.Size) {
	pxSize := r.Canvas.PixelRatio

	imgWidth := r.Canvas.Cols * pxSize
	imgHeight := r.Canvas.Rows * pxSize

	r.Image.Move(fyne.NewPos(r.Canvas.CanvasOffset.X, r.Canvas.CanvasOffset.Y))
	r.Image.Resize(fyne.NewSize(float32(imgWidth), float32(imgHeight)))
}

func (r *Renderer) LayoutBorder(size fyne.Size) {
	offset := r.Canvas.CanvasOffset
	imgWidth := r.Image.Size().Width
	imgHeight := r.Image.Size().Height

	left := &r.Border[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	top := &r.Border[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y)

	right := &r.Border[2]
	right.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)

	bottom := &r.Border[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)
}
