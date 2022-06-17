package ui

import (
	"errors"
	"regexp"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/frei-0xff/inmak_invoice_generator/internal/util"
	"github.com/shopspring/decimal"
)

func CreateMaster(window fyne.Window) fyne.CanvasObject {
	window.SetMaster()
	header := widget.NewCard("Генератор счетов", "Создать новый счет + акт", nil)
	return container.NewVScroll(container.NewBorder(header, nil, nil, nil, header, createMainForm(window)))
}

type mainFormData struct {
	date   binding.String
	number binding.String
	payer  binding.String
	inn    binding.String
	items  []itemData
}

func newMainFormData() *mainFormData {
	d := mainFormData{
		date:   binding.NewString(),
		number: binding.NewString(),
		payer:  binding.NewString(),
		inn:    binding.NewString(),
		items:  []itemData{},
	}
	d.date.Set(time.Now().Format("02.01.2006"))

	d.items = append(d.items, itemData{
		name:     "Продовження реєстрації доменного імені: forpost.com.ua",
		unit:     "12 міс.",
		quantity: "1",
		price:    "321,00",
	})
	d.items = append(d.items, itemData{
		name:     "Хостинг сайта: forpost.com.ua",
		unit:     "12 міс.",
		quantity: "1",
		price:    "400,00",
	})

	return &d
}

type itemData struct {
	name     string
	unit     string
	quantity string
	price    string
}

func createMainForm(window fyne.Window) fyne.CanvasObject {
	d := newMainFormData()

	form := widget.NewForm()
	form.OnSubmit = func() {
	}
	form.SubmitText = "Сгенерировать"

	// Дата
	date := widget.NewEntryWithData(d.date)
	date.Validator = func(s string) error {
		if s == "" {
			return errors.New("Необходимо ввести дату")
		}
		if ok, _ := regexp.MatchString(`\d\d\.\d\d\.\d\d\d\d`, s); !ok {
			return errors.New("Дата должна быть в формате дд.мм.гггг")
		}
		if _, err := time.Parse("02.01.2006", s); err != nil {
			return errors.New("Неверная дата")
		}
		return nil
	}
	date.OnChanged = func(s string) {
		df := "00000000"
		if d, err := time.Parse("02.01.2006", s); err == nil {
			df = d.Format("20060102")
		}
		d.number.Set("1-" + df + "-1")
	}
	date.SetText(util.GetBoundString(d.date))
	form.Append("Дата", date)

	// Номер счета
	number := widget.NewEntryWithData(d.number)
	number.Validator = func(s string) error {
		if s == "" {
			return errors.New("Необходимо ввести номер счета")
		}
		return nil
	}
	number.SetText(util.GetBoundString(d.number))
	form.Append("Номер счета", number)

	// Плательщик
	payer := widget.NewEntryWithData(d.payer)
	form.Append("Плательщик", payer)

	// ИНН
	inn := widget.NewEntryWithData(d.inn)
	inn.Validator = func(s string) error {
		if ok, _ := regexp.MatchString(`^\d*$`, s); !ok {
			return errors.New("ИНН должен состоять только из цифр")
		}
		return nil
	}
	form.Append("ИНН", inn)

	// Итого
	total := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{})
	onUpdate := func() {
		if len(d.items) == 0 {
			total.SetText("")
			return
		}
		t := decimal.NewFromInt(0)
		for _, item := range d.items {
			price, err := util.ParseDecimal(item.price)
			if err != nil {
				continue
			}
			quantity, err := util.ParseDecimal(item.quantity)
			if err != nil {
				continue
			}
			t = t.Add(price.Mul(quantity))
		}
		total.SetText(util.FormatDecimal(t))
	}

	// Услуги
	form.Append("Услуги", createItemsList(window, &d.items, onUpdate))

	form.Append("Итого", total)

	return container.NewVBox(widget.NewLabel(""), form)
}
