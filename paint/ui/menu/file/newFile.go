package file

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"takken.io/paint/canvas"
	"takken.io/paint/global"
)

func NewMenu(app *global.App, canvas *canvas.Canvas) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil || width <= 0 {
				return errors.New("must be a positive integer")
			}
			return nil
		}

		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}

		dialog.ShowForm("New image", "Create", "Cancel", formItems, func(ok bool) {
			if !ok {
				return
			}

			pixelWidth := 0
			pixelHeight := 0

			if widthEntry.Validate() != nil {
				dialog.ShowError(errors.New("invalid width"), app.Window)
			}
			if heightEntry.Validate() != nil {
				dialog.ShowError(errors.New("invalid height"), app.Window)
			}

			pixelWidth, _ = strconv.Atoi(widthEntry.Text)
			pixelHeight, _ = strconv.Atoi(heightEntry.Text)

			// Todo - use AppState to open new tab instead
			canvas.NewDrawing(pixelWidth, pixelHeight)

		}, app.Window)
	})
}
