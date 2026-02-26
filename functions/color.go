package functions

import (
	"io"
)

type color string

const (
	Reset color = "\x1b[0m"
	Green color = "\x1b[32m"
	Blue  color = "\x1b[34m"
)

func (c color) ColorPrint(w io.Writer, s string) {
	io.WriteString(w, string(c)+s+string(Reset))
}
