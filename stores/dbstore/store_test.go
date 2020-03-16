// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dbstore

import (
	"testing"
	"time"
)

func TestExpiration(t *testing.T) {
	s := &Store{}
	expiration := time.Minute
	Expiration(expiration)(s)
	if s.expiration != expiration {
		t.Errorf("expected expiration %v, got %v", expiration, s.expiration)
	}
}

func TestGCInterval(t *testing.T) {
	s := &Store{}
	interval := time.Minute
	GCInterval(interval)(s)
	if s.gcInterval != interval {
		t.Errorf("expected gc interval %v, got %v", interval, s.gcInterval)
	}
}

func TestCategory(t *testing.T) {
	s := &Store{}
	category := "foo"
	Category(category)(s)
	if s.category != category {
		t.Errorf("expected category %s, got %s", category, s.category)
	}
}

func TestTableName(t *testing.T) {
	s := &Store{}
	tableName := "bar"
	TableName(tableName)(s)
	if s.tableName != tableName {
		t.Errorf("expected table name %s, got %s", tableName, s.tableName)
	}
}
