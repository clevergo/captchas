// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import "testing"

func TestStringHeight(t *testing.T) {
	s := &str{}
	StringHeight(4)(s)
	if s.height != 4 {
		t.Errorf("expected height %d, got %d", 4, s.height)
	}
}

func TestStringWidth(t *testing.T) {
	d := &str{}
	StringWidth(4)(d)
	if d.width != 4 {
		t.Errorf("expected width %d, got %d", 4, d.width)
	}
}
