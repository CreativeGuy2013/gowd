package mdl

import (
	"github.com/dtylman/gowd"
)

type List struct {
	*gowd.Element
	Elements []ListElement
}

type ListElement struct {
	*gowd.Element
	Primary struct {
		*gowd.Element
		SubText *gowd.Element
	}
	Secondary struct {
		*gowd.Element
		Action *gowd.Element
		Info   *gowd.Element
	}
}

func NewList() *List {
	newList := new(List)
	return newList
}

func NewListElement(primaryContent *gowd.Element, secondaryAction *gowd.Element, secondaryInfo *gowd.Element) *ListElement {
	newListElement := new(ListElement)

	newListElement.Element = NewElement("li", "mdl-list__item")
	newListElement.Primary.Element = NewElement("span", "mdl-list__item-primary-content")

	newListElement.Secondary.Element = NewElement("span", "mdl-list__item-secondary-content")
	
	if action != nil {
		newListElement.Secondary.Action = NewElement("span", "mdl-list__item-primary-content")
	}

	return newListElement
}

func NewListElementShortSubtitle(primaryContent *gowd.Element, primarySubtitle *gowd.Element, secondaryAction *gowd.Element, secondaryInfo *gowd.Element) *ListElement {
	newListElement := new(ListElement)

	return newListElement
}

func NewListElementLongSubtitle(primaryContent *gowd.Element, primarySubtitle *gowd.Element, secondaryAction *gowd.Element, secondaryInfo *gowd.Element) *ListElement {
	newListElement := new(ListElement)
	return newListElement
}
