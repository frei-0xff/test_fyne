// Package assets contains bundled static resources.
package assets

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed icon/icon.png
var icon []byte

var Icon = &fyne.StaticResource{
	StaticContent: icon,
}