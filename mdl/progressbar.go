package mdl

import (
	"fmt"
	"strings"

	"github.com/dtylman/gowd"
)

const (
	//SpinnerSingleColour makes the button the primary mdl color.
	SpinnerSingleColour = "mdl-spinner--single-color"
)

type ProgressBar struct {
	*gowd.Element
	percentage int
	buffering  int
}

//NewProgressBar creates a new progressbar
func NewProgressBar() *ProgressBar {
	progress := new(ProgressBar)
	progress.Element = NewElement("div", "mdl-progress mdl-js-progress")
	progress.SetPercent(0)
	return progress
}

func NewLoadingBar() *gowd.Element {
	loadingBar := NewElement("div", "mdl-progress mdl-js-progress mdl-progress__indeterminate")
	return loadingBar
}

//SetPercent sets the value of the progress bar as a percentage
func (pb *ProgressBar) SetPercent(percent int) {
	pb.percentage = percent
	execJS(fmt.Sprintf("document.getElementById('%q').MaterialProgress.setProgress(%d)", pb.GetID(), pb.percentage))
}

//SetBuffering sets the value of the progress bar as a percentage
func (pb *ProgressBar) SetBuffering(percent int) {
	pb.percentage = percent
	execJS(fmt.Sprintf("document.getElementById('%q').MaterialProgress.setBuffer(%d)", pb.GetID(), pb.percentage))
}

func newSpinner(spinnertype ...string) *gowd.Element {
	loadingBar := NewElement("div", "mdl-spinner mdl-js-spinner is-active"+strings.Join(spinnertype, " "))
	return loadingBar
}
