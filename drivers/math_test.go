// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"image/color"
	"reflect"
	"testing"
)

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

func TestMathFonts(t *testing.T) {
	m := &math{}
	fonts := []string{"wqy_microhei.ttc"}
	MathFonts(fonts)(m)
	if !reflect.DeepEqual(fonts, m.fonts) {
		t.Errorf("expected fonts %v, got %v", fonts, m.fonts)
	}
}

func TestMathNoiseCount(t *testing.T) {
	m := &math{}
	MathNoiseCount(4)(m)
	if m.noiseCount != 4 {
		t.Errorf("expected noise count %d, got %d", 4, m.noiseCount)
	}
}

func TestMathBGColor(t *testing.T) {
	m := &math{}
	color := &color.RGBA{1, 2, 3, 4}
	MathBGColor(color)(m)
	if !reflect.DeepEqual(color, m.bgColor) {
		t.Errorf("expected background color %v, got %v", color, m.bgColor)
	}
}

func TestNewMath(t *testing.T) {
	d := NewMath(
		MathHeight(3),
	)

	m, _ := d.(*math)
	if m.height != 3 {
		t.Errorf("expected height %d, got %d", 3, m.height)
	}
}
