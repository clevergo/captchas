// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import "testing"

func TestDigitHeight(t *testing.T) {
	d := &digit{}
	DigitHeight(4)(d)
	if d.height != 4 {
		t.Errorf("expected height %d, got %d", 4, d.height)
	}
}

func TestDigitWidth(t *testing.T) {
	d := &digit{}
	DigitWidth(4)(d)
	if d.width != 4 {
		t.Errorf("expected width %d, got %d", 4, d.width)
	}
}
