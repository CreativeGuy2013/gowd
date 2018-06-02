package bootstrap

import (
	"strings"

	"github.com/dtylman/gowd"
)

const (
	//ButtonPrimary makes the button the primary mdl color.
	ButtonPrimary = "mdl-button--primary"

	//ButtonAccent makes the button the secondary mdl color.
	ButtonAccent = "mdl-button--accent"

	//ButtonRippled makes the button the have a ripple effect.
	ButtonRippled = "mdl-js-ripple-effect"
)

//NewButtonFab creates a new fab mdl <button> element
func NewButtonFab(icon string, enabled bool, buttontype ...string) *gowd.Element {
	btn := NewElement("button", "mdl-button mdl-js-button mdl-button--fab "+strings.Join(buttontype, " "))

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

func NewbuttonRaised(caption string, enabled bool, buttontype ...string) *gowd.Element {
	btn := NewElement("button", "mdl-button mdl-js-button mdl-button--raised "+strings.Join(buttontype, " "))

	if caption != "" {
		btn.SetText(caption)
	}
	if !enabled {
		btn.Disable()
	}
	return btn
}

func NewButtonFlat(caption string, enabled bool, buttontype ...string) *gowd.Element {
	btn := NewElement("button", "mdl-button mdl-js-button "+strings.Join(buttontype, " "))

	if caption != "" {
		btn.SetText(caption)
	}
	if !enabled {
		btn.Disable()
	}
	return btn
}
