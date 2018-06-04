package mdl

import (
	"fmt"
	"strconv"

	"github.com/jaicewizard/gowd"
)

type Slider struct {
	*gowd.Element
}

func NewSlider(min, max, value, step int) *Slider {
	slider := new(Slider)
	slider.Element = NewElement("input", "mdl-slider mdl-js-slider")
	slider.Element.SetAttribute("type", "range")
	slider.Element.SetAttribute("min", fmt.Sprintf("%d", min))
	slider.Element.SetAttribute("max", fmt.Sprintf("%d", max))
	slider.Element.SetAttribute("value", fmt.Sprintf("%d", value))
	slider.Element.SetAttribute("step", fmt.Sprintf("%d", step))
	return slider
}

func (slider *Slider) GetValue() int {
	value, _ := slider.Element.GetAttribute("value")
	i, err := strconv.Atoi(value)
	if err != nil {
		return -1
	}
	return i
}
