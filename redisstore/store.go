// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package redisstore

import (
	"fmt"
	"time"

	"github.com/clevergo/captchas"
	"github.com/go-redis/redis/v7"
)

// Option is a function that receives a pointer of redis store.
type Option func(s *store)

// Prefix sets the prefix of key.
func Prefix(prefix string) Option {
	return func(s *store) {
		s.prefix = prefix
	}
}

type store struct {
	client     *redis.Client
	expiration time.Duration
	prefix     string
}

// New returns a redis store.
func New(client *redis.Client, expiration time.Duration, opts ...Option) captchas.Store {
	s := &store{
		client:     client,
		prefix:     "captchas",
		expiration: expiration,
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
	tx := s.client.TxPipeline()
	get := tx.Get(key)
	var del *redis.IntCmd
	if clear {
		del = tx.Del(key)
	}
	_, err := tx.Exec()
	if err != nil {
		return "", err
	}
	val, err := get.Result()
	if err != nil {
		return "", fmt.Errorf("failed to get key: %s", key)
	}

	if clear {
		if _, err = del.Result(); err != nil {
			return "", fmt.Errorf("failed to delete key: %s", key)
		}
	}

	return val, nil
}

// Set implements Store.Set.
func (s *store) Set(id string, value string) error {
	key := s.getKey(id)
	_, err := s.client.Set(key, value, s.expiration).Result()
	if err != nil {
		return fmt.Errorf("failed to set key: %s", key)
	}
	return nil
}
