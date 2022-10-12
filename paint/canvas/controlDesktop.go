package canvas

import (
	"fyne.io/fyne/v2/driver/desktop"
	"takken.io/paint/canvas/brush"
)

const scrollSpeed = 2

// Mouseable - https://developer.fyne.io/api/v2.2/driver/desktop/mouseable.html

func (c *Canvas) MouseDown(event *desktop.MouseEvent) {
	c.State.IsBrushing = true
	brush.TryBrush(c.appState, c, event)
}

func (c *Canvas) MouseUp(event *desktop.MouseEvent) {
	c.State.IsBrushing = false
}

// Hoverable - https://developer.fyne.io/api/v2.2/driver/desktop/hoverable.html

func (c *Canvas) MouseIn(event *desktop.MouseEvent) {}

func (c *Canvas) MouseOut() {}

func (c *Canvas) MouseMoved(event *desktop.MouseEvent) {
	defer func() {
		c.Refresh()
		c.mouseState.previousCoord = &event.PointEvent
	}()

	// Panning
	c.TryPan(c.mouseState.previousCoord, event)

	x, y := c.MouseToCanvasXY(event)
	if x == nil || y == nil {
		c.renderer.SetCursor(nil)

		// Further interactions only need to be done inside the canvas
		return
	}

	// Show cursor
	cursor := brush.Cursor(c.CanvasConfig, c.appState.BrushType, event, *x, *y)
	c.renderer.SetCursor(cursor)

	// Continue brushing
	if c.State.IsBrushing {
		brush.TryBrush(c.appState, c, event)
	}

}
