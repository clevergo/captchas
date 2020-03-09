// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dbstore

import "fmt"

// Dialect is a database dialect.
type Dialect interface {
	// BindVar bind variable.
	BindVar(i int) string
	// Quote escapes key.
	Quote(key string) string
}

// CommonDialect is a common dialect.
var CommonDialect = commonDialect{}

type commonDialect struct {
}

// BindVar implements Dialect.BindVar.
func (d commonDialect) BindVar(i int) string {
	return "?"
}

// Quote implements Dialect.Quote.
func (d commonDialect) Quote(key string) string {
	return fmt.Sprintf("`%s`", key)
}
