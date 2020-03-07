// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import "testing"

func TestMathHeight(t *testing.T) {
	m := &math{}
	MathHeight(4)(m)
	if m.height != 4 {
		t.Errorf("expected height %d, got %d", 4, m.height)
	}
}

func TestMathWidth(t *testing.T) {
	m := &math{}
	MathWidth(4)(m)
	if m.width != 4 {
		t.Errorf("expected width %d, got %d", 4, m.width)
	}
}
