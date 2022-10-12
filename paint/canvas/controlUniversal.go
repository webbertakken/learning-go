package canvas

import "fyne.io/fyne/v2"

// Scrollable - https://developer.fyne.io/api/v2.2/scrollable.html

func (c *Canvas) Scrolled(event *fyne.ScrollEvent) {
  c.scale(int(event.Scrolled.DY * scrollSpeed))
  c.Refresh()
}
