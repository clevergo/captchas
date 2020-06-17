// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package mysqlstore

import (
	"database/sql"

	"clevergo.tech/captchas"
	"clevergo.tech/captchas/stores/dbstore"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

// Store is a MySQL store.
type Store struct {
	*dbstore.Store
}

var _ captchas.Store = New(nil)

// New returns store instance.
func New(db *sql.DB, opts ...dbstore.Option) *Store {
	return &Store{dbstore.New(db, goqu.Dialect("mysql"), opts...)}
}
