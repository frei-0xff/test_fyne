package ui

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/frei-0xff/inmak_invoice_generator/internal/util"
)

func openEditor(item *itemData, isNew bool, onSave func()) {
	title, subtitle := "", ""
	if isNew {
		title, subtitle = "Добавить услугу", "Новая услуга"
	} else {
		title, subtitle = "Изменить услугу", "Редактирование"
	}
	app := fyne.CurrentApp()
	w := app.NewWindow(title)

	header := widget.NewCard(title, subtitle, nil)
	w.SetContent(container.NewVScroll(container.NewBorder(header, nil, nil, nil, header, createEditorForm(w, item, isNew, onSave))))

	w.Resize(fyne.NewSize(1280, 500))
	w.Show()
}

func createEditorForm(window fyne.Window, item *itemData, isNew bool, onSave func()) fyne.CanvasObject {
	form := widget.NewForm()

	// Название
	name := widget.NewEntry()
	name.SetText(item.name)
	name.Validator = func(s string) error {
		if s == "" {
			return errors.New("Необходимо ввести название")
		}
		return nil
	}

	// Единицы измерения
	unit := widget.NewEntry()
	unit.SetText(item.unit)
	unit.Validator = func(s string) error {
		if s == "" {
			return errors.New("Необходимо ввести единицы измерения")
		}
		return nil
	}

	// Количество
	quantity := widget.NewEntry()
	quantity.SetText(item.quantity)
	quantity.Validator = func(s string) error {
		if s == "" {
			return errors.New("Необходимо ввести количество")
		}
		if _, err := util.ParseDecimal(s); err != nil {
			return errors.New("Количество должно быть числом")
		}
		return nil
	}

	// Цена
	price := widget.NewEntry()
	price.SetText(item.price)
	price.Validator = func(s string) error {
		if s == "" {
			return errors.New("Необходимо ввести цену")
		}
		if _, err := util.ParseDecimal(s); err != nil {
			return errors.New("Цена должна быть числом")
		}
		return nil
	}

	form.OnSubmit = func() {
		item.name = name.Text
		item.unit = unit.Text
		item.quantity = quantity.Text
		if dec, err := util.ParseDecimal(price.Text); err == nil {
			item.price = util.FormatDecimal(dec)
		} else {
			item.price = price.Text
		}
		onSave()
		window.Close()
	}
	if isNew {
		form.SubmitText = "Добавить"
	} else {
		form.SubmitText = "Изменить"
	}
	form.OnCancel = func() {
		window.Close()
	}
	form.CancelText = "Отмена"

	form.Append("Название", name)
	form.Append("Ед. изм.", unit)
	form.Append("Количество", quantity)
	form.Append("Цена", price)

	return container.NewVBox(widget.NewLabel(""), form)
}
