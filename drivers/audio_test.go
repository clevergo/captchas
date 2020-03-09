// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package drivers

import "testing"

func TestAudioLength(t *testing.T) {
	a := &Audio{}
	length := 8
	AudioLength(length)(a)
	if a.length != length {
		t.Errorf("expected length %d, got %d", length, a.length)
	}
}

func TestAudioLanguage(t *testing.T) {
	a := &Audio{}
	language := "zh"
	AudioLangauge(language)(a)
	if a.language != language {
		t.Errorf("expected language %s, got %s", language, a.language)
	}
}

func TestNewAudio(t *testing.T) {
	a := NewAudio(
		AudioLength(4),
		AudioLangauge("zh"),
	)

	if a.length != 4 {
		t.Errorf("expected length %d, got %d", 4, a.length)
	}
	if a.language != "zh" {
		t.Errorf("expected language %s, got %s", "zh", a.language)
	}
}
