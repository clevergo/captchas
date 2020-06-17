// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"clevergo.tech/captchas"
	"github.com/mojocn/base64Captcha"
)

// DigitOption is a function that receives a pointer of digit driver.
type DigitOption func(*Digit)

// DigitHeight sets height.
func DigitHeight(height int) DigitOption {
	return func(d *Digit) {
		d.height = height
	}
}

// DigitWidth sets width.
func DigitWidth(width int) DigitOption {
	return func(d *Digit) {
		d.width = width
	}
}

// DigitLength sets length.
func DigitLength(length int) DigitOption {
	return func(d *Digit) {
		d.length = length
	}
}

// DigitMaxSkew sets max skew.
func DigitMaxSkew(maxSkew float64) DigitOption {
	return func(d *Digit) {
		d.maxSkew = maxSkew
	}
}

// DigitDotCount sets dot count.
func DigitDotCount(count int) DigitOption {
	return func(d *Digit) {
		d.dotCount = count
	}
}

// Digit is a digit driver.
type Digit struct {
	*driver
	// captcha png height in pixel.
	height int
	// captcha png width in pixel.
	width int
	// number of digits in captcha solution.
	length int
	// max absolute skew factor of a single digit.
	maxSkew float64
	// number of background circles.
	dotCount int
}

var _ captchas.Driver = NewDigit()

// NewDigit return a digit driver.
func NewDigit(opts ...DigitOption) *Digit {
	d := &Digit{
		driver:   &driver{htmlTag: htmlTagIMG},
		height:   80,
		width:    220,
		length:   6,
		maxSkew:  0.7,
		dotCount: 80,
	}

	for _, f := range opts {
		f(d)
	}

	d.driver.driver = base64Captcha.NewDriverDigit(d.height, d.width, d.length, d.maxSkew, d.dotCount)

	return d
}
