package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/thomas-marquis/kleo-back/desktop/app"
	"github.com/thomas-marquis/kleo-back/desktop/ui/components"
)

func GetImportationView(ctx *app.AppContext) *fyne.Container {
	content := container.NewCenter(widget.NewLabel("Hello"))

	return container.NewBorder(
		nil, nil,
		components.MakeMenubar(ctx),
		nil,
		content,
	)
}
