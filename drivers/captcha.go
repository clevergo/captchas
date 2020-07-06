// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package drivers

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/mojocn/base64Captcha"
)

var tmplContent = `
<input type="hidden" name="{{ .fieldName }}" value="{{ .captcha.ID }}">
{{ if .captcha.IsTagAudio }}
<audio controls {{ .captcha.MediaAttr }} />
{{ else }}
<img {{ .captcha.MediaAttr }} />
{{ end }}
`
var tmpl = template.Must(template.New("captcha").Parse(tmplContent))

const (
	htmlTagIMG   = "img"
	htmlTagAudio = "audio"
)

// Captcha implements captchas.Captcha interface.
type Captcha struct {
	id     string
	answer string
	item   base64Captcha.Item
	tag    string
}

func newCaptcha(id, answer, tag string, item base64Captcha.Item) *Captcha {
	return &Captcha{
		id:     id,
		answer: answer,
		item:   item,
		tag:    tag,
	}
}

// ID implements Captcha.ID.
func (c *Captcha) ID() string {
	return c.id
}

// Answer implements Captcha.Answer.
func (c *Captcha) Answer() string {
	return c.answer
}

// EncodeToString implements Captcha.EncodeToString.
func (c *Captcha) EncodeToString() string {
	return c.item.EncodeB64string()
}

// HTMLField implements Captcha.HTMLField.
func (c *Captcha) HTMLField(fieldName string) template.HTML {
	buf := &bytes.Buffer{}
	tmpl.Execute(buf, map[string]interface{}{
		"captcha":   c,
		"fieldName": fieldName,
	})
	return template.HTML(buf.String())
}

// MediaAttr returns template.HTMLAttr.
func (c *Captcha) MediaAttr() template.HTMLAttr {
	return template.HTMLAttr(fmt.Sprintf(`src="%s"`, c.item.EncodeB64string()))
}

// IsTagAudio returns a bool value indicates whehter the tag is audio.
func (c *Captcha) IsTagAudio() bool {
	return c.tag == htmlTagAudio
}
