package mdl

import (
	"fmt"
	"strings"

	"github.com/dtylman/gowd"
)

const (
	//SpinnerSingleColour makes the button the primary mdl color.
	SpinnerSingleColour = "mdl-spinner--single-color"

	//LoadingBar this makes a progressbar a loader
	LoadingBar = "mdl-progress__indeterminate"
)

//ProgressBar is a progress bar element
type ProgressBar struct {
	*gowd.Element
	percentage int
	buffering  int
}

//NewProgressBar creates a new progressbar, you can set the percentage and buffer value
func NewProgressBar(style string) *ProgressBar {
	progress := new(ProgressBar)
	progress.Element = NewElement("div", "mdl-progress mdl-js-progress "+style)
	return progress
}

//SetPercentage sets the value of the progress bar as a percentage
func (pb *ProgressBar) SetPercentage(percent int) {
	pb.percentage = percent
	gowd.ExecJSNow(fmt.Sprintf("document.getElementById(%q).MaterialProgress.setProgress(%d)", pb.GetID(), pb.percentage))
}

//SetBuffering sets the value of the progress bar as a percentage
func (pb *ProgressBar) SetBuffering(percent int) {
	pb.percentage = percent
	gowd.ExecJSNow(fmt.Sprintf("document.getElementById(%q).MaterialProgress.setBuffer(%d)", pb.GetID(), pb.percentage))
}

//NewSpinner creates a new spinning loader
func NewSpinner(spinnertype ...string) *gowd.Element {
	loadingBar := NewElement("div", "mdl-spinner mdl-js-spinner is-active"+strings.Join(spinnertype, " "))
	return loadingBar
}
