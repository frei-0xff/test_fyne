package wgt

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ValidatableContaner struct {
	widget.BaseWidget
	container           *fyne.Container
	validator           func() error
	onValidationChanged func(error)
}

func NewValidatableContaner(container *fyne.Container, validator func() error) *ValidatableContaner {
	w := &ValidatableContaner{
		container: container,
		validator: validator,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (vc *ValidatableContaner) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(vc.container)
}

var _ fyne.Validatable = (*ValidatableContaner)(nil)

func (vc *ValidatableContaner) Validate() error {
	err := vc.validator()
	if vc.onValidationChanged != nil {
		vc.onValidationChanged(err)
	}
	return err
}

func (vc *ValidatableContaner) SetOnValidationChanged(callback func(error)) {
	vc.onValidationChanged = callback
	vc.Validate()
}
