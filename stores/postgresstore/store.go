// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package postgresstore

import (
	"database/sql"
	"fmt"

	"clevergo.tech/captchas/stores/dbstore"
)

// Store is a PostgreSQL store.
type Store struct {
	*dbstore.Store
}

// New returns store instance.
func New(db *sql.DB, opts ...dbstore.Option) *Store {
	return &Store{dbstore.New(db, &dialect{}, opts...)}
}

type dialect struct {
}

// BindVar implements Dialect.BindVar.
func (d dialect) BindVar(i int) string {
	return fmt.Sprintf("$%d", i)
}

// Quote implements Dialect.Quote.
func (d dialect) Quote(key string) string {
	return fmt.Sprintf(`"%s"`, key)
}
