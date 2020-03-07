// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"image/color"

	"github.com/clevergo/captchas"
	"github.com/mojocn/base64Captcha"
)

// MathOption is a function that receives a pointer of math driver.
type MathOption func(*math)

// MathHeight sets height.
func MathHeight(height int) MathOption {
	return func(m *math) {
		m.height = height
	}
}

// MathWidth sets width.
func MathWidth(width int) MathOption {
	return func(m *math) {
		m.width = width
	}
}

// MathNoiseCount sets noise count.
func MathNoiseCount(count int) MathOption {
	return func(m *math) {
		m.noiseCount = count
	}
}

// MathBGColor sets background color.
func MathBGColor(color *color.RGBA) MathOption {
	return func(m *math) {
		m.bgColor = color
	}
}

// MathFonts sets fonts.
func MathFonts(fonts []string) MathOption {
	return func(m *math) {
		m.fonts = fonts
	}
}

type math struct {
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

// NewMath return a math driver.
func NewMath(opts ...MathOption) captchas.Driver {
	d := &math{
		driver:     &driver{htmlTag: htmlTagIMG},
		height:     80,
		width:      220,
		noiseCount: 0,
	}

	for _, f := range opts {
		f(d)
	}

	d.driver.driver = base64Captcha.NewDriverMath(d.height, d.width, d.noiseCount, d.showLineOptions, d.bgColor, d.fonts)

	return d
}
