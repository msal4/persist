// Copyright 2020 Mohammed Salman. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package persist

import (
	"fmt"
	"os"
	"testing"
)

var dbname = "testdb.db"

func TestOpen(t *testing.T) {
	if err := Open(dbname); err != nil {
		t.Errorf("Failed to create/open db file %s. %s", dbname, err)
	}
}

func TestPut(t *testing.T) {
	if err := Put("key", "value"); err != nil {
		t.Errorf("Failed to put to db %s. %s", dbname, err)
	}
}

func TestGet(t *testing.T) {
	key := "key"
	var got string
	if err := Get(key, &got); err != nil {
		t.Errorf("Failed to get key \"%s\" from db %s. %s", key, dbname, err)
	}
	want := "value"
	if got != want {
		t.Errorf("value = \"%s\"; wanted \"%s\"", got, want)
	}
}

func TestClose(t *testing.T) {
	if err := Close(); err != nil {
		t.Errorf("Failed to close file %s", dbname)
	}
	if err := os.Remove(dbname); err != nil {
		t.Errorf("Failed to remove file %s", dbname)
	}
}

func Example() {
	// Open db file
	Must(Open("mydb.db"))
	// Always close the file
	defer Close()
	// Store a value with the key "key"
	if err := Put("key", "value"); err != nil {
		fmt.Println(err)
	}
	// Get the value
	var value string
	Get("key", &value)

	fmt.Println(value) // Output: value
}
