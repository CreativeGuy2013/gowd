package mdl

import (
	"strings"

	"github.com/dtylman/gowd"
)

const (
	//TooltipLocationLeft makes the tooltip appear on the left hand side of the element.
	TooltipLocationLeft = "mdl-tooltip--left"

	//TooltipLocationRight makes the tooltip appear on the right hand side of the element.
	TooltipLocationRight = "mdl-tooltip--right"

	//TooltipLocationTop makes the tooltip appear on the top side of the element.
	TooltipLocationTop = "mdl-tooltip--top"

	//TooltipLocationBottom makes the tooltip appear on the bottom side of the element.
	TooltipLocationBottom = "mdl-tooltip--bottom"

	//TooltipLarge makes the tooltip larger.
	TooltipLarge = "mdl-tooltip--large"
)

//NewTooltip creates a new tooltip
func NewTooltip(tooltip string, tooltipElement *gowd.Element, tooltipTags ...string) *gowd.Element {
	newTooltip := new(gowd.Element)
	newTooltip = NewElement("span", "mdl-tooltip "+strings.Join(tooltipTags, " "))
	newTooltip.SetAttribute("for", tooltipElement.GetID())
	newTooltip.SetText(tooltip)

	return newTooltip
}
