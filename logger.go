package main

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

type Logger struct {
	widget.TextGrid
}

func NewLogger() *Logger {
	grid := &Logger{}
	grid.ExtendBaseWidget(grid)
	grid.ShowLineNumbers = true
	grid.ShowWhitespace = true
	return grid
}

func (t *Logger) log(msg string) {
	t.SetText(fmt.Sprintf("%s%s\n", t.Text(), msg))
}
