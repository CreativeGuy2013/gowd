package mdl

import "github.com/dtylman/gowd"

type Input struct {
	*gowd.Element
	Input   *gowd.Element
	Label  	*gowd.Element
	Help 	*gowd.Element
}

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

/*//NewInput creates a new text input
func NewInputText(label string, enabled bool, inputTags ...string) *Input {
	newInput := new(Input)
	newInput.Element = NewElement("div", "mdl-textfield mdl-js-textfield")

	newInput.Input = NewElement("input", "mdl-textfield__input")
	newInput.Input.SetAttribute("type", "text")

	newInput.Label = NewElement("label", "mdl-textfield__label")
	newInput.Label.SetAttribute("for", newInput.Input.GetID())
	newInput.Label.SetText(label)
	
	if !enabled {
		newInput.Input.Disable()
	}

	newInput.AddElement(newInput.Input)
	newInput.AddElement(newInput.Label)

	return newInput
}*/

//NewInput creates a new text input with a provided pattern
func NewInputText(label string, enabled bool, pattern) *Input {
	newInput := new(Input)
	newInput.Element = NewElement("div", "mdl-textfield mdl-js-textfield")

	newInput.Input = NewElement("input", "mdl-textfield__input")
	newInput.Input.SetAttribute("type", "text")

	newInput.Label = NewElement("label", "mdl-textfield__label")
	newInput.Label.SetAttribute("for", newInput.Input.GetID())
	newInput.Label.SetText(label)
	
	if !enabled {
		newInput.Input.Disable()
	}

	newInput.AddElement(newInput.Input)
	newInput.AddElement(newInput.Label)

	return newInput
}*/
