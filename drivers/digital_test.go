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

func TestDigitLength(t *testing.T) {
	d := &digit{}
	DigitLength(4)(d)
	if d.length != 4 {
		t.Errorf("expected length %d, got %d", 4, d.length)
	}
}

func TestDigitMaxSkew(t *testing.T) {
	d := &digit{}
	DigitMaxSkew(0.78)(d)
	if d.maxSkew != 0.78 {
		t.Errorf("expected max skew %f, got %f", 0.78, d.maxSkew)
	}
}

func TestDigitDotCount(t *testing.T) {
	d := &digit{}
	DigitDotCount(4)(d)
	if d.dotCount != 4 {
		t.Errorf("expected dot count %d, got %d", 4, d.dotCount)
	}
}

func TestNewDigit(t *testing.T) {
	d := NewDigit(
		DigitHeight(4),
	)

	digit, _ := d.(*digit)
	if digit.height != 4 {
		t.Errorf("expected height %d, got %d", 4, digit.height)
	}
}
