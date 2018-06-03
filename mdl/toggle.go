package mdl

import (
	"fmt"

	"github.com/dtylman/gowd"
)

//CheckBox is a MDl checkbox element
type CheckBox struct {
	*gowd.Element
	input *gowd.Element
	text  *gowd.Element
}

//IconToggle is a MDl icontoggle element
type IconToggle struct {
	*gowd.Element
	input *gowd.Element
	text  *gowd.Element
}

//Toggle is a MDl toggle switch element
type Toggle struct {
	*gowd.Element
	input *gowd.Element
	text  *gowd.Element
}

//RadioButton is a MDl radio button element
type RadioButton struct {
	*gowd.Element
	input *gowd.Element
	text  *gowd.Element
}

//RadioButtonArray is set of mdl radio buttons in a row, with the same name and work together without hasste
type RadioButtonArray struct {
	*gowd.Element
	buttons []*RadioButton
}

//NewCheckBox creates a new checkbox element
func NewCheckBox(labelText string) *CheckBox {
	checkbox := new(CheckBox)

	checkbox.Element = NewElement("label", "mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect")

	checkbox.input = NewElement("input", "mdl-checkbox__input")
	checkbox.input.SetAttribute("type", "checkbox")

	checkbox.Element.SetAttribute("for", checkbox.input.GetID())

	checkbox.text = NewElement("span", "mdl-checkbox__label")
	checkbox.text.SetText(labelText)

	checkbox.Element.AddElement(checkbox.input)
	checkbox.Element.AddElement(checkbox.text)

	return checkbox
}

//IsChecked checks is a checkbox is checked
func (Checkbox *CheckBox) IsChecked() bool {
	_, worked := Checkbox.input.GetAttribute("checked")
	return worked
}

//NewRadioButton creates a new radio button, groupName is the name of selectable items, label text is the text
func NewRadioButton(groupName string, labelText string, value int) *RadioButton {
	radioButton := new(RadioButton)

	radioButton.Element = NewElement("label", "mdl-radio mdl-js-radio mdl-js-ripple-effect")

	radioButton.input = NewElement("input", "mdl-radio__button")
	radioButton.input.SetAttribute("type", "radio")
	radioButton.input.SetAttribute("value", string(value))
	radioButton.input.SetAttribute("name", groupName)
	radioButton.input.SetAttribute("onclick", fmt.Sprintf("update_radio_button('%s')", radioButton.input.GetID()))

	radioButton.Element.SetAttribute("for", radioButton.input.GetID())

	radioButton.text = NewElement("span", "mdl-checkbox__label")
	radioButton.text.SetText(labelText)
	radioButton.text.SetID(labelText)

	radioButton.Element.AddElement(radioButton.input)
	radioButton.Element.AddElement(radioButton.text)

	return radioButton
}

//NewRadioButtonArray creates a new radio button array, automaticly assigns the groupName and value
func NewRadioButtonArray(labelNames ...string) *RadioButtonArray {
	radioButtonArray := new(RadioButtonArray)
	radioButtonArray.Element = NewElement("div", "")

	for i, labelName := range labelNames {
		radioButton := NewRadioButton(radioButtonArray.Element.GetID(), labelName, i)
		radioButtonArray.buttons = append(radioButtonArray.buttons, radioButton)
		radioButtonArray.Element.AddElement(radioButtonArray.buttons[i].Element)
	}
	return radioButtonArray
}

//IsSelected checks if a radiobutton is selected
func (rb *RadioButton) IsSelected() bool {
	_, worked := rb.input.GetAttribute("checked")
	return worked
}

//GetSelected returns the name of the radiobutton of the selected radiobutton
func (radioButtonArray *RadioButtonArray) GetSelected() string {
	for _, rb := range radioButtonArray.buttons {
		fmt.Printf("%v\n", rb.text.GetID())
		if rb.IsSelected() {
			return rb.text.GetID()
		}
	}
	return ""
}

//NewIconToggle returns a new iconToggle mdl element
func NewIconToggle(iconText string) *IconToggle {
	iconToggle := new(IconToggle)

	iconToggle.Element = NewElement("label", "mdl-icon-toggle mdl-js-icon-toggle mdl-js-ripple-effect")

	iconToggle.input = NewElement("input", "mdl-icon-toggle__input")
	iconToggle.input.SetAttribute("type", "checkbox")

	iconToggle.Element.SetAttribute("for", iconToggle.input.GetID())

	iconToggle.text = NewElement("i", "mdl-icon-toggle__label material-icons")
	iconToggle.text.SetText(iconText)

	iconToggle.Element.AddElement(iconToggle.input)
	iconToggle.Element.AddElement(iconToggle.text)
	return iconToggle
}

//IsChecked checks if a iconToggle is checked
func (iconToggle *IconToggle) IsChecked() bool {
	_, worked := iconToggle.input.GetAttribute("checked")
	return worked
}

//NewToggle returns a new Toggle mdl element
func NewToggle(labelText string) *Toggle {
	toggle := new(Toggle)

	toggle.Element = NewElement("label", "mdl-switch mdl-js-switch mdl-js-ripple-effect")

	toggle.input = NewElement("input", "mdl-switch__input")
	toggle.input.SetAttribute("type", "checkbox")

	toggle.Element.SetAttribute("for", toggle.input.GetID())

	toggle.text = NewElement("span", "mdl-switch__label")
	toggle.text.SetText(labelText)

	toggle.Element.AddElement(toggle.input)
	toggle.Element.AddElement(toggle.text)
	return toggle
}

//IsChecked checks if a oggle is checked
func (toggle *Toggle) IsChecked() bool {
	_, worked := toggle.input.GetAttribute("checked")
	return worked
}
