// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"fmt"
	"html/template"
	"strings"
	"testing"

	"github.com/mojocn/base64Captcha"
)

func TestCaptcha(t *testing.T) {
	id := "test"
	answer := "expected"
	tag := htmlTagIMG
	var item base64Captcha.Item
	captcha := newCaptcha(id, answer, tag, item)
	if captcha.ID() != id {
		t.Errorf("expected ID %s, got %s", id, captcha.ID())
	}
	if captcha.Answer() != answer {
		t.Errorf("expected answer %s, got %s", answer, captcha.Answer())
	}
	if captcha.tag != tag {
		t.Errorf("expected tag %s, got %s", tag, captcha.tag)
	}
	if captcha.item != item {
		t.Errorf("expected item %v, got %v", item, captcha.item)
	}
}

func TestCaptchaEncodeToString(t *testing.T) {
	driver := NewDigit()
	c, err := driver.Generate()
	if err != nil {
		t.Fatal(err)
	}

	if c.EncodeToString() != c.item.EncodeB64string() {
		t.Errorf("expected base64 encode string %q, got %q", c.item.EncodeB64string(), c.EncodeToString())
	}
}

func TestCaptchaMediaAttr(t *testing.T) {
	driver := NewDigit()
	c, err := driver.Generate()
	if err != nil {
		t.Fatal(err)
	}

	expected := template.HTMLAttr(fmt.Sprintf(`src="%s"`, c.item.EncodeB64string()))
	if c.MediaAttr() != expected {
		t.Errorf("expected media attr %v, got %v", expected, c.MediaAttr())
	}
}

func TestCaptchaHTMLField(t *testing.T) {
	driver := NewDigit()
	c, err := driver.Generate()
	if err != nil {
		t.Fatal(err)
	}

	fieldName := "captcha_id"
	content := string(c.HTMLField(fieldName))
	if !strings.Contains(content, `name="`+fieldName+`"`) {
		t.Errorf("HTML output doesn't contians input field")
	}
	if !strings.Contains(content, `src="`+c.EncodeToString()+`"`) {
		t.Errorf("HTML output doesn't contians media field")
	}
}

func TestCaptchaIsTagAudio(t *testing.T) {
	imgCaptcha := &Captcha{tag: htmlTagIMG}
	if imgCaptcha.IsTagAudio() {
		t.Errorf("expected %t, got %t", true, imgCaptcha.IsTagAudio())
	}

	audioCaptcha := &Captcha{tag: htmlTagAudio}
	if !audioCaptcha.IsTagAudio() {
		t.Errorf("expected %t, got %t", true, audioCaptcha.IsTagAudio())
	}
}
