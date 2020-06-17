// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"image/color"

	"clevergo.tech/captchas"
	"github.com/mojocn/base64Captcha"
)

// StringOption is a function that receives a pointer of string driver.
type StringOption func(*Str)

// StringHeight sets height.
func StringHeight(height int) StringOption {
	return func(s *Str) {
		s.height = height
	}
}

// StringWidth sets width.
func StringWidth(width int) StringOption {
	return func(s *Str) {
		s.width = width
	}
}

// StringLength sets length.
func StringLength(length int) StringOption {
	return func(s *Str) {
		s.length = length
	}
}

// StringSource sets source.
func StringSource(source string) StringOption {
	return func(s *Str) {
		s.source = source
	}
}

// StringNoiseCount sets noise count.
func StringNoiseCount(count int) StringOption {
	return func(s *Str) {
		s.noiseCount = count
	}
}

// StringBGColor sets background color.
func StringBGColor(color *color.RGBA) StringOption {
	return func(s *Str) {
		s.bgColor = color
	}
}

// StringFonts sets fonts.
func StringFonts(fonts []string) StringOption {
	return func(s *Str) {
		s.fonts = fonts
	}
}

// Str is a string driver.
type Str struct {
	*driver
	// captcha png height in pixel.
	height int
	// captcha png width in pixel.
	width  int
	length int
	source string
	// text noise count.
	noiseCount      int
	showLineOptions int
	// background color.
	bgColor *color.RGBA
	fonts   []string
}

const defaultStringSource = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var _ captchas.Driver = NewString()

// NewString returns a string driver.
func NewString(opts ...StringOption) *Str {
	d := &Str{
		driver:     &driver{htmlTag: htmlTagIMG},
		height:     80,
		width:      220,
		noiseCount: 0,
		length:     4,
		source:     defaultStringSource,
	}

	for _, f := range opts {
		f(d)
	}

	d.driver.driver = base64Captcha.NewDriverString(d.height, d.width, d.noiseCount, d.showLineOptions, d.length, d.source, d.bgColor, d.fonts)

	return d
}
