// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"clevergo.tech/captchas"
	"github.com/mojocn/base64Captcha"
)

type driver struct {
	driver  base64Captcha.Driver
	htmlTag string
}

func (d *driver) Generate() (captchas.Captcha, error) {
	id, question, answer := d.driver.GenerateIdQuestionAnswer()
	item, err := d.driver.DrawCaptcha(question)
	if err != nil {
		return nil, err
	}

	return newCaptcha(id, answer, d.htmlTag, item), nil
}
