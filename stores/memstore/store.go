// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package memstore

import (
	"context"
	"sync"
	"time"

	"clevergo.tech/captchas"
)

// Option is a function that receives a pointer of store.
type Option func(*Store)

// Expiration sets expiration.
func Expiration(expiration time.Duration) Option {
	return func(s *Store) {
		s.expiration = expiration
	}
}

// GCInterval sets garbage collection .
func GCInterval(interval time.Duration) Option {
	return func(s *Store) {
		s.gcInterval = interval
	}
}

type item struct {
	expiration int64
	answer     string
}

// Store is a memory store.
type Store struct {
	mu         *sync.RWMutex
	expiration time.Duration
	gcInterval time.Duration
	items      map[string]*item
}

var _ captchas.Store = New()

// New returns a memory store.
func New(opts ...Option) *Store {
	s := &Store{
		mu:         &sync.RWMutex{},
		expiration: 10 * time.Minute,
		gcInterval: time.Minute,
		items:      make(map[string]*item),
	}

	for _, f := range opts {
		f(s)
	}

	go s.gc()

	return s
}

// Get implements Store.Get.
func (s *Store) Get(ctx context.Context, id string, clear bool) (string, error) {
	if clear {
		item, err := s.getAndDel(id)
		if err != nil {
			return "", err
		}
		return item.answer, nil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	item, err := s.get(id)
	if err != nil {
		return "", err
	}
	return item.answer, nil
}

func (s *Store) get(id string) (*item, error) {
	item, ok := s.items[id]
	if !ok {
		return nil, captchas.ErrCaptchaIncorrect
	}
	if time.Now().UnixNano() > item.expiration {
		return nil, captchas.ErrCaptchaExpired
	}

	return item, nil
}

func (s *Store) getAndDel(id string) (*item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, err := s.get(id)
	if err != nil {
		return nil, err
	}

	delete(s.items, id)

	return item, nil
}

// Set implements Store.Set.
func (s *Store) Set(ctx context.Context, id, answer string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[id] = &item{
		expiration: time.Now().Add(s.expiration).UnixNano(),
		answer:     answer,
	}
	return nil
}

func (s *Store) gc() {
	ticker := time.NewTicker(s.gcInterval)
	for {
		select {
		case <-ticker.C:
			s.deleteExpired()
		}
	}
}

func (s *Store) deleteExpired() {
	now := time.Now().UnixNano()
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, item := range s.items {
		if now > item.expiration {
			delete(s.items, id)
		}
	}
}
