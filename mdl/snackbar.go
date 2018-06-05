package mdl

import (
	"fmt"

	"github.com/jaicewizard/gowd"
)

type Snackbar struct {
	*gowd.Element
	textElement   *gowd.Element
	buttonElement *gowd.Element
	message       string
	timeout       int
	actionText    string
}

type Toast struct {
	*gowd.Element
	textElement *gowd.Element
	message     string
	timeout     int
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

func NewToast(message string, timeout int) *Toast {
	toast := new(Toast)

	toast.Element = NewElement("div", "mdl-js-snackbar mdl-snackbar")

	toast.textElement = NewElement("div", "mdl-snackbar__text")

	toast.Element.AddElement(toast.textElement)

	toast.message = message
	toast.timeout = timeout

	return toast
}

func (snackbar *Snackbar) Open() {
	gowd.ExecJS(fmt.Sprintf("document.querySelector('#%s').MaterialSnackbar.showSnackbar({message: '%s',timeout: %d,actionHandler: function(event) {},actionText: '%s'});", snackbar.Element.GetID(), snackbar.message, snackbar.timeout, snackbar.actionText))
}

func (toast *Toast) Open() {
	gowd.ExecJS(fmt.Sprintf("document.querySelector('"+toast.Element.GetID()+"').MaterialSnackbar.showSnackbar({message: '%s',timeout: %d});", toast.message, toast.timeout))
}
