// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dbstore

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/clevergo/captchas"
)

type item struct {
	ID        string `db:"id"`
	Category  string `db:"category"`
	Answer    string `db:"answer"`
	CreatedAt int64  `db:"created_at"`
	ExpiresIn int64  `db:"expires_in"`
}

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

// Category sets captcha category.
func Category(category string) Option {
	return func(s *Store) {
		s.category = category
	}
}

// TableName sets captcha table name.
func TableName(name string) Option {
	return func(s *Store) {
		s.tableName = name
	}
}

// Store is a database store.
type Store struct {
	db         *sql.DB
	dialect    Dialect
	tableName  string
	category   string
	expiration time.Duration
	gcInterval time.Duration
}

// New returns a db store.
func New(db *sql.DB, dialect Dialect, opts ...Option) *Store {
	s := &Store{
		db:         db,
		dialect:    dialect,
		tableName:  "captchas",
		category:   "default",
		expiration: 10 * time.Minute,
		gcInterval: time.Hour,
	}

	for _, f := range opts {
		f(s)
	}

	go s.gc()

	return s
}

// Get implements Store.Get.
func (s *Store) Get(id string, clear bool) (string, error) {
	stat := fmt.Sprintf(
		"SELECT %s, %s, %s FROM %s WHERE id=%s AND category=%s",
		s.dialect.Quote("id"),
		s.dialect.Quote("answer"),
		s.dialect.Quote("expires_in"),
		s.dialect.Quote(s.tableName),
		s.dialect.BindVar(1),
		s.dialect.BindVar(2),
	)
	row := s.db.QueryRow(stat, id, s.category)
	if row == nil {
		return "", captchas.ErrCaptchaIncorrect
	}
	item := item{}
	if err := row.Scan(&item.ID, &item.Answer, &item.ExpiresIn); err != nil {
		if err == sql.ErrNoRows {
			return "", captchas.ErrCaptchaIncorrect
		}
		return "", err
	}
	if time.Now().Unix() > item.ExpiresIn {
		return "", captchas.ErrCaptchaExpired
	}

	if clear {
		stat := fmt.Sprintf(
			"DELETE FROM %s WHERE id=%s AND category=%s",
			s.dialect.Quote(s.tableName),
			s.dialect.BindVar(1),
			s.dialect.BindVar(2),
		)
		_, err := s.db.Exec(stat, id, s.category)
		if err != nil {
			return "", err
		}
	}

	return item.Answer, nil
}

// Set implements Store.Set.
func (s *Store) Set(id, answer string) error {
	now := time.Now()
	stat := fmt.Sprintf(
		"INSERT INTO %s(%s, %s, %s, %s, %s) VALUES(%s, %s, %s, %s, %s)",
		s.dialect.Quote(s.tableName),
		s.dialect.Quote("id"),
		s.dialect.Quote("category"),
		s.dialect.Quote("answer"),
		s.dialect.Quote("created_at"),
		s.dialect.Quote("expires_in"),
		s.dialect.BindVar(1),
		s.dialect.BindVar(2),
		s.dialect.BindVar(3),
		s.dialect.BindVar(4),
		s.dialect.BindVar(5),
	)
	_, err := s.db.Exec(stat, id, s.category, answer, now.Unix(), now.Add(s.expiration).Unix())
	return err
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
	stat := fmt.Sprintf(
		"DELETE FROM %s WHERE category=%s AND expires_in<%s",
		s.dialect.Quote(s.tableName),
		s.dialect.BindVar(1),
		s.dialect.BindVar(2),
	)
	s.db.Exec(stat, s.category, time.Now().Unix())
}
