package ui

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/frei-0xff/inmak_invoice_generator/internal/ui/wgt"
	"github.com/frei-0xff/inmak_invoice_generator/internal/util"
	"github.com/shopspring/decimal"
)

func createItemsList(window fyne.Window, items *[]itemData, onUpdate func()) fyne.CanvasObject {
	var con *wgt.ValidatableContaner
	var selected widget.ListItemID
	var add, edit, remove, clone, up, down *widget.Button

	list := wgt.NewList(
		func() int {
			return len(*items)
		},
		func() fyne.CanvasObject {
			return container.NewVBox(
				widget.NewLabel("template"),
				container.NewGridWithColumns(4,
					widget.NewLabelWithStyle("tmpl", fyne.TextAlignTrailing, fyne.TextStyle{}),
					widget.NewLabelWithStyle("tmpl", fyne.TextAlignTrailing, fyne.TextStyle{}),
					widget.NewLabelWithStyle("tmpl", fyne.TextAlignTrailing, fyne.TextStyle{}),
					widget.NewLabelWithStyle("tmpl", fyne.TextAlignTrailing, fyne.TextStyle{}),
				))
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			row1 := o.(*fyne.Container).Objects[0].(*widget.Label)
			row1.SetText((*items)[i].name)

			row2 := o.(*fyne.Container).Objects[1].(*fyne.Container)
			row2.Objects[0].(*widget.Label).SetText((*items)[i].unit)
			row2.Objects[1].(*widget.Label).SetText((*items)[i].quantity)
			row2.Objects[2].(*widget.Label).SetText((*items)[i].price)

			s := decimal.NewFromInt(0)
			if q, err := util.ParseDecimal((*items)[i].quantity); err == nil {
				if p, err := util.ParseDecimal((*items)[i].price); err == nil {
					s = q.Mul(p)
				}
			}
			row2.Objects[3].(*widget.Label).SetText(util.FormatDecimal(s))
		})
	list.OnSelected = func(id widget.ListItemID) {
		selected = id
		edit.Enable()
		remove.Enable()
		clone.Enable()
		if id > 0 {
			up.Enable()
		} else {
			up.Disable()
		}
		if id < len(*items)-1 {
			down.Enable()
		} else {
			down.Disable()
		}
	}

	add = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		newItem := itemData{}
		openEditor(&newItem, true, func() {
			*items = append(*items, newItem)
			list.Refresh()
			con.Validate()
			onUpdate()
			if selected < len(*items) {
				list.UnselectAll()
				list.Select(selected)
			}
		})
	})

	edit = widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {
		if len(*items) > 0 && selected < len(*items) {
			openEditor(&(*items)[selected], false, func() {
				list.Refresh()
				onUpdate()
			})
		}
	})
	edit.Disable()

	remove = widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), func() {
		if len(*items) > 0 && selected < len(*items) {
			d := dialog.NewConfirm("Удалить услугу", "Вы действильно хотите удалить эту услугу?", func(confirm bool) {
				if confirm {
					*items = append((*items)[0:selected], (*items)[selected+1:]...)
					onUpdate()
					list.UnselectAll()
					if selected < len(*items) {
						list.Select(selected)
					} else {
						selected = int(^uint(0) >> 1)
						edit.Disable()
						remove.Disable()
						clone.Disable()
						up.Disable()
						down.Disable()
					}
				}
				con.Validate()
			}, window)
			d.SetConfirmText("Да")
			d.SetDismissText("Нет")
			d.Show()
		}
	})
	remove.Disable()

	clone = widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		if len(*items) > 0 && selected < len(*items) {
			newItem := (*items)[selected]
			*items = append((*items)[:selected+1], (*items)[selected:]...)
			(*items)[selected+1] = newItem
			list.Refresh()
			onUpdate()
			list.UnselectAll()
			list.Select(selected)
		}
	})
	clone.Disable()

	up = widget.NewButtonWithIcon("", theme.MoveUpIcon(), func() {
		if len(*items) > 0 && selected < len(*items) && selected > 0 {
			(*items)[selected], (*items)[selected-1] = (*items)[selected-1], (*items)[selected]
			list.Refresh()
			onUpdate()
			list.Select(selected - 1)
		}
	})
	up.Disable()

	down = widget.NewButtonWithIcon("", theme.MoveDownIcon(), func() {
		if len(*items) > 0 && selected < len(*items) && selected < len(*items)-1 {
			(*items)[selected], (*items)[selected+1] = (*items)[selected+1], (*items)[selected]
			list.Refresh()
			onUpdate()
			list.Select(selected + 1)
		}
	})
	down.Disable()

	con = wgt.NewValidatableContaner(
		container.NewVBox(container.NewHBox(up, down, layout.NewSpacer(), add, edit, clone, remove), list),
		func() error {
			if len(*items) < 1 {
				return errors.New("")
			}
			return nil
		})
	return con
}
