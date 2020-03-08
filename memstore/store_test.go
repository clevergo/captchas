// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package memstore

import (
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	expiration := 10 * time.Minute
	gcInterval := time.Minute
	s := New(expiration, gcInterval)

	mem, _ := s.(*store)
	if mem.expiration != expiration {
		t.Errorf("expected expiration %v, got %v", expiration, mem.expiration)
	}
	if mem.gcInterval != gcInterval {
		t.Errorf("expected gcInterval %v, got %v", gcInterval, mem.gcInterval)
	}
}

func TestStoreGet(t *testing.T) {
	s := New(10*time.Minute, time.Minute)
	_, err := s.Get("foo", true)
	if err == nil {
		t.Error("expected a non-nil error, got nil")
	}

	err = s.Set("foo", "bar")
	if err != nil {
		t.Fatalf("failed to set: %s", err)
	}
	for _, clear := range []bool{false, true} {
		value, err := s.Get("foo", clear)
		if err != nil {
			t.Fatalf("expected non error, got %s", err)
		}
		if value != "bar" {
			t.Errorf("expected value %q, got %q", "bar", value)
		}
	}

	_, err = s.Get("foo", true)
	if err == nil {
		t.Error("expected a non-nil error, got nil")
	}
}

func TestStoreSet(t *testing.T) {
	s := &store{
		mu: &sync.RWMutex{},
		items: map[string]*item{
			"expired": {int64(0), "expired"},
			"active":  {time.Now().Add(time.Second).UnixNano(), "active"},
		},
	}

	s.deleteExpired()
	if len(s.items) != 1 {
		t.Errorf("expected items count %d, got %d", 1, len(s.items))
	}
	if _, ok := s.items["expired"]; ok {
		t.Errorf("expected item %q to be deleted", "expired")
	}

	time.Sleep(time.Second)
	s.deleteExpired()
	if len(s.items) != 0 {
		t.Errorf("expected items count %d, got %d", 0, len(s.items))
	}
}
