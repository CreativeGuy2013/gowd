package mdl

import (
	"strings"

	"github.com/jaicewizard/gowd"
)

//Badge is a badge element
type Badge struct {
	*gowd.Element
	Content *gowd.Element
}

const (
	//BadgeOverlap makes the badges overlap wth its container.
	BadgeOverlap = "mdl-badge--overlap"

	//BadgeNoBackground applies open-circle effect to badge.
	BadgeNoBackground = "mdl-badge--no-background"
)

//NewBadge creates a new badge
func NewBadge(badge string, content *gowd.Element, badgeTags ...string) *Badge {
	newBadge := new(Badge)
	newBadge.Element = NewElement("span", "mdl-badge "+strings.Join(badgeTags, " "))
	newBadge.Element.SetAttribute("data-badge", badge)

	newBadge.Content = content

	newBadge.Element.AddElement(newBadge.Content)

	return newBadge
}

//NewBadgeIcon creates a new badge with icon
func NewBadgeIcon(badge string, icon string, badgeTags ...string) *Badge {
	newBadge := new(Badge)
	newBadge.Element = NewElement("div", "material-icons mdl-badge "+strings.Join(badgeTags, " "))
	newBadge.Element.SetText(icon)
	newBadge.Element.SetAttribute("data-badge", badge)

	return newBadge
}

//GetBadge Gets the current text in the badge.
func (badge *Badge) GetBadge() string {
	badgeReturn, _ := badge.Element.GetAttribute("data-badge")
	return badgeReturn
}

//SetBadge Sets text in the badge.
func (badge *Badge) SetBadge(badgeText string) {
	badge.Element.SetAttribute("data-badge", badgeText)
}
