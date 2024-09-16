package app

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/thomas-marquis/kleo-back/desktop/ui/viewmodel"
)

type kleoApp struct {
	ctx       *AppContext
	veiws     map[string]func(*AppContext) *fyne.Container
	startView string
}

func New(views map[string]func(*AppContext) *fyne.Container, startView string) *kleoApp {
	a := app.New()
	w := a.NewWindow("Kl€o")

	vm := viewmodel.New()
	ctx := NewContext(vm, w)
	return &kleoApp{ctx, views, startView}
}

func (a *kleoApp) Start() error {
	view, ok := a.veiws[a.startView]
	if !ok {
		return errors.New("specified start view does not exist")
	}
	a.ctx.W.SetContent(view(a.ctx))
	a.ctx.W.Resize(fyne.NewSize(800, 600))
	a.ctx.W.ShowAndRun()
	return nil
}
