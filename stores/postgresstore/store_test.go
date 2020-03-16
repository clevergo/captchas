// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package postgresstore

import (
	"fmt"
	"testing"
)

func TestDialectBindVar(t *testing.T) {
	d := dialect{}
	for i := 0; i < 5; i++ {
		v := d.BindVar(i)
		expected := fmt.Sprintf("$%d", i)
		if v != expected {
			t.Errorf("expected %s, got %s", expected, v)
		}
	}
}

func TestDialectQuote(t *testing.T) {
	d := dialect{}
	tests := []struct {
		key      string
		expected string
	}{
		{"foo", `"foo"`},
		{"bar", `"bar"`},
	}
	for _, test := range tests {
		v := d.Quote(test.key)
		if v != test.expected {
			t.Errorf("expected %s, got %s", test.expected, v)
		}
	}
}
