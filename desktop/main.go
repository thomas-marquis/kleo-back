package main

import (
	"fyne.io/fyne/v2"
	"github.com/thomas-marquis/kleo-back/desktop/app"
	"github.com/thomas-marquis/kleo-back/desktop/ui/views"
)

func main() {
	v := make(map[string]func(*app.AppContext) *fyne.Container)
	v["transactions"] = views.GetTransactionListView
	v["importation"] = views.GetImportationView

	app := app.New(v, "transactions")
	if err := app.Start(); err != nil {
		panic(err)
	}
}
