// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"image/color"

	"github.com/clevergo/captchas"
	"github.com/mojocn/base64Captcha"
)

// ChineseOption is a function that receives a pointer of chinese driver.
type ChineseOption func(*chinese)

// ChineseHeight sets height.
func ChineseHeight(height int) ChineseOption {
	return func(c *chinese) {
		c.height = height
	}
}

// ChineseWidth sets width.
func ChineseWidth(width int) ChineseOption {
	return func(c *chinese) {
		c.width = width
	}
}

// ChineseLength sets length.
func ChineseLength(length int) ChineseOption {
	return func(c *chinese) {
		c.length = length
	}
}

// ChineseSource sets source.
func ChineseSource(source string) ChineseOption {
	return func(c *chinese) {
		c.source = source
	}
}

// ChineseNoiseCount sets noise count.
func ChineseNoiseCount(count int) ChineseOption {
	return func(c *chinese) {
		c.noiseCount = count
	}
}

// ChineseBGColor sets background color.
func ChineseBGColor(color *color.RGBA) ChineseOption {
	return func(c *chinese) {
		c.bgColor = color
	}
}

// ChineseFonts sets fonts.
func ChineseFonts(fonts []string) ChineseOption {
	return func(c *chinese) {
		c.fonts = fonts
	}
}

type chinese struct {
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

const defaultChineseSource = "零一二三四五六七八九十"

// NewChinese returns a chinese driver.
func NewChinese(opts ...ChineseOption) captchas.Driver {
	d := &chinese{
		driver:     &driver{htmlTag: htmlTagIMG},
		height:     80,
		width:      220,
		noiseCount: 0,
		length:     4,
		source:     defaultChineseSource,
		fonts:      []string{"wqy-microhei.ttc"},
	}

	for _, f := range opts {
		f(d)
	}

	d.driver.driver = base64Captcha.NewDriverChinese(d.height, d.width, d.noiseCount, d.showLineOptions, d.length, d.source, d.bgColor, d.fonts)

	return d
}
