package assets

import (
	_ "embed"
)

//go:embed invoice.docx
var InvoiceTemplate []byte

//go:embed act.docx
var ActTemplate []byte
