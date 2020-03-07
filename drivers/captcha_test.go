// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
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
