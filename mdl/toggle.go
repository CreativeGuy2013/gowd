package mdl

import (
	"fmt"

	"github.com/dtylman/gowd"
)

type Checkbox struct {
	*gowd.Element
	input *gowd.Element
	text  *gowd.Element
}

type RadioButton struct {
	*gowd.Element
	input *gowd.Element
	text  *gowd.Element
}

type RadioButtonArray struct {
	*gowd.Element
	buttons []*RadioButton
}

func NewCheckBox(labelText string) *Checkbox {
	cb := new(Checkbox)

	cb.Element = NewElement("label", "mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect")

	cb.input = NewElement("input", "mdl-checkbox__input")
	cb.input.SetAttribute("type", "checkbox")

	cb.Element.SetAttribute("for", cb.input.GetID())

	cb.text = NewElement("span", "mdl-checkbox__label")
	cb.text.SetText(labelText)

	cb.Element.AddElement(cb.input)
	cb.Element.AddElement(cb.text)

	return cb
}

func (cb *Checkbox) IsChecked() bool {
	_, worked := cb.input.GetAttribute("checked")
	return worked
}

func NewRadioButton(groupName string, labelText string, value int) *RadioButton {
	rb := new(RadioButton)

	rb.Element = NewElement("label", "mdl-radio mdl-js-radio mdl-js-ripple-effect")

	rb.input = NewElement("input", "mdl-radio__button")
	rb.input.SetAttribute("type", "radio")
	rb.input.SetAttribute("value", string(value))
	rb.input.SetAttribute("name", groupName)
	rb.input.SetAttribute("onclick", fmt.Sprintf("update_radio_button(%s)", rb.input.GetID()))

	rb.Element.SetAttribute("for", rb.input.GetID())

	rb.text = NewElement("span", "mdl-checkbox__label")
	rb.text.SetText(labelText)

	rb.Element.AddElement(rb.input)
	rb.Element.AddElement(rb.text)

	return rb
}

func NewRadioButtonArray(labelNames ...string) *RadioButtonArray {
	rba := new(RadioButtonArray)
	rba.Element = NewElement("div", "")

	for i, labelName := range labelNames {
		radioButton := NewRadioButton(rba.Element.GetID(), labelName, i)
		rba.buttons = append(rba.buttons, radioButton)
		rba.Element.AddElement(rba.buttons[i].Element)
	}
	return rba
}

func (rb *RadioButton) IsChecked() bool {
	_, worked := rb.input.GetAttribute("checked")
	return worked
}

func (rba *RadioButtonArray) GetChecked() (string, error) {
	/*for _, rb := range rba.buttons {
		if rb.IsChecked() {
			return rb.text.GetValue(), nil
		}
	}
	return "", errors.New("could not find the checked box")*/
	return "nil", nil
}
