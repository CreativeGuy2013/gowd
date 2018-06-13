package mdl

import (
	"github.com/dtylman/gowd"
)

//TableRow represents a <td> or <th>
type TableCell struct {
	*gowd.Element
	Modifier string
}

//Table represents <table>
type Table struct {
	*gowd.Element
	Head *gowd.Element
	Body *gowd.Element
}

//NewTable creates a new table with type
func NewTable(headRow []TableCell, table [][]TableCell) *Table {
	newTable := new(Table)
	newTable.Element = NewElement("table", "mdl-data-table mdl-js-data-table")
	newTable.Head = gowd.NewElement("thead")
	newTable.AddElement(newTable.Head)
	newTable.Body = gowd.NewElement("tbody")
	newTable.AddElement(newTable.Body)
	headerRow := NewElement("tr", "")
	newTable.Head.AddElement(headerRow)

	for _, element := range headRow {
		newCell := NewElement("th", element.Modifier)
		newCell.AddElement(element.Element)
		headerRow.AddElement(newCell)
	}

	for _, row := range table {
		newRow := NewElement("tr", "")
		newTable.Body.AddElement(newRow)
		for _, element := range row {
			newCell := NewElement("td", element.Modifier)
			newCell.AddElement(element.Element)
			newRow.AddElement(newCell)
		}
	}
	return newTable
}
