package mdl

import (
	"github.com/dtylman/gowd"
)

//NewIcon makes a new icon
func NewIcon(iconName string) *gowd.Element {
	newIcon := NewElement("i", "material-icons")
	newIcon.SetText(iconName)

	return newIcon
}
