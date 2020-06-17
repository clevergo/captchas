// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package mysqlstore

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	psw := os.Getenv("MYSQL_PASSWORD")
	var err error
	testDB, err = sql.Open("mysql", fmt.Sprintf("root:%s@tcp(localhost:3306)/test?multiStatements=true", psw))
	if err != nil {
		panic(err)
	}

	m.Run()
}

func TestStoreGet(t *testing.T) {
	s := New(testDB)
	ctx := context.TODO()
	_, err := s.Get(ctx, "foo", true)
	if err == nil {
		t.Error("expected a non-nil error, got nil")
	}

	err = s.Set(ctx, "foo", "bar")
	if err != nil {
		t.Fatalf("failed to set: %s", err)
	}
	for _, clear := range []bool{false, true} {
		value, err := s.Get(ctx, "foo", clear)
		if err != nil {
			t.Fatalf("expected non error, got %s", err)
		}
		if value != "bar" {
			t.Errorf("expected value %q, got %q", "bar", value)
		}
	}

	_, err = s.Get(ctx, "foo", true)
	if err == nil {
		t.Error("expected a non-nil error, got nil")
	}
}

func TestStoreSet(t *testing.T) {
	// TBD
}
