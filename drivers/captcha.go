// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/mojocn/base64Captcha"
)

var imgTmpl = template.Must(template.New("img").Parse(`<img {{ . }} />`))
var audioTmpl = template.Must(template.New("audio").Parse(`<audio controls {{ . }} />`))

const (
	htmlTagIMG   = "img"
	htmlTagAudio = "audio"
)

type captcha struct {
	id     string
	answer string
	item   base64Captcha.Item
	tag    string
}

func newCaptcha(id, answer, tag string, item base64Captcha.Item) *captcha {
	return &captcha{
		id:     id,
		answer: answer,
		item:   item,
		tag:    tag,
	}
}

// ID implements Captcha.ID.
func (c *captcha) ID() string {
	return c.id
}

// ID implements Captcha.Answer.
func (c *captcha) Answer() string {
	return c.answer
}

// ID implements Captcha.EncodeToString.
func (c *captcha) EncodeToString() string {
	return c.item.EncodeB64string()
}

// ID implements Captcha.HTMLField.
func (c *captcha) HTMLField() template.HTML {
	return c.InputField() + c.MediaField()
}

func (c *captcha) MediaField() template.HTML {
	if c.tag == "audio" {
		return c.audio()
	}

	return c.img()
}

const (
	fieldName = "captcha_id"
)

func (c *captcha) InputField() template.HTML {
	return template.HTML(fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`, fieldName, c.id))
}

func (c *captcha) img() template.HTML {
	buf := &bytes.Buffer{}
	imgTmpl.Execute(buf, template.HTMLAttr(fmt.Sprintf(`src="%s"`, c.item.EncodeB64string())))
	return template.HTML(buf.String())
}

func (c *captcha) audio() template.HTML {
	buf := &bytes.Buffer{}
	audioTmpl.Execute(buf, template.HTMLAttr(fmt.Sprintf(`src="%s"`, c.item.EncodeB64string())))
	return template.HTML(buf.String())
}
