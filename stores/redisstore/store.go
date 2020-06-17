// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package redisstore

import (
	"context"
	"time"

	"clevergo.tech/captchas"
	"github.com/go-redis/redis/v7"
)

// Option is a function that receives a pointer of redis store.
type Option func(s *Store)

// Prefix sets the prefix of key.
func Prefix(prefix string) Option {
	return func(s *Store) {
		s.prefix = prefix
	}
}

// Expiration sets the expiration.
func Expiration(expiration time.Duration) Option {
	return func(s *Store) {
		s.expiration = expiration
	}
}

// Store is a redis store.
type Store struct {
	client     *redis.Client
	expiration time.Duration
	prefix     string
}

var _ captchas.Store = New(nil)

// New returns a redis store.
func New(client *redis.Client, opts ...Option) *Store {
	s := &Store{
		client:     client,
		prefix:     "captchas",
		expiration: 10 * time.Minute,
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
	tx := s.client.WithContext(ctx).TxPipeline()
	get := tx.Get(key)
	var del *redis.IntCmd
	if clear {
		del = tx.Del(key)
	}
	_, err := tx.Exec()
	if err != nil {
		return "", s.handleError(err)
	}
	val, err := get.Result()
	isNil := false
	if err != nil {
		isNil = err == redis.Nil
		if err == redis.Nil {
			return "", captchas.ErrCaptchaIncorrect
		}
		return "", s.handleError(err)
	}

	if clear && !isNil {
		if _, err = del.Result(); err != nil {
			return "", err
		}
	}

	return val, nil
}

// Set implements Store.Set.
func (s *Store) Set(ctx context.Context, id string, value string) error {
	key := s.getKey(id)
	_, err := s.client.WithContext(ctx).Set(key, value, s.expiration).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) handleError(err error) error {
	if err == redis.Nil {
		return captchas.ErrCaptchaIncorrect
	}

	return err
}
