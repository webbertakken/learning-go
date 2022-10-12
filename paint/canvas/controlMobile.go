package canvas

import "fyne.io/fyne/v2/driver/mobile"

// Touchable - https://developer.fyne.io/api/v2.2/driver/mobile/touchable.html

func (c *Canvas) TouchDown(event *mobile.TouchEvent) {}

func (c *Canvas) TouchUp(event *mobile.TouchEvent) {}

func (c *Canvas) TouchCancel(event *mobile.TouchEvent) {}
