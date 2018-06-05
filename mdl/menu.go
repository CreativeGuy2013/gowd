package mdl

import (
	"strings"

	"github.com/jaicewizard/gowd"
)

//NewMenu creates a new menu
func NewMenu(bindElement *gowd.Element, options []*gowd.Element, menuModifiers ...string) *gowd.Element {
	newMenu := new(gowd.Element)
	newMenu = NewElement("ul", "mdl-menu mdl-js-menu "+strings.Join(menuModifiers, " "))
	newMenu.SetAttribute("data-mdl-for", bindElement.GetID())

	for _, element := range options {
		newOption := NewElement("li", "mdl-menu__item")
		newOption.AddElement(element)
		newMenu.AddElement(newOption)
	}

	return newMenu
}
