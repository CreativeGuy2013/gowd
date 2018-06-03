package mdl

import (
	"github.com/dtylman/gowd"
)

//Chip is a chip element
type Chip struct {
	*gowd.Element
	Text    *gowd.Element
	Contact *gowd.Element
	Action  struct {
		*gowd.Element
		Icon *gowd.Element
	}
}

//NewChipBasic creates a new basic chip
func NewChipBasic(text string) *Chip {
	newChip := new(Chip)
	newChip.Element = NewElement("span", "mdl-chip")

	newChip.Text = NewElement("span", "mdl-chip__text")
	newChip.Text.SetText(text)

	newChip.AddElement(newChip.Text)

	return newChip
}

//NewChipContactImage creates a new chip with a contact image
func NewChipContactImage(text string, contactImageURL string) *Chip {
	newChip := new(Chip)
	newChip.Element = NewElement("span", "mdl-chip mdl-chip--contact")

	newChip.Text = NewElement("span", "mdl-chip__text")
	newChip.Text.SetText(text)

	newChip.Contact = NewElement("img", "mdl-chip__contact")
	newChip.Contact.SetAttribute("src", contactImageURL)

	newChip.AddElement(newChip.Contact)
	newChip.AddElement(newChip.Text)

	return newChip
}

//NewChipContactText creates a new chip with a contact letter (colors must be a valid MDL colors)
func NewChipContactText(text string, contactLetter string, backgroundColor string, textColor string) *Chip {
	newChip := new(Chip)
	newChip.Element = NewElement("span", "mdl-chip mdl-chip--contact")

	newChip.Text = NewElement("span", "mdl-chip__text")
	newChip.Text.SetText(text)

	newChip.Contact = NewElement("img", "mdl-chip__contact mdl-color--"+backgroundColor+" mdl-color-text--"+textColor)
	newChip.Contact.SetText(contactLetter)

	newChip.AddElement(newChip.Contact)
	newChip.AddElement(newChip.Text)

	return newChip
}

//NewChipAction creates a new chip with an action
func NewChipAction(text string, actionIcon string) *Chip {
	newChip := new(Chip)
	newChip.Element = NewElement("span", "mdl-chip mdl-chip--deletable")

	newChip.Text = NewElement("span", "mdl-chip__text")
	newChip.Text.SetText(text)

	if actionIcon != "" {
		newChip.Action.Element = NewElement("button", "mdl-chip__action")
		newChip.Action.Icon = NewElement("i", "material-icons")
		newChip.Action.Icon.SetText(actionIcon)
		newChip.Action.Element.AddElement(newChip.Action.Icon)
	} else {
		return nil
	}

	newChip.AddElement(newChip.Text)
	newChip.AddElement(newChip.Action.Element)

	return newChip
}

//NewChipContactImageAction creates a new chip with an action and a contact image
func NewChipContactImageAction(text string, contactImageURL string, actionIcon string) *Chip {
	newChip := new(Chip)
	newChip.Element = NewElement("span", "mdl-chip mdl-chip--contact mdl-chip--deletable")

	newChip.Text = NewElement("span", "mdl-chip__text")
	newChip.Text.SetText(text)

	newChip.Contact = NewElement("img", "mdl-chip__contact")
	newChip.Contact.SetAttribute("src", contactImageURL)

	if actionIcon != "" {
		newChip.Action.Element = NewElement("button", "mdl-chip__action")
		newChip.Action.Icon = NewElement("i", "material-icons")
		newChip.Action.Icon.SetText(actionIcon)
		newChip.Action.Element.AddElement(newChip.Action.Icon)
	} else {
		return nil
	}

	newChip.AddElement(newChip.Contact)
	newChip.AddElement(newChip.Text)
	newChip.AddElement(newChip.Action.Element)

	return newChip
}

//NewChipContactTextAction creates a new chip with an action and contact letter (colors must be a valid MDL colors)
func NewChipContactTextAction(text string, contactLetter string, backgroundColor string, textColor string, actionIcon string) *Chip {
	newChip := new(Chip)
	newChip.Element = NewElement("span", "mdl-chip mdl-chip--contact mdl-chip--deletable")

	newChip.Text = NewElement("span", "mdl-chip__text")
	newChip.Text.SetText(text)

	newChip.Contact = NewElement("span", "mdl-chip__contact mdl-color--"+backgroundColor+" mdl-color-text--"+textColor)
	newChip.Contact.SetText(contactLetter)


	if actionIcon != "" {
		newChip.Action.Element = NewElement("button", "mdl-chip__action")
		newChip.Action.Icon = NewElement("i", "material-icons")
		newChip.Action.Icon.SetText(actionIcon)
		newChip.Action.Element.AddElement(newChip.Action.Icon)
	} else {
		return nil
	}

	newChip.AddElement(newChip.Contact)
	newChip.AddElement(newChip.Text)
	newChip.AddElement(newChip.Action.Element)

	return newChip
}
