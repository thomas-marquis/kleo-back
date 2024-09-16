package main

import (
	"fyne.io/fyne/v2"
	"github.com/thomas-marquis/kleo-back/desktop/app"
	"github.com/thomas-marquis/kleo-back/desktop/ui"
)

func main() {
	views := make(map[string]func(*app.AppContext) *fyne.Container)
	views["transactions"] = ui.GetTransactionListView
	app := app.New(views, "transactions")
	if err := app.Start(); err != nil {
		panic(err)
	}
}
