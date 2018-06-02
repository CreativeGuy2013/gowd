# gowd Styling

This is a addon to dtylman's gowd, a library to directly interface with HTML/JS/CSS using GO, that adds more options to style elements.


### How to use this library:

1. Install (gowd)[github.com/dtylman/gowd]`
1. Choose the styling package you want
1. For **mdl** 
	1. Get the package using the GO CLI: `go get github.com/CreativeGuy2013/gowd-styling/[name of styling package]` e.g. `go get github.com/CreativeGuy2013/gowd-styling/mdl`.
	1. In your HTML files import the MDL CSS and JS. These can be found (here)[https://getmdl.io/started/index.html]
1. Add elements like shown below.


### Available Objects:

**MDL:**
	- Buttons
		- Elements
			- ButtonFlat
			- ButtonRaised
			- ButtonFAB
			- ButtonFABMini
			- ButtonIcon
		- Modifiers
			- ButtonPrimary
			- ButtonAccent
			- ButtonRipple

	- Inputs
		- Elements
			- Input
			- InputPattern
			- InputMultiline
			- InputPatternMultiline
			- InputExpanding
			- InputPatternExpanding
		- Modifiers
			- InputFloatingLabel

	- Progress Bar
		- Elements
			- ProgressBar
			- LoadingBar
			- Spinner
		- Modifiers
			- SpinnerSingleColour (Spinner only)

### Usage (MDL)

Adding a button:

```go

import (
	"github.com/dtylman/gowd"
	"github.com/creativeguy2013/gowd-styling/mdl"
)

func main() {
	body := gowd.NewElement("div")
	btn := mdl.NewButtonRaised("Click Me")
	btn.OnEvent(gowd.OnClick, btnClicked)

	body.AddElement(btn)
	gowd.Run(body)
}

func btnClicked(sender *gowd.Element, event *gowd.EventElement) {
	sender.SetText("Clicked!")
}
```

