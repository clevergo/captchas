// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package memcachedstore

import (
	"context"

	"clevergo.tech/captchas"
	"github.com/bradfitz/gomemcache/memcache"
)

// Option is a function that receives a pointer of memcached store.
type Option func(s *Store)

// Prefix sets the prefix of key.
func Prefix(prefix string) Option {
	return func(s *Store) {
		s.prefix = prefix
	}
}

// Expiration sets the expiration of captcha.
func Expiration(expiration int32) Option {
	return func(s *Store) {
		s.expiration = expiration
	}
}

// Store is a memcached store.
type Store struct {
	client     *memcache.Client
	prefix     string
	expiration int32
}

var _ captchas.Store = New(nil)

// New returns a memcache store
func New(client *memcache.Client, opts ...Option) *Store {
	s := &Store{
		client:     client,
		prefix:     "captchas",
		expiration: 600,
	}

	for _, f := range opts {
		f(s)
	}

	return s
}

func (s *Store) getKey(id string) string {
	return s.prefix + ":" + id
}

// Get implements Store.Get.
func (s *Store) Get(ctx context.Context, id string, clear bool) (string, error) {
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
func (s *Store) Set(ctx context.Context, id, answer string) error {
	item := &memcache.Item{
		Key:        s.getKey(id),
		Value:      []byte(answer),
		Expiration: s.expiration,
	}

	return s.client.Set(item)
}
