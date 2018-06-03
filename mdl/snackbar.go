package mdl

import (
	"fmt"

	"github.com/dtylman/gowd"
)

type Snackbar struct {
	*gowd.Element
	textElement   *gowd.Element
	buttonElement *gowd.Element
	eventHandler  *gowd.Element
	message       string
	timeout       int
	actionhandler string
	actionText    string
}

func NewSnackbar(message string, timeout int, eventlistner gowd.EventHandler) *Snackbar {
	snackbar := new(Snackbar)

	snackbar.Element = NewElement("div", "mdl-js-snackbar mdl-snackbar")

	snackbar.textElement = NewElement("div", "mdl-snackbar__text")

	snackbar.buttonElement = NewElement("button", "mdl-snackbar__action")
	snackbar.buttonElement.OnEvent(gowd.OnClick, eventlistner)

	snackbar.actionhandler = fmt.Sprintf("fire_event('onclick',document.getElementById(%s));", snackbar.buttonElement.GetID())

	snackbar.eventHandler = NewElement("script", "")
	ExecJS(fmt.Sprintf("var handler%s = function(event) {%s};", snackbar.Element.GetID(), snackbar.actionhandler))
	snackbar.Element.AddElement(snackbar.textElement)
	snackbar.Element.AddElement(snackbar.buttonElement)
	snackbar.Element.AddElement(snackbar.eventHandler)

	return snackbar
}
