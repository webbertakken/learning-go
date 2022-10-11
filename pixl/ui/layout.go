package ui

func Setup(app *AppInit) {
	swatchesContainer := buildSwatches(app)

	app.PixlWindow.SetContent(swatchesContainer)
}
