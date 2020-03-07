// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import (
	"image/color"
	"reflect"
	"testing"
)

func TestStringHeight(t *testing.T) {
	s := &str{}
	StringHeight(4)(s)
	if s.height != 4 {
		t.Errorf("expected height %d, got %d", 4, s.height)
	}
}

func TestStringWidth(t *testing.T) {
	s := &str{}
	StringWidth(4)(s)
	if s.width != 4 {
		t.Errorf("expected width %d, got %d", 4, s.width)
	}
}

func TestStringLength(t *testing.T) {
	s := &str{}
	StringLength(4)(s)
	if s.length != 4 {
		t.Errorf("expected length %d, got %d", 4, s.length)
	}
}

func TestStringFonts(t *testing.T) {
	s := &str{}
	fonts := []string{"wqy_microhei.ttc"}
	StringFonts(fonts)(s)
	if !reflect.DeepEqual(fonts, s.fonts) {
		t.Errorf("expected fonts %v, got %v", fonts, s.fonts)
	}
}

func TestStringSource(t *testing.T) {
	s := &str{}
	source := "foobar"
	StringSource(source)(s)
	if s.source != source {
		t.Errorf("expected source %s, got %s", source, s.source)
	}
}

func TestNoiseCount(t *testing.T) {
	s := &str{}
	StringNoiseCount(4)(s)
	if s.noiseCount != 4 {
		t.Errorf("expected noise count %d, got %d", 4, s.noiseCount)
	}
}

func TestStringBGColor(t *testing.T) {
	s := &str{}
	color := &color.RGBA{1, 2, 3, 4}
	StringBGColor(color)(s)
	if !reflect.DeepEqual(color, s.bgColor) {
		t.Errorf("expected background color %v, got %v", color, s.bgColor)
	}
}

func TestNewString(t *testing.T) {
	d := NewString(
		StringLength(3),
	)

	s, _ := d.(*str)
	if s.length != 3 {
		t.Errorf("expected length %d, got %d", 3, s.length)
	}
}
