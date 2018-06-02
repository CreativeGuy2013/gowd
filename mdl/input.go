package mdl

import "github.com/dtylman/gowd"

//NewInputGroup creates new bootsrap input group from the given elements
func NewInputGroup(elems ...*gowd.Element) *gowd.Element {
	inputGroup := NewElement("div", "input-group")
	for _, elem := range elems {
		inputGroup.AddElement(elem)
	}
	return inputGroup
}

const (
	//InputFloatingLabel is <input type=text>
	InputFloatingLabel = "mdl-textfield--floating-label"
)

//NewInput creates a new input with a provided type
func NewInputText(inputType string) *gowd.Element {
	input := NewElement("input")
	input.SetAttribute("type", inputType)
	return input
}
