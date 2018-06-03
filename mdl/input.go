package mdl

import (
	"strings"

	"github.com/dtylman/gowd"
)

type Input struct {
	*gowd.Element
	Input *gowd.Element
	Label *gowd.Element
	Error *gowd.Element
}

type InputExpandable struct {
	*gowd.Element
	ExpandableHolder struct {
		*gowd.Element
		Input *gowd.Element
		Label *gowd.Element
		Error *gowd.Element
	}
	Icon struct {
		*gowd.Element
		Icon *gowd.Element
	}
}

const (
	//InputFloatingLabel is <input type=text>
	InputFloatingLabel = "mdl-textfield--floating-label"

	//InputTypeText is type text
	InputTypeText = "text"

	//InputTypeNumber is type number
	InputTypeNumber = "number"

	//InputTypePassword is type password
	InputTypePassword = "password"
)

//NewInput creates a new input
func NewInput(label string, inputType string, enabled bool, inputTags ...string) *Input {
	newInput := new(Input)
	newInput.Element = NewElement("div", "mdl-textfield mdl-js-textfield "+strings.Join(inputTags, " "))

	newInput.Input = NewElement("input", "mdl-textfield__input")
	newInput.Input.SetAttribute("type", inputType)

	newInput.Label = NewElement("label", "mdl-textfield__label")
	newInput.Label.SetAttribute("for", newInput.Input.GetID())
	newInput.Label.SetText(label)

	if !enabled {
		newInput.Input.Disable()
	}

	newInput.AddElement(newInput.Input)
	newInput.AddElement(newInput.Label)

	return newInput
}

//NewInputPattern creates a new input with a provided pattern
func NewInputPattern(label string, inputType string, enabled bool, pattern string, errorLabel string, inputTags ...string) *Input {
	newInput := new(Input)
	newInput.Element = NewElement("div", "mdl-textfield mdl-js-textfield "+strings.Join(inputTags, " "))

	newInput.Input = NewElement("input", "mdl-textfield__input")
	newInput.Input.SetAttribute("type", inputType)
	newInput.Input.SetAttribute("pattern", pattern)

	newInput.Label = NewElement("label", "mdl-textfield__label")
	newInput.Label.SetAttribute("for", newInput.Input.GetID())
	newInput.Label.SetText(label)

	newInput.Error = NewElement("span ", "mdl-textfield__error")
	newInput.Error.SetText(errorLabel)

	if !enabled {
		newInput.Input.Disable()
	}

	newInput.AddElement(newInput.Input)
	newInput.AddElement(newInput.Label)
	newInput.AddElement(newInput.Error)

	return newInput
}

//NewInputMultiline creates a new multiline input
func NewInputMultiline(label string, inputType string, enabled bool, rows int, inputTags ...string) *Input {
	newInput := new(Input)
	newInput.Element = NewElement("div", "mdl-textfield mdl-js-textfield "+strings.Join(inputTags, " "))

	newInput.Input = NewElement("textarea", "mdl-textfield__input")
	newInput.Input.SetAttribute("type", inputType)
	newInput.Input.SetAttribute("rows", string(rows))

	newInput.Label = NewElement("label", "mdl-textfield__label")
	newInput.Label.SetAttribute("for", newInput.Input.GetID())
	newInput.Label.SetText(label)

	if !enabled {
		newInput.Input.Disable()
	}

	newInput.AddElement(newInput.Input)
	newInput.AddElement(newInput.Label)

	return newInput
}

//NewInputPatternMultiline creates a new multiline input with a provided pattern
func NewInputPatternMultiline(label string, inputType string, enabled bool, pattern string, errorLabel string, rows int, inputTags ...string) *Input {
	newInput := new(Input)
	newInput.Element = NewElement("div", "mdl-textfield mdl-js-textfield "+strings.Join(inputTags, " "))

	newInput.Input = NewElement("textarea", "mdl-textfield__input")
	newInput.Input.SetAttribute("type", inputType)
	newInput.Input.SetAttribute("pattern", pattern)
	newInput.Input.SetAttribute("rows", string(rows))

	newInput.Label = NewElement("label", "mdl-textfield__label")
	newInput.Label.SetAttribute("for", newInput.Input.GetID())
	newInput.Label.SetText(label)

	newInput.Error = NewElement("span ", "mdl-textfield__error")
	newInput.Error.SetText(errorLabel)

	if !enabled {
		newInput.Input.Disable()
	}

	newInput.AddElement(newInput.Input)
	newInput.AddElement(newInput.Label)
	newInput.AddElement(newInput.Error)

	return newInput
}

//NewInputExpanding creates a new expanable input
func NewInputExpanding(label string, inputType string, enabled bool, icon string, inputTags ...string) *InputExpandable {
	newInput := new(InputExpandable)
	newInput.Element = NewElement("div", "mdl-textfield mdl-js-textfield mdl-textfield--expandable "+strings.Join(inputTags, " "))

	newInput.ExpandableHolder.Element = NewElement("div", "mdl-textfield__expandable-holder")

	newInput.ExpandableHolder.Input = NewElement("input", "mdl-textfield__input")
	newInput.ExpandableHolder.Input.SetAttribute("type", inputType)

	newInput.ExpandableHolder.Label = NewElement("label", "mdl-textfield__label")
	newInput.ExpandableHolder.Label.SetAttribute("for", newInput.ExpandableHolder.Input.GetID())
	newInput.ExpandableHolder.Label.SetText(label)

	newInput.Icon.Element = NewElement("label", "mdl-button mdl-js-button mdl-button--icon")
	newInput.Icon.Element.SetAttribute("for", newInput.ExpandableHolder.Input.GetID())

	if icon != "" {
		newInput.Icon.Icon = NewElement("i", "material-icons")
		newInput.Icon.Icon.SetText(icon)
	} else {
		return nil
	}

	if !enabled {
		newInput.ExpandableHolder.Input.Disable()
	}

	newInput.ExpandableHolder.AddElement(newInput.ExpandableHolder.Input)
	newInput.ExpandableHolder.AddElement(newInput.ExpandableHolder.Label)
	newInput.ExpandableHolder.AddElement(newInput.ExpandableHolder.Error)
	newInput.AddElement(newInput.ExpandableHolder.Element)
	newInput.Icon.AddElement(newInput.Icon.Icon)
	newInput.AddElement(newInput.Icon.Element)

	return newInput
}

//NewInputPatternExpanding creates a new expandable input with a provided pattern
func NewInputPatternExpanding(label string, inputType string, enabled bool, pattern string, errorLabel string, icon string, inputTags ...string) *InputExpandable {
	newInput := new(InputExpandable)
	newInput.Element = NewElement("div", "mdl-textfield mdl-js-textfield mdl-textfield--expandable "+strings.Join(inputTags, " "))

	newInput.ExpandableHolder.Element = NewElement("div", "mdl-textfield__expandable-holder")

	newInput.ExpandableHolder.Input = NewElement("input", "mdl-textfield__input")
	newInput.ExpandableHolder.Input.SetAttribute("type", inputType)
	newInput.ExpandableHolder.Input.SetAttribute("pattern", pattern)

	newInput.ExpandableHolder.Label = NewElement("label", "mdl-textfield__label")
	newInput.ExpandableHolder.Label.SetAttribute("for", newInput.ExpandableHolder.Input.GetID())
	newInput.ExpandableHolder.Label.SetText(label)

	newInput.ExpandableHolder.Error = NewElement("span ", "mdl-textfield__error")
	newInput.ExpandableHolder.Error.SetText(errorLabel)

	newInput.Icon.Element = NewElement("label", "mdl-button mdl-js-button mdl-button--icon")
	newInput.Icon.Element.SetAttribute("for", newInput.ExpandableHolder.Input.GetID())

	if icon != "" {
		newInput.Icon.Icon = NewElement("i", "material-icons")
		newInput.Icon.Icon.SetText(icon)
	} else {
		return nil
	}

	if !enabled {
		newInput.ExpandableHolder.Input.Disable()
	}

	newInput.ExpandableHolder.AddElement(newInput.ExpandableHolder.Input)
	newInput.ExpandableHolder.AddElement(newInput.ExpandableHolder.Label)
	newInput.ExpandableHolder.AddElement(newInput.ExpandableHolder.Error)
	newInput.AddElement(newInput.ExpandableHolder.Element)
	newInput.Icon.AddElement(newInput.Icon.Icon)
	newInput.AddElement(newInput.Icon.Element)

	return newInput
}
