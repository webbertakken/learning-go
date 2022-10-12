package ui

import (
	"fyne.io/fyne/v2/container"
	"takken.io/paint/global"
	"takken.io/paint/ui/menu"
)

func Setup(app *global.App) {
	mainMenu := menu.BuildNewMenu(app, app.Canvas)
	app.Window.SetMainMenu(mainMenu)

	swatchesContainer := buildSwatches(app)
	colorPicker := SetupColorPicker(app)
	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker, app.Canvas)
	app.Window.SetContent(appLayout)
}
