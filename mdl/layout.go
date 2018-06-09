package mdl

import (
	"github.com/dtylman/gowd"
)

type LayoutHeader struct {
	*gowd.Element
	items                         []*gowd.Element
	spacer, title, nav, container *gowd.Element
}
type LayoutDrawer struct {
	*gowd.Element
	title *gowd.Element
	nav   *gowd.Element
	items []*gowd.Element
}

type NavigationLayout struct {
	*gowd.Element
	css    *gowd.Element
	header *LayoutHeader
	drawer *LayoutDrawer
	body   *gowd.Element
}

type navHeaderModification int

var navHeaderMods = [...]string{
	"mdl-layout__header--transparent",
}

const (
	//NavigationHeaderTransparent makes the header transparent
	NavigationHeaderTransparent navHeaderModification = 1 + iota
)

type navModification int

var navMods = [...]string{
	"mdl-layout--fixed-header",
	"mdl-layout--fixed-drawer",
}

const (
	//NavigationFixedHeader makes the header fixed
	NavigationFixedHeader navModification = 1 + iota
	//NavigationFixedDrawer makes the drawer fixed
	NavigationFixedDrawer
)

//NewNavigationHeader creates a new navigation header for the navigationLayout
func NewNavigationHeader(headerTitle string, navigaionElements []*gowd.Element, headerMods ...navHeaderModification) *LayoutHeader {
	headerLayout := new(LayoutHeader)

	headerLayout.Element = NewElement("header", "mdl-layout__header"+navHeaderToString(headerMods))

	headerLayout.title = NewElement("div", "mdl-layout-title")
	headerLayout.title.SetText(headerTitle)

	headerLayout.spacer = NewElement("div", "mdl-layout-spacer")
	headerLayout.nav = NewElement("nav", "mdl-navigation")

	headerLayout.container = NewElement("div", "mdl-layout__header-row")

	headerLayout.items = navigaionElements
	for _, elem := range headerLayout.items {
		headerLayout.nav.AddElement(elem)
	}

	headerLayout.container.AddElement(headerLayout.title)
	headerLayout.container.AddElement(headerLayout.spacer)
	headerLayout.container.AddElement(headerLayout.nav)

	headerLayout.Element.AddElement(headerLayout.container)
	return headerLayout
}

//NewNavigationDrawer creates a new navigation drawer for the navigationLayout
func NewNavigationDrawer(drawerTitle string, drawerElements []*gowd.Element) *LayoutDrawer {
	drawerlayout := new(LayoutDrawer)

	drawerlayout.title = NewElement("div", "mdl-layout-title")
	drawerlayout.title.SetText(drawerTitle)

	drawerlayout.Element = NewElement("div", "mdl-layout__drawer")
	drawerlayout.nav = NewElement("span", "mdl-navigation")
	drawerlayout.items = drawerElements
	for _, elem := range drawerElements {
		drawerlayout.nav.AddElement(elem)
	}

	drawerlayout.Element.AddElement(drawerlayout.title)
	drawerlayout.Element.AddElement(drawerlayout.nav)
	return drawerlayout
}

//NewNavigationLayout creates a new navigation layout
func NewNavigationLayout(header *LayoutHeader, drawer *LayoutDrawer, body *gowd.Element, navigationTypes ...navModification) *gowd.Element {
	navigationLayout := new(NavigationLayout)
	navigationLayout.header = header
	navigationLayout.drawer = drawer
	navigationLayout.body = NewElement("div", "mdl-layout__content")
	navigationLayout.body.AddElement(body)
	navigationLayout.Element = NewElement("div", "mdl-layout mdl-js-layout "+navToString(navigationTypes))

	navigationLayout.Element.AddElement(navigationLayout.header.Element)
	navigationLayout.Element.AddElement(navigationLayout.drawer.Element)
	navigationLayout.Element.AddElement(navigationLayout.body)

	return navigationLayout.Element
}

func navHeaderToString(nav []navHeaderModification) (res string) {
	for _, mod := range nav {
		res = res + " " + navHeaderMods[mod]
	}
	return
}

func navToString(nav []navModification) (res string) {
	for _, mod := range nav {
		res = res + " " + navMods[mod]
	}
	return
}
