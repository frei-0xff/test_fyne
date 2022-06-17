package wgt

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const ShowItems = 4

type baseList = widget.List

type List struct {
	baseList
}

func (cl *List) MinSize() fyne.Size {
	size := cl.baseList.MinSize()
	n := cl.Length()
	if n > ShowItems {
		n = ShowItems
	}
	size.Height = size.Height*float32(n) + theme.SeparatorThicknessSize()*float32(n)
	return size
}

func NewList(length func() int, createItem func() fyne.CanvasObject, updateItem func(widget.ListItemID, fyne.CanvasObject)) *List {
	list := &List{baseList: widget.List{BaseWidget: widget.BaseWidget{}, Length: length, CreateItem: createItem, UpdateItem: updateItem}}
	list.ExtendBaseWidget(list)
	return list
}
