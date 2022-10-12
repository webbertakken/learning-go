package file

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"image"
	"takken.io/paint/canvas"
	"takken.io/paint/core"
	"takken.io/paint/global"
)

func OpenMenu(app *global.App, canvas *canvas.Canvas) *fyne.MenuItem {
	return fyne.NewMenuItem("Open...", func() {
		dialog.ShowFileOpen(func(blob fyne.URIReadCloser, err error) {
			if blob == nil {
				return
			}

			image, format, err := image.Decode(blob)
			if err != nil {
				dialog.ShowError(err, app.Window)
				return
			}

			canvas.LoadImage(image, format)
			canvas.State.SetFilePath(blob.URI().Path())
			imgColors := core.GetImageColors(image)

			i := 0
			for c := range imgColors {
				if i >= len(app.Swatches) {
					break
				}

				app.Swatches[i].SetColor(c)

				i += 1
			}

		}, app.Window)
	})
}
