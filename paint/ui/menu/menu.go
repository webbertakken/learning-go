package menu

import (
	"fyne.io/fyne/v2"
	"takken.io/paint/canvas"
	"takken.io/paint/global"
	"takken.io/paint/ui/menu/file"
)

func BuildNewMenu(app *global.App, canvas *canvas.Canvas) *fyne.MainMenu {
	return &fyne.MainMenu{Items: []*fyne.Menu{file.BuildMenu(app, canvas)}}
}
