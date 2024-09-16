package app

import (
	"fyne.io/fyne/v2"
	"github.com/thomas-marquis/kleo-back/desktop/ui/components/navigation"
	"github.com/thomas-marquis/kleo-back/desktop/ui/viewmodel"
)

type AppContext struct {
	VM *viewmodel.ViewModel
	W  fyne.Window

	currentRoute navigation.Route
	views        map[navigation.Route]func(*AppContext) *fyne.Container
}

func NewContext(vm *viewmodel.ViewModel, w fyne.Window, initialRoute navigation.Route, views map[navigation.Route]func(*AppContext) *fyne.Container) *AppContext {
	return &AppContext{vm, w, initialRoute, views}
}

func (ctx *AppContext) Navigate(route navigation.Route) {
	view, ok := ctx.views[route]
	if !ok {
		return
	}
	ctx.currentRoute = route
	ctx.W.SetContent(view(ctx))
}

func (ctx *AppContext) CurrentRoute() navigation.Route {
	return ctx.currentRoute
}
