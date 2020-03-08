// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package memcachedstore

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/clevergo/captchas"
)

// Option is a function that receives a pointer of memcached store.
type Option func(s *store)

// Prefix sets the prefix of key.
func Prefix(prefix string) Option {
	return func(s *store) {
		s.prefix = prefix
	}
}

// Expiration sets the expiration of captcha.
func Expiration(expiration int32) Option {
	return func(s *store) {
		s.expiration = expiration
	}
}

type store struct {
	client     *memcache.Client
	prefix     string
	expiration int32
}

// New returns a memcache store
func New(client *memcache.Client, opts ...Option) captchas.Store {
	s := &store{
		client:     client,
		prefix:     "captchas",
		expiration: 600,
	}

	for _, f := range opts {
		f(s)
	}

	return s
}

func (s *store) getKey(id string) string {
	return s.prefix + ":" + id
}

// Get implements Store.Get.
func (s *store) Get(id string, clear bool) (string, error) {
	key := s.getKey(id)
	item, err := s.client.Get(key)
	if err != nil {
		return "", err
	}
	if clear {
		if err := s.client.Delete(key); err != nil {
			return "", err
		}
	}
	return string(item.Value), nil
}

// Set implements Store.Get.
func (s *store) Set(id, answer string) error {
	item := &memcache.Item{
		Key:        s.getKey(id),
		Value:      []byte(answer),
		Expiration: s.expiration,
	}

	return s.client.Set(item)
}
