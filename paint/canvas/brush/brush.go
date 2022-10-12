package brush

import (
	"fyne.io/fyne/v2/driver/desktop"
	"takken.io/paint/appTypes"
)

const (
	Pixel = iota
)

func TryBrush(appState *appTypes.AppState, brushable appTypes.Brushable, event *desktop.MouseEvent) bool {
	switch {
	case appState.BrushType == Pixel:
		return TryPaintPixel(appState, brushable, event)
	default:
		return false
	}
}

func TryPaintPixel(appState *appTypes.AppState, brushable appTypes.Brushable, event *desktop.MouseEvent) bool {
	x, y := brushable.MouseToCanvasXY(event)
	if x == nil || y == nil {
		return false
	}

	if event.Button == desktop.MouseButtonPrimary {
		brushable.SetColor(appState.BrushColor, *x, *y)
		return true
	}

	return false
}
