package mdl

import (
	"strings"

	"github.com/dtylman/gowd"
)

const (
	//CardBorder A mdl card border that can be applied to any inner element of a card
	CardBorder = "mdl-card--border"

	//CardShadow2 is a shadow behind the card with a value of 2 dp
	CardShadow2 = "mdl-shadow--2dp"

	//CardShadow3 is a shadow behind the card with a value of 3 dp
	CardShadow3 = "mdl-shadow--3dp"

	//CardShadow4 is a shadow behind the card with a value of 4 dp
	CardShadow4 = "mdl-shadow--4dp"

	//CardShadow6 is a shadow behind the card with a value of 6 dp
	CardShadow6 = "mdl-shadow--6dp"

	//CardShadow8 is a shadow behind the card with a value of 8 dp
	CardShadow8 = "mdl-shadow--8dp"

	//CardShadow16 is a shadow behind the card with a value of 16 dp
	CardShadow16 = "mdl-shadow--16dp"
)

//NewCardNoImage creates a new card that doesnt have a image
func NewCardNoImageNoMenu(title *gowd.Element, supportingText *gowd.Element, actions *gowd.Element, actionCSS string, cardModifiers ...string) *gowd.Element {
	newCard := NewElement("div", "mdl-card "+strings.Join(cardModifiers, " "))

	newTitle := NewElement("div", "mdl-card__title")
	newTitle.AddElement(title)

	newSupportingText := NewElement("div", "mdl-card__supporting-text")
	newSupportingText.AddElement(supportingText)

	newActions := NewElement("div", "mdl-card__actions "+actionCSS)
	newActions.AddElement(actions)

	newCard.AddElement(newTitle)
	newCard.AddElement(newSupportingText)
	newCard.AddElement(newActions)

	return newCard
}

//NewCardTitle generates a element to use as a title in mdl cards.
func NewCardTitle(title string, css string) *gowd.Element {
	newTitle := NewElement("h2", "mdl-card__title-text")
	newTitle.SetAttribute("style", css)
	newTitle.SetText(title)
	return newTitle
}
