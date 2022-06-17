package util

import (
	"strings"

	"fyne.io/fyne/v2/data/binding"
	"github.com/shopspring/decimal"
)

func GetBoundString(str binding.String) string {
	if s, err := str.Get(); err == nil {
		return s
	}
	return ""
}

func ParseDecimal(str string) (decimal.Decimal, error) {
	str = strings.ReplaceAll(str, ",", ".")
	return decimal.NewFromString(str)
}

func FormatDecimal(dec decimal.Decimal) string {
	return strings.ReplaceAll(dec.StringFixedBank(2), ".", ",")
}
