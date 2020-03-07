// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"github.com/clevergo/captchas"
	"github.com/mojocn/base64Captcha"
)

// AudioOption is a function that receives a pointer of audio driver.
type AudioOption func(*audio)

// AudioLength sets audio length.
func AudioLength(length int) AudioOption {
	return func(a *audio) {
		a.length = length
	}
}

// AudioLangauge sets audio language.
func AudioLangauge(language string) AudioOption {
	return func(a *audio) {
		a.language = language
	}
}

type audio struct {
	*driver
	// number of digits in captcha solution.
	length int
	// max absolute skew factor of a single audio.
	language string
}

// NewAudio returns an audio driver.
func NewAudio(opts ...AudioOption) captchas.Driver {
	d := &audio{
		driver:   &driver{htmlTag: htmlTagAudio},
		length:   6,
		language: "en",
	}

	for _, f := range opts {
		f(d)
	}

	d.driver.driver = base64Captcha.NewDriverAudio(d.length, d.language)

	return d
}
