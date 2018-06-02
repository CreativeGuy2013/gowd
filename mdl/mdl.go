package mdl

import (
	"fmt"
	"os"

	"github.com/dtylman/gowd"
)

//NewElement returns new mdl element
func NewElement(tag, class string, kids ...*gowd.Element) *gowd.Element {
	elem := gowd.NewElement(tag)
	if class != "" {
		elem.SetAttribute("class", class)
	}
	for _, kid := range kids {
		elem.AddElement(kid)
	}
	return elem
}

//NewContainer returns new mdl container.
func NewContainer(fluid bool, kids ...*gowd.Element) *gowd.Element {
	if fluid {
		return NewElement("div", "container-fluid", kids...)
	}
	return NewElement("div", "container", kids...)
}

//NewFormGroup returns new bootsrap form group
func NewFormGroup(elems ...*gowd.Element) *gowd.Element {
	return NewElement("div", "form-group", elems...)
}

//NewRow return new mdl row
func NewRow(elems ...*gowd.Element) *gowd.Element {
	return NewElement("div", "row", elems...)
}

const (
	//ColumnLarge <col-lg>
	ColumnLarge = "col-lg"
	//ColumnMedium <col-md>
	ColumnMedium = "col-md"
	//ColumnSmall <col-sm>
	ColumnSmall = "col-sm"
	//ColumnXtraSmall <col-xs>
	ColumnXtraSmall = "col-xs"
)

//NewColumn returns new mdl column
func NewColumn(size string, span int, elems ...*gowd.Element) *gowd.Element {
	return NewElement("div", fmt.Sprintf("%s-%d", size, span), elems...)
}

func execJS(js string) {
	fmt.Fprintln(os.Stdout, js)
}
