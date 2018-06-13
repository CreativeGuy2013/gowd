package mdl

import (
	"fmt"

	"github.com/dtylman/gowd"
)

//CheckBox is a MDl checkbox element
type CheckBox struct {
	checkBox
}

//IconToggle is a MDl icontoggle element
type IconToggle struct {
	checkBox
}

//Toggle is a MDl toggle switch element
type Toggle struct {
	checkBox
}

//RadioButton is a MDl radio button element
type RadioButton struct {
	checkBox
}

type checkBox struct {
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
func NewCheckBox(labelText string, isChecked bool) *CheckBox {
	checkbox := new(CheckBox)

	checkbox.Element = NewElement("label", "mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect")

	checkbox.input = NewElement("input", "mdl-checkbox__input")
	checkbox.input.SetAttribute("type", "checkbox")
	if isChecked {
		checkbox.input.SetAttribute("checked", "")
	}
	checkbox.Element.SetAttribute("for", checkbox.input.GetID())

	checkbox.text = NewElement("span", "mdl-checkbox__label")
	checkbox.text.SetText(labelText)

	checkbox.Element.AddElement(checkbox.input)
	checkbox.Element.AddElement(checkbox.text)

	return checkbox
}

//NewRadioButton creates a new radio button, groupName is the name of selectable items, label text is the text
func NewRadioButton(groupName string, labelText string, value int, isChecked bool) *RadioButton {
	radioButton := new(RadioButton)

	radioButton.Element = NewElement("label", "mdl-radio mdl-js-radio mdl-js-ripple-effect")

	radioButton.input = NewElement("input", "mdl-radio__button")
	radioButton.input.SetAttribute("type", "radio")
	radioButton.input.SetAttribute("value", string(value))
	radioButton.input.SetAttribute("name", groupName)
	radioButton.input.SetAttribute("onclick", fmt.Sprintf("update_radio_button('%s')", radioButton.input.GetID()))
	if isChecked {
		radioButton.input.SetAttribute("checked", "")
	}
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
		radioButton := NewRadioButton(radioButtonArray.Element.GetID(), labelName, i, false)
		radioButtonArray.buttons = append(radioButtonArray.buttons, radioButton)
		radioButtonArray.Element.AddElement(radioButtonArray.buttons[i].Element)
	}
	return radioButtonArray
}

//GetSelected returns the name of the radiobutton of the selected radiobutton
func (radioButtonArray *RadioButtonArray) GetSelected() string {
	for _, rb := range radioButtonArray.buttons {
		fmt.Printf("%v\n", rb.text.GetID())
		if rb.IsChecked() {
			return rb.text.GetID()
		}
	}
	return ""
}

//NewIconToggle returns a new iconToggle mdl element
func NewIconToggle(iconText string, isChecked bool) *IconToggle {
	iconToggle := new(IconToggle)

	iconToggle.Element = NewElement("label", "mdl-icon-toggle mdl-js-icon-toggle mdl-js-ripple-effect")

	iconToggle.input = NewElement("input", "mdl-icon-toggle__input")
	iconToggle.input.SetAttribute("type", "checkbox")
	if isChecked {
		iconToggle.input.SetAttribute("checked", "")
	}

	iconToggle.Element.SetAttribute("for", iconToggle.input.GetID())

	iconToggle.text = NewElement("i", "mdl-icon-toggle__label material-icons")
	iconToggle.text.SetText(iconText)

	iconToggle.Element.AddElement(iconToggle.input)
	iconToggle.Element.AddElement(iconToggle.text)
	return iconToggle
}

//NewToggle returns a new Toggle mdl element
func NewToggle(labelText string, isChecked bool) *Toggle {
	toggle := new(Toggle)

	toggle.Element = NewElement("label", "mdl-switch mdl-js-switch mdl-js-ripple-effect")

	toggle.input = NewElement("input", "mdl-switch__input")
	toggle.input.SetAttribute("type", "checkbox")
	if isChecked {
		toggle.input.SetAttribute("checked", "")
	}
	toggle.Element.SetAttribute("for", toggle.input.GetID())

	toggle.text = NewElement("span", "mdl-switch__label")
	toggle.text.SetText(labelText)

	toggle.Element.AddElement(toggle.input)
	toggle.Element.AddElement(toggle.text)
	return toggle
}

//IsChecked checks if a checkbox is checked is checked
func (checkbox *checkBox) IsChecked() bool {
	_, worked := checkbox.input.GetAttribute("checked")
	return worked
}
