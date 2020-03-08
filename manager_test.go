// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package captchas

import (
	"errors"
	"reflect"
	"testing"
)

func TestCaseSensitive(t *testing.T) {
	m := &Manager{}
	CaseSensitive(true)(m)
	if !m.caseSensitive {
		t.Error("expected to enable case sensitive")
	}
	CaseSensitive(false)(m)
	if m.caseSensitive {
		t.Error("expected to disable case sensitive")
	}
}

func TestManagerIsEqual(t *testing.T) {
	m := &Manager{}

	if m.isEqual("", "bar") {
		t.Errorf("expected %t, got %t", false, m.isEqual("", "bar"))
	}
	if m.isEqual("foo", "") {
		t.Errorf("expected %t, got %t", false, m.isEqual("foo", ""))
	}

	CaseSensitive(true)(m)
	if m.isEqual("foo", "Foo") {
		t.Errorf("expected %q not equals %q", "foo", "Foo")
	}

	CaseSensitive(false)(m)
	if !m.isEqual("foo", "Foo") {
		t.Errorf("expected %q equals %q", "foo", "Foo")
	}
}

type testStore struct {
}

func (s *testStore) Get(id string, clear bool) (string, error) {
	if clear {
		return "getAndDel", nil
	}
	return "get", nil
}

func (s *testStore) Set(id, answer string) error {
	return nil
}

type testDriver struct {
}

func (d *testDriver) Generate() (Captcha, error) {
	return nil, errors.New("unsupport to generate captcha")
}

func TestNew(t *testing.T) {
	store := &testStore{}
	driver := &testDriver{}
	m := New(store, driver, CaseSensitive(false))
	if m.caseSensitive {
		t.Error("expected to disable case sensitive")
	}
	if !reflect.DeepEqual(store, m.store) {
		t.Errorf("expected store %v, got %v", store, m.store)
	}
	if !reflect.DeepEqual(driver, m.driver) {
		t.Errorf("expected driver %v, got %v", driver, m.driver)
	}
}

func TestManagerGenerate(t *testing.T) {
	driver := &testDriver{}
	m := New(&testStore{}, driver)
	captcha1, err1 := m.Generate()
	captcha2, err2 := driver.Generate()
	if !reflect.DeepEqual(captcha1, captcha2) {
		t.Errorf("expected captcha %v, got %v", captcha2, captcha1)
	}
	if !reflect.DeepEqual(err1, err2) {
		t.Errorf("expected err %v, got %v", err2, err1)
	}
}

func TestManagerGet(t *testing.T) {
	store := &testStore{}
	m := New(store, &testDriver{})
	for _, clear := range []bool{true, false} {
		val1, err1 := m.Get("foo", clear)
		val2, err2 := store.Get("foo", clear)
		if val1 != val2 {
			t.Errorf("expected value %v, got %v", val2, val1)
		}
		if !reflect.DeepEqual(err1, err2) {
			t.Errorf("expected err %v, got %v", err2, err1)
		}
	}
}

func TestManagerVerify(t *testing.T) {
	store := &testStore{}
	m := New(store, &testDriver{})
	for _, clear := range []bool{true, false} {
		err1 := m.Verify("foo", "bar", clear)
		if !reflect.DeepEqual(err1, ErrIncorrectCaptcha) {
			t.Errorf("expected err %v, got %v", ErrIncorrectCaptcha, err1)
		}
	}
}
