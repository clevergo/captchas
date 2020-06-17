// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dbstore

import (
	"context"
	"database/sql"
	"log"
	"time"

	"clevergo.tech/captchas"
	"github.com/doug-martin/goqu/v9"
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
	dialect    goqu.DialectWrapper
	tableName  string
	category   string
	expiration time.Duration
	gcInterval time.Duration
}

var _ captchas.Store = New(nil, goqu.DialectWrapper{})

// New returns a db store.
func New(db *sql.DB, dialect goqu.DialectWrapper, opts ...Option) *Store {
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
func (s *Store) Get(ctx context.Context, id string, clear bool) (string, error) {
	query, args, err := s.dialect.From(s.tableName).
		Select("id", "answer", "expires_in").
		Where(goqu.Ex{
			"id":       id,
			"category": s.category,
		}).ToSQL()
	if err != nil {
		return "", err
	}
	row := s.db.QueryRowContext(ctx, query, args...)
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
		query, args, err = s.dialect.Delete(s.tableName).
			Where(goqu.Ex{
				"id":       id,
				"category": s.category,
			}).ToSQL()
		if err != nil {
			return "", err
		}
		if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
			return "", err
		}
	}

	return item.Answer, nil
}

// Set implements Store.Set.
func (s *Store) Set(ctx context.Context, id, answer string) error {
	now := time.Now()
	query, args, err := s.dialect.Insert(s.tableName).Rows(
		goqu.Record{
			"id":         id,
			"category":   s.category,
			"answer":     answer,
			"created_at": now.Unix(),
			"expires_in": now.Add(s.expiration).Unix(),
		},
	).ToSQL()
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, query, args...)
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
	query, args, err := s.dialect.Delete(s.tableName).Where(
		goqu.Ex{
			"category":   s.category,
			"expires_in": goqu.Op{"lt": time.Now().Unix()},
		},
	).ToSQL()
	if err != nil {
		log.Println(err)
	}
	if _, err := s.db.Exec(query, args...); err != nil {
		log.Println(err)
	}
}
