package mdl

import (
	"strings"

	"github.com/dtylman/gowd"
)

//Dialog is a dialog element
type Dialog struct {
	*gowd.Element
	Title   *gowd.Element
	Content *gowd.Element
	Actions *gowd.Element
}

const (
	//DialogFullWidthActions makes the buttons in the action bar of the dialog full width.
	DialogFullWidthActions = "mdl-dialog__actions--full-width"
)

//NewDialog creates a new dialog
func NewDialog(title string, content *gowd.Element, actions *gowd.Element, dialogActionTags ...string) *Dialog {
	newDialog := new(Dialog)
	newDialog.Element = NewElement("dialog", "mdl-dialog")

	newDialog.Title = NewElement("h4", "mdl-dialog__title")
	newDialog.Title.SetText(title)

	newDialog.Content = NewElement("div ", "mdl-dialog__content")
	newDialog.Content.AddElement(content)

	newDialog.Actions = NewElement("div", "mdl-mdl-dialog__actions "+strings.Join(dialogActionTags, " "))
	newDialog.Actions.AddElement(actions)

	newDialog.AddElement(newDialog.Title)
	newDialog.AddElement(newDialog.Content)
	newDialog.AddElement(newDialog.Actions)

	return newDialog
}

//Show Shows the dialog without modal.
func (dialog *Dialog) Show() {
	gowd.ExecJSNow("document.getElementById('" + dialog.Element.GetID() + "').show()")
}

//ShowModalNow Shows the dialog with modal now.
func (dialog *Dialog) ShowModalNow() {
	gowd.ExecJSNow("document.getElementById('" + dialog.Element.GetID() + "').showModal()")
}

//ShowModal Shows the dialog with modal.
func (dialog *Dialog) ShowModal() {
	gowd.ExecJS("document.getElementById('" + dialog.Element.GetID() + "').showModal()")
}

//Close Closes the dialog.
func (dialog *Dialog) Close() {
	gowd.ExecJSNow("document.getElementById('" + dialog.Element.GetID() + "').close()")
}
