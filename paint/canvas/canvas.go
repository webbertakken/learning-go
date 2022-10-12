package canvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	image "image"
	"image/color"
	"takken.io/paint/appTypes"
)

type CanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type Canvas struct {
	widget.BaseWidget
	appTypes.CanvasConfig
	renderer          *Renderer
	PixelData         image.Image
	mouseState        CanvasMouseState
	appState          *appTypes.AppState
	State             *appTypes.CanvasState
	shouldReloadImage bool
	Format            string
}

func (c *Canvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(c.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &Renderer{
		Canvas: c,
		Image:  canvasImage,
		Border: canvasBorder,
	}
	c.renderer = renderer

	return renderer
}

func (c *Canvas) Bounds() image.Rectangle {
	x0 := int(c.CanvasOffset.X)
	y0 := int(c.CanvasOffset.Y)
	x1 := int(c.Cols*c.PixelRatio + int(c.CanvasOffset.X))
	y1 := int(c.Rows*c.PixelRatio + int(c.CanvasOffset.Y))

	return image.Rectangle{Min: image.Point{x0, y0}, Max: image.Point{x1, y1}}
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X < float32(bounds.Min.X) || pos.X > float32(bounds.Max.X) {
		return false
	}

	if pos.Y < float32(bounds.Min.Y) || pos.Y > float32(bounds.Max.Y) {
		return false
	}

	return true
}

func NewBlankImage(cols, rows int, color color.Color) image.Image {
	img := image.NewNRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{cols, rows}})
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, color)
		}
	}

	return img
}

func NewCanvas(state *appTypes.AppState, config appTypes.CanvasConfig) *Canvas {
	canvas := &Canvas{
		CanvasConfig: config,
		appState:     state,
		State:        &appTypes.CanvasState{},
	}

	canvas.PixelData = NewBlankImage(canvas.Cols, canvas.Rows, color.NRGBA{128, 128, 128, 255})
	canvas.ExtendBaseWidget(canvas)

	return canvas
}

func (c *Canvas) TryPan(mouseDragStart *fyne.PointEvent, event *desktop.MouseEvent) {
	if event.Button == desktop.MouseButtonTertiary && mouseDragStart != nil {
		c.Pan(*mouseDragStart, event.PointEvent)
	}
}

func (c *Canvas) SetColor(color color.Color, x, y int) {
	if nrgba, ok := c.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, color)
	}

	if rgba, ok := c.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, color)
	}
}

func (c *Canvas) MouseToCanvasXY(event *desktop.MouseEvent) (*int, *int) {
	bounds := c.Bounds()
	if !InBounds(event.Position, bounds) {
		return nil, nil
	}

	x := int((event.Position.X - c.CanvasOffset.X) / float32(c.PixelRatio))
	y := int((event.Position.Y - c.CanvasOffset.Y) / float32(c.PixelRatio))

	return &x, &y
}
