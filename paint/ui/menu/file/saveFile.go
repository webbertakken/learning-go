package file

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"image/png"
	"os"
	"takken.io/paint/canvas"
	"takken.io/paint/global"
)

func saveFileDialog(app *global.App, canvas *canvas.Canvas, window fyne.Window) {
	dialog.ShowFileSave(func(blob fyne.URIWriteCloser, err error) {
		if blob == nil {
			return
		} else {
			err := png.Encode(blob, canvas.PixelData)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			canvas.State.SetFilePath(blob.URI().Path())

		}
	}, window)
}

func save(filePath string, canvas *canvas.Canvas, window fyne.Window) {
	fileHandle, err := os.Create(filePath)

	defer func() {
		err := fileHandle.Close()
		if err != nil {
			dialog.ShowError(err, window)
		}
	}()

	if err != nil {
		dialog.ShowError(err, window)
	}

	err = png.Encode(fileHandle, canvas.PixelData)
	if err != nil {
		dialog.ShowError(err, window)
		return
	}
}

func SaveAsMenu(app *global.App, canvas *canvas.Canvas) *fyne.MenuItem {
	return fyne.NewMenuItem("Save as...", func() {
		saveFileDialog(app, canvas, app.Window)
	})
}

func SaveMenu(app *global.App, canvas *canvas.Canvas) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		filePath := canvas.State.File.FilePath
		if filePath != "" {
			save(filePath, canvas, app.Window)
		} else {
			saveFileDialog(app, canvas, app.Window)
		}
	})
}
