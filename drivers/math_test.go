// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package drivers

import (
	"image/color"
	"reflect"
	"testing"
)

func TestMathHeight(t *testing.T) {
	m := &Math{}
	MathHeight(4)(m)
	if m.height != 4 {
		t.Errorf("expected height %d, got %d", 4, m.height)
	}
}

func TestMathWidth(t *testing.T) {
	m := &Math{}
	MathWidth(4)(m)
	if m.width != 4 {
		t.Errorf("expected width %d, got %d", 4, m.width)
	}
}

func TestMathFonts(t *testing.T) {
	m := &Math{}
	fonts := []string{"wqy_microhei.ttc"}
	MathFonts(fonts)(m)
	if !reflect.DeepEqual(fonts, m.fonts) {
		t.Errorf("expected fonts %v, got %v", fonts, m.fonts)
	}
}

func TestMathNoiseCount(t *testing.T) {
	m := &Math{}
	MathNoiseCount(4)(m)
	if m.noiseCount != 4 {
		t.Errorf("expected noise count %d, got %d", 4, m.noiseCount)
	}
}

func TestMathBGColor(t *testing.T) {
	m := &Math{}
	color := &color.RGBA{1, 2, 3, 4}
	MathBGColor(color)(m)
	if !reflect.DeepEqual(color, m.bgColor) {
		t.Errorf("expected background color %v, got %v", color, m.bgColor)
	}
}

func TestNewMath(t *testing.T) {
	m := NewMath(
		MathHeight(3),
	)

	if m.height != 3 {
		t.Errorf("expected height %d, got %d", 3, m.height)
	}
}
