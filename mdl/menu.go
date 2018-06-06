package mdl

import (
	"strings"

	"github.com/dtylman/gowd"
)

//NewMenu creates a new menu
func NewMenu(bindElement *gowd.Element, options []*gowd.Element, menuModifiers ...string) (*gowd.Element, []*gowd.Element) {
	newMenu := new(gowd.Element)
	newMenu = NewElement("ul", "mdl-menu mdl-js-menu "+strings.Join(menuModifiers, " "))
	newMenu.SetAttribute("data-mdl-for", bindElement.GetID())

	var outElements []*gowd.Element

	for _, element := range options {
		newOption := NewElement("li", "mdl-menu__item")
		newOption.AddElement(element)
		newMenu.AddElement(newOption)
		outElements = append(outElements, newOption)
	}

	return newMenu, outElements
}
