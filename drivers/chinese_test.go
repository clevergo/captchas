// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import "testing"

func TestChineseHeight(t *testing.T) {
	c := &chinese{}
	ChineseHeight(4)(c)
	if c.height != 4 {
		t.Errorf("expected height %d, got %d", 4, c.height)
	}
}

func TestChineseWidth(t *testing.T) {
	c := &chinese{}
	ChineseWidth(4)(c)
	if c.width != 4 {
		t.Errorf("expected width %d, got %d", 4, c.width)
	}
}
