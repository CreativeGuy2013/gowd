package bootstrap

import (
	"github.com/dtylman/gowd"
)

const (
	//ButtonColoured
	ButtonColoured = "mdl-button--colored"
)

//NewButton creates a new bootstrap <button> element
func NewButtonDefault(buttontype string, caption string) *gowd.Element {
	btn := NewElement("button", "btn "+buttontype)
	if caption != "" {
		btn.SetText(caption)
	}
	return btn
}

func NewButtonFAB(buttontype string, icon string) *gowd.Element {
	btn := NewElement("button", "mdl-button mdl-js-button mdl-button--fab"+buttontype)

	if icon != "" {
		btnIcon := NewElement("i", "material-icons")
		btnIcon.SetText(icon)
		btn.AddElement(btnIcon)
	} else {
		return nil
	}
	return btn
}

//NewLinkButton creates a new bootstrap link button (<a>)
func NewLinkButton(caption string) *gowd.Element {
	linkBtn := gowd.NewElement("a")
	linkBtn.SetAttribute("href", "#")
	if caption != "" {
		linkBtn.AddElement(gowd.NewText(caption))
	}
	return linkBtn
}
