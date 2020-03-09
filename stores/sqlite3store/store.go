// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package sqlite3store

import (
	"database/sql"

	"github.com/clevergo/captchas/stores/dbstore"
)

type Store struct {
	*dbstore.Store
}

func New(db *sql.DB, opts ...dbstore.Option) *Store {
	return &Store{dbstore.New(db, dbstore.CommonDialect, opts...)}
}
