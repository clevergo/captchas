// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package mysqlstore

import (
	"database/sql"

	"clevergo.tech/captchas/stores/dbstore"
)

// Store is a MySQL store.
type Store struct {
	*dbstore.Store
}

// New returns store instance.
func New(db *sql.DB, opts ...dbstore.Option) *Store {
	return &Store{dbstore.New(db, dbstore.CommonDialect, opts...)}
}
