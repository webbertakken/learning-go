package file

import (
  "fyne.io/fyne/v2"
  "takken.io/paint/canvas"
  "takken.io/paint/global"
)

func BuildMenu(app *global.App, canvas *canvas.Canvas) *fyne.Menu {
  return &fyne.Menu{
    Label: "File",
    Items: []*fyne.MenuItem{
      NewMenu(app, canvas),
      OpenMenu(app, canvas),
      SaveMenu(app, canvas),
      SaveAsMenu(app, canvas),
    },
  }
}
