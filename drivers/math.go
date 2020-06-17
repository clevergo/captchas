// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"image/color"

	"clevergo.tech/captchas"
	"github.com/mojocn/base64Captcha"
)

// MathOption is a function that receives a pointer of math driver.
type MathOption func(*Math)

// MathHeight sets height.
func MathHeight(height int) MathOption {
	return func(m *Math) {
		m.height = height
	}
}

// MathWidth sets width.
func MathWidth(width int) MathOption {
	return func(m *Math) {
		m.width = width
	}
}

// MathNoiseCount sets noise count.
func MathNoiseCount(count int) MathOption {
	return func(m *Math) {
		m.noiseCount = count
	}
}

// MathBGColor sets background color.
func MathBGColor(color *color.RGBA) MathOption {
	return func(m *Math) {
		m.bgColor = color
	}
}

// MathFonts sets fonts.
func MathFonts(fonts []string) MathOption {
	return func(m *Math) {
		m.fonts = fonts
	}
}

// Math is a math driver.
type Math struct {
	*driver
	// captcha png height in pixel.
	height int
	// captcha png width in pixel.
	width int
	// text noise count.
	noiseCount      int
	showLineOptions int
	// background color.
	bgColor *color.RGBA
	fonts   []string
}

var _ captchas.Driver = NewMath()

// NewMath return a math driver.
func NewMath(opts ...MathOption) *Math {
	d := &Math{
		driver:     &driver{htmlTag: htmlTagIMG},
		height:     80,
		width:      220,
		noiseCount: 0,
		fonts:      []string{"wqy-microhei.ttc"},
	}

	for _, f := range opts {
		f(d)
	}

	d.driver.driver = base64Captcha.NewDriverMath(d.height, d.width, d.noiseCount, d.showLineOptions, d.bgColor, d.fonts)

	return d
}
