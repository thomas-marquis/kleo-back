package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/thomas-marquis/kleo-back/desktop/app"
	apptheme "github.com/thomas-marquis/kleo-back/desktop/ui/theme"
	"github.com/thomas-marquis/kleo-back/internal/domain"
)

func makeTransactionDetailsContenr(t domain.Transaction) fyne.CanvasObject {
	parsedDate := t.Date.Format("2006/05/04")

	content := container.New(layout.NewVBoxLayout(),
		widget.NewLabel(parsedDate),
		widget.NewLabel(fmt.Sprintf("%.2f €", t.Amount)),
	)

	if len(t.Allocations) > 0 {
		allocs := make([]struct {
			UserID domain.UserId
			Ratio  float64
		}, len(t.Allocations))
		for a, r := range t.Allocations {
			allocs = append(allocs, struct {
				UserID domain.UserId
				Ratio  float64
			}{a, r})
		}
		allTable := widget.NewTable(
			func() (int, int) {
				return len(t.Allocations), 2
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("-")
			},
			func(cellId widget.TableCellID, obj fyne.CanvasObject) {
				al := allocs[cellId.Row]
				cell := obj.(*widget.Label)
				if cellId.Col == 0 {
					cell.SetText("TOOT")
				}
				if cellId.Col == 1 {
					cell.SetText(fmt.Sprintf("%d%%", int(al.Ratio*100)))
				}
				cell.Resize(fyne.NewSize(200, cell.Size().Height)) // TODO
			},
		)
		content.Add(allTable)
	}

	return content
}

func getColorFromCategory(cat *domain.Category) color.Color {
	c := apptheme.Unclassified
	if cat != nil {
		switch cat.SubCategory.MovmentType {
		case domain.BaseExpense:
			c = apptheme.Expense
		case domain.BaseIncome:
			c = apptheme.Income
		case domain.BaseSaving:
			c = apptheme.Saving
		case domain.BaseTransfer:
			c = apptheme.Transfer
		}
	}
	return c
}

func getColorFromAmount(amount float64) color.Color {
	c := apptheme.Unclassified
	if amount < 0 {
		c = apptheme.Expense
	} else if amount > 0 {
		c = apptheme.Income
	}
	return c
}

// makeBadge creates a badge with the given text, background color, and text color.
func makeBadge(text string, bgColor color.Color, txtColor color.Color) *fyne.Container {
	label := canvas.NewText(text, txtColor)

	rect := canvas.NewRectangle(bgColor)
	content := container.NewStack(rect, label)

	return container.NewPadded(content)
}

func updateBadge(b *fyne.Container, text string, bgColor color.Color, txtColor color.Color) {
	content := b.Objects[0].(*fyne.Container)
	rect := content.Objects[0].(*canvas.Rectangle)
	label := content.Objects[1].(*canvas.Text)

	rect.FillColor = bgColor

	label.Text = text
	label.Color = txtColor
}

// makePriceText create a text object with the given transaction amount.
func makePriceText(t domain.Transaction) *canvas.Text {
	pText := fmt.Sprintf("%.2f", t.Amount)
	c := getColorFromAmount(t.Amount)
	txt := canvas.NewText(pText, c)
	txt.TextStyle.Bold = true
	return txt
}

// updatePriceText update a text object with the given transaction amount.
func updatePriceText(txt *canvas.Text, t domain.Transaction) {
	pText := fmt.Sprintf("%.2f", t.Amount)
	c := getColorFromAmount(t.Amount)
	txt.Text = pText
	txt.Color = c
}

func updateCategoryBadge(b *fyne.Container, t domain.Transaction) {
	if cat := t.Category; cat != nil {
		updateBadge(b, cat.Label, getColorFromCategory(cat), apptheme.TextPrimary)
	} else {
		updateBadge(b, "Non classé", apptheme.Unclassified, apptheme.TextPrimary)
	}
}

func makeDetailButton(callback func()) *widget.Button {
	btn := widget.NewButton("", callback)
	btn.SetIcon(theme.NavigateNextIcon())
	btn.Resize(fyne.NewSize(30, 30))
	return btn
}

func updateDetailButton(btn *widget.Button, callback func()) {
	btn.OnTapped = callback
}

func makeTransactionItemRaw() *fyne.Container {
	cat := makeBadge("", apptheme.Unclassified, apptheme.TextPrimary)
	label := widget.NewLabel("")
	price := makePriceText(domain.Transaction{Amount: 100})

	content := container.NewBorder(nil, nil, label, cat)

	return container.NewBorder(
		nil,
		nil,
		price,                       // 1
		makeDetailButton(func() {}), // 2
		content,                     // 0
	)
}

func makeTransactionItem(ctx *app.AppContext, c *fyne.Container, t domain.Transaction) {
	content := c.Objects[0].(*fyne.Container)

	label := content.Objects[0].(*widget.Label)
	label.Bind(binding.BindString(&t.Label))

	cat := content.Objects[1].(*fyne.Container)
	updateCategoryBadge(cat, t)

	price := c.Objects[1].(*canvas.Text)
	updatePriceText(price, t)

	btn := c.Objects[2].(*widget.Button)
	updateDetailButton(btn, func() {
		dialogContent := makeTransactionDetailsContenr(t)
		dial := dialog.NewCustom(t.Label, "Fermer", dialogContent, ctx.W)
		dial.Resize(fyne.NewSize(700, 500))
		dial.Show()
	})
}

func GetTransactionListView(ctx *app.AppContext) *fyne.Container {
	content := widget.NewListWithData(
		ctx.VM.Transactions(),
		func() fyne.CanvasObject {
			return makeTransactionItemRaw()
		},
		func(bd binding.DataItem, item fyne.CanvasObject) {
			di, _ := bd.(binding.Untyped).Get()
			t := di.(domain.Transaction)
			c := item.(*fyne.Container)

			makeTransactionItem(ctx, c, t)
		},
	)

	searchBtn := widget.NewButton("Rechercher", func() {
		ctx.VM.LoadingStart()
		ctx.VM.SearchTransactions()
		ctx.VM.LoadingEnd()
	})

	prog := widget.NewProgressBarInfinite()
	prog.Resize(fyne.NewSize(ctx.W.Canvas().Size().Width, 100))
	progressBar := container.NewCenter(prog)
	progressBar.Hide()

	c := container.NewBorder(
		searchBtn,
		nil,
		nil,
		nil,
		progressBar,
		content,
	)

	ctx.VM.IsLoading().AddListener(binding.NewDataListener(func() {
		if val, _ := ctx.VM.IsLoading().Get(); val {
			progressBar.Show()
			content.Hide()
		} else {
			content.Show()
			progressBar.Hide()
		}
	}))

	return c
}
