// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package postgresstore

import (
	"database/sql"
	"fmt"

	"github.com/clevergo/captchas/stores/dbstore"
)

type Store struct {
	*dbstore.Store
}

func New(db *sql.DB, opts ...dbstore.Option) *Store {
	return &Store{dbstore.New(db, &dialect{}, opts...)}
}

type dialect struct {
}

func (d dialect) BindVar(i int) string {
	return fmt.Sprintf("$%d", i)
}

func (d dialect) Quote(key string) string {
	return fmt.Sprintf(`"%s"`, key)
}
