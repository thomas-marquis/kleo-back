package app

import (
	"fyne.io/fyne/v2"
	"github.com/thomas-marquis/kleo-back/desktop/ui/viewmodel"
)

type AppContext struct {
	VM *viewmodel.ViewModel
	W  fyne.Window
}

func NewContext(vm *viewmodel.ViewModel, w fyne.Window) *AppContext {
	return &AppContext{vm, w}
}
