// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dbstore

import "testing"

func TestCommonDialectBindVar(t *testing.T) {
	d := commonDialect{}
	for i := 0; i < 5; i++ {
		v := d.BindVar(i)
		if v != "?" {
			t.Errorf("expected %s, got %s", "?", v)
		}
	}
}

func TestCommonDialectQuote(t *testing.T) {
	d := commonDialect{}
	tests := []struct {
		key      string
		expected string
	}{
		{"foo", "`foo`"},
		{"bar", "`bar`"},
	}
	for _, test := range tests {
		v := d.Quote(test.key)
		if v != test.expected {
			t.Errorf("expected %s, got %s", test.expected, v)
		}
	}
}
