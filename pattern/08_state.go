package pattern

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type (
	WritingState interface {
		Write(text string)
	}

	UpperCaseState struct{}
	DefaultState   struct{}

	TextEditor struct {
		WritingState
	}
)

func (state UpperCaseState) Write(text string) {
	fmt.Println(strings.ToUpper(text))
}

func (state DefaultState) Write(text string) {
	fmt.Println(text)
}

func (textEditor *TextEditor) SetState(state WritingState) {
	textEditor.WritingState = state
}

func (textEditor *TextEditor) Input(text string) {
	textEditor.WritingState.Write(text)
}

func main8() {
	editor := &TextEditor{DefaultState{}}
	editor.Input("ОбыЧнЫй")
	editor.SetState(UpperCaseState{})
	editor.Input("аппер КеЙс")
}
