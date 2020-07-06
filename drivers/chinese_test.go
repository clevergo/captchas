// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package drivers

import (
	"image/color"
	"reflect"
	"testing"
)

func TestChineseHeight(t *testing.T) {
	c := &Chinese{}
	ChineseHeight(4)(c)
	if c.height != 4 {
		t.Errorf("expected height %d, got %d", 4, c.height)
	}
}

func TestChineseWidth(t *testing.T) {
	c := &Chinese{}
	ChineseWidth(4)(c)
	if c.width != 4 {
		t.Errorf("expected width %d, got %d", 4, c.width)
	}
}

func TestChineseLength(t *testing.T) {
	c := &Chinese{}
	ChineseLength(4)(c)
	if c.length != 4 {
		t.Errorf("expected length %d, got %d", 4, c.length)
	}
}

func TestChineseFonts(t *testing.T) {
	c := &Chinese{}
	fonts := []string{"wqy_microhei.ttc"}
	ChineseFonts(fonts)(c)
	if !reflect.DeepEqual(fonts, c.fonts) {
		t.Errorf("expected fonts %v, got %v", fonts, c.fonts)
	}
}

func TestChineseSource(t *testing.T) {
	c := &Chinese{}
	source := "foobar"
	ChineseSource(source)(c)
	if c.source != source {
		t.Errorf("expected source %s, got %s", source, c.source)
	}
}

func TestChineseNoiseCount(t *testing.T) {
	c := &Chinese{}
	ChineseNoiseCount(4)(c)
	if c.noiseCount != 4 {
		t.Errorf("expected noise count %d, got %d", 4, c.noiseCount)
	}
}

func TestChineseBGColor(t *testing.T) {
	c := &Chinese{}
	color := &color.RGBA{1, 2, 3, 4}
	ChineseBGColor(color)(c)
	if !reflect.DeepEqual(color, c.bgColor) {
		t.Errorf("expected background color %v, got %v", color, c.bgColor)
	}
}

func TestNewChinese(t *testing.T) {
	c := NewChinese(
		ChineseLength(3),
	)

	if c.length != 3 {
		t.Errorf("expected length %d, got %d", 3, c.length)
	}
}
