// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package memstore

import (
	"sync"
	"time"

	"github.com/clevergo/captchas"
)

type item struct {
	expiration int64
	answer     string
}

type store struct {
	mu         *sync.RWMutex
	expiration time.Duration
	gcInterval time.Duration
	items      map[string]*item
}

// New returns a memory store.
func New(expiration time.Duration, gcInterval time.Duration) captchas.Store {
	s := &store{
		mu:         &sync.RWMutex{},
		expiration: expiration,
		gcInterval: gcInterval,
		items:      make(map[string]*item),
	}

	go s.gc()

	return s
}

// Get implements Store.Get.
func (s *store) Get(id string, clear bool) (string, error) {
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

func (s *store) get(id string) (*item, error) {
	item, ok := s.items[id]
	if !ok {
		return nil, captchas.ErrIncorrectCaptcha
	}
	if time.Now().UnixNano() > item.expiration {
		return nil, captchas.ErrExpiredCaptcha
	}

	return item, nil
}

func (s *store) getAndDel(id string) (*item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, err := s.get(id)
	if err != nil {
		return nil, err
	}

	delete(s.items, id)

	return item, err
}

// Set implements Store.Set.
func (s *store) Set(id, answer string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[id] = &item{
		expiration: time.Now().Add(s.expiration).UnixNano(),
		answer:     answer,
	}
	return nil
}

func (s *store) gc() {
	ticker := time.NewTicker(s.gcInterval)
	for {
		select {
		case <-ticker.C:
			s.deleteExpired()
		}
	}
}

func (s *store) deleteExpired() {
	now := time.Now().UnixNano()
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, item := range s.items {
		if now > item.expiration {
			delete(s.items, id)
		}
	}
}
