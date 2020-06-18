// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package captchas

import (
	"context"
	"errors"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseSensitive(t *testing.T) {
	m := &Manager{}
	CaseSensitive(true)(m)
	assert.True(t, m.caseSensitive)
	CaseSensitive(false)(m)
	assert.False(t, m.caseSensitive)
}

func TestManagerIsEqual(t *testing.T) {
	m := &Manager{}
	assert.False(t, m.isEqual("foo", ""))
	assert.False(t, m.isEqual("", "foo"))
	CaseSensitive(true)(m)
	assert.False(t, m.isEqual("foo", "Foo"))
	CaseSensitive(false)(m)
	assert.True(t, m.isEqual("foo", "Foo"))
}

type testStore struct {
	errGet error
	errSet error
}

func (s *testStore) Get(ctx context.Context, id string, clear bool) (string, error) {
	if s.errGet != nil {
		return "", s.errGet
	}
	if clear {
		return "getAndDel", nil
	}
	return "get", nil
}

func (s *testStore) Set(ctx context.Context, id, answer string) error {
	if s.errSet != nil {
		return s.errSet
	}
	return nil
}

type testDriver struct {
	captcha Captcha
}

func (d *testDriver) Generate() (Captcha, error) {
	if d.captcha != nil {
		return d.captcha, nil
	}
	return nil, errors.New("unsupport to generate captcha")
}

type testCaptcha struct {
	id     string
	answer string
}

func (c *testCaptcha) ID() string {
	return c.id
}

func (c *testCaptcha) Answer() string {
	return c.answer
}

func (c *testCaptcha) EncodeToString() string {
	return ""
}

func (c *testCaptcha) HTMLField(fieldName string) template.HTML {
	return template.HTML("")
}

func TestNew(t *testing.T) {
	store := &testStore{}
	driver := &testDriver{}
	m := New(store, driver, CaseSensitive(false))
	assert.False(t, m.caseSensitive)
	assert.Equal(t, store, m.store)
	assert.Equal(t, driver, m.driver)
}

func TestManagerGenerate(t *testing.T) {
	driver := &testDriver{}
	m := New(&testStore{}, driver)
	captcha1, err1 := m.Generate(nil)
	captcha2, err2 := driver.Generate()
	assert.Equal(t, captcha2, captcha1)
	assert.Equal(t, err2, err1)

	expectedCaptcha := &testCaptcha{id: "foo", answer: "bar"}
	m.driver = &testDriver{captcha: expectedCaptcha}
	captcha, err := m.Generate(nil)
	assert.Nil(t, err)
	assert.Equal(t, expectedCaptcha, captcha)
}

func TestManagerGet(t *testing.T) {
	store := &testStore{}
	m := New(store, &testDriver{})
	for _, clear := range []bool{true, false} {
		val1, err1 := m.Get(context.TODO(), "foo", clear)
		val2, err2 := store.Get(context.TODO(), "foo", clear)
		assert.Equal(t, val2, val1)
		assert.Equal(t, err2, err1)
	}
}

func TestManagerVerify(t *testing.T) {
	store := &testStore{}
	m := New(store, &testDriver{})
	for _, clear := range []bool{true, false} {
		err1 := m.Verify(nil, "foo", "bar", clear)
		assert.Equal(t, ErrCaptchaIncorrect, err1)
	}
	assert.Nil(t, m.Verify(nil, "", "get", false))
	assert.Nil(t, m.Verify(nil, "", "getAndDel", true))

	errGet := errors.New("failed to get answer")
	m.store = &testStore{errGet: errGet}
	assert.Equal(t, errGet, m.Verify(nil, "", "", false))
}
