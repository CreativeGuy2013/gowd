package bootstrap

import (
	"strings"

	"github.com/dtylman/gowd"
)

const (
	//ButtonColored makes the button the primary color.
	ButtonColored = "mdl-button--colored"

	//ButtonColored makes the button the have a ripple effect.
	ButtonRippled = "mdl-js-ripple-effect"
)

//NewButtonDefault creates a new standard mdl <button> element
func NewButtonDefault(caption string, buttontype ...string) *gowd.Element {
	btn := NewElement("button", "btn "+strings.Join(buttontype, ", "))
	if caption != "" {
		btn.SetText(caption)
	}
	return btn
}

//NewButtonFab creates a new fab mdl <button> element
func NewButtonFab(icon string, enabled bool, buttontype ...string) *gowd.Element {
	btn := NewElement("button", "mdl-button mdl-js-button mdl-button--fab"+strings.Join(buttontype, ", "))

	if icon != "" {
		btnIcon := NewElement("i", "material-icons")
		btnIcon.SetText(icon)
		btn.AddElement(btnIcon)
	} else {
		return nil
	}
	if !enabled {
		btn.Disable()
	}
	return btn
}

//NewButtonFabMini creates a new mini fab mdl <button> element
func NewButtonFabMini(icon string, enabled bool, buttontype ...string) *gowd.Element {
	btn := NewElement("button", "mdl-button mdl-js-button mdl-button--fab button--mini-fab"+strings.Join(buttontype, ", "))

	if icon != "" {
		btnIcon := NewElement("i", "material-icons")
		btnIcon.SetText(icon)
		btn.AddElement(btnIcon)
	} else {
		return nil
	}
	if !enabled {
		btn.Disable()
	}
	return btn
}

//NewButtonIcon creates a new icon mdl <button> element
func NewButtonIcon(icon string, enabled bool, buttontype ...string) *gowd.Element {
	btn := NewElement("button", "mdl-button mdl-js-button mdl-button--icon"+strings.Join(buttontype, ", "))

	if icon != "" {
		btnIcon := NewElement("i", "material-icons")
		btnIcon.SetText(icon)
		btn.AddElement(btnIcon)
	} else {
		return nil
	}
	if !enabled {
		btn.Disable()
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
