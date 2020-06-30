// Copyright 2020 Mohammed Salman. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package persist

import (
	"fmt"
	"os"
	"testing"
)

func TestPut(t *testing.T) {
	s := NewStore("testdb.db")
	if err := s.Open(); err != nil {
		t.Errorf("Failed to create/open db file %s. %s", s.path, err)
	}
	defer s.Close()
	defer os.Remove(s.path)
	if err := s.Put("key", "value"); err != nil {
		t.Errorf("Failed to put to db %s. %s", s.path, err)
	}

	key := "key"
	got, err := s.Get(key)
	if err != nil {
		t.Errorf("Failed to get key \"%s\" from db %s. %s", key, s.path, err)
		return
	}
	want := "value"
	if got != want {
		t.Errorf("value = \"%s\"; wanted \"%s\"", got, want)
	}
}

func TestGet(t *testing.T) {
	s := NewStore("testdb.db")
	if err := s.Open(); err != nil {
		t.Errorf("Failed to create/open db file %s. %s", s.path, err)
	}
	defer s.Close()
	defer os.Remove(s.path)

	key := "key"
	_, err := s.Get(key)
	if err == nil {
		t.Errorf("Get should return an error for \"%s\" \"%s\"", key, s.path)
		return
	}
}

func Example() {
	// Create a new Store
	s := NewStore("mydb.db")
	// Open db file
	Must(s.Open())
	// Always close the file
	defer s.Close()
	// Store a value with the key "key"
	if err := s.Put("key", "value"); err != nil {
		fmt.Println(err)
	}
	// Get the value
	v, _ := s.Get("key")
	value := v.(string)

	fmt.Println(value) // Output: value
}
