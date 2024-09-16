package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/thomas-marquis/kleo-back/desktop/app"
)

func MakeMenubar(ctx *app.AppContext) *fyne.Container {
	return container.NewVBox(
		widget.NewButton("Transactions", func() { ctx.Navigate("transactions") }),
		widget.NewButton("Importation", func() { ctx.Navigate("importation") }),
	)
}
