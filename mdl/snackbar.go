package mdl

import (
	"fmt"

	"github.com/dtylman/gowd"
)

type Snackbar struct {
	*gowd.Element
	textElement   *gowd.Element
	buttonElement *gowd.Element
	message       string
	timeout       int
	actionText    string
}

func NewSnackbar(message string, timeout int, actionText string, eventlistner gowd.EventHandler) *Snackbar {
	snackbar := new(Snackbar)

	snackbar.Element = NewElement("div", "mdl-js-snackbar mdl-snackbar")

	snackbar.textElement = NewElement("div", "mdl-snackbar__text")

	snackbar.buttonElement = NewElement("button", "mdl-snackbar__action")
	snackbar.buttonElement.OnEvent(gowd.OnClick, eventlistner)

	snackbar.Element.AddElement(snackbar.textElement)
	snackbar.Element.AddElement(snackbar.buttonElement)

	snackbar.message = message
	snackbar.timeout = timeout
	snackbar.actionText = actionText

	return snackbar
}

func (snackbar *Snackbar) Open() {
	ExecJS(fmt.Sprintf("document.querySelector('#_div4').MaterialSnackbar.showSnackbar({message: '%s',timeout: %d,actionHandler: function(event) {},actionText: '%s'});", snackbar.message, snackbar.timeout, snackbar.actionText))

}
