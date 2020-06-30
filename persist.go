// Copyright 2020 Mohammed Salman. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
  Package "persist" is a key value pair storage for simple use cases, it can
  store premitive values only such as strings, integers and booleans.
  "persist" uses a file to read and store it's values, and for putting new
  values it clears the whole file and writes the key-value pairs with the
  newly added pair.
*/
package persist

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

type Store struct {
	// The db file path
	path string
	// The db file to be used for reading and storing the key value pairs.
	db *os.File
	// The decoded map containing the key-value pairs in the db file.
	store map[string]interface{}
}

// Creates and new store.
func NewStore(path string) *Store {
	return &Store{path: path, store: map[string]interface{}{}}
}

// Opens the store file with the specified path.
// If the `path` does not exist it will be created otherwise
// the existing file will be used.
// The connection must be closed using `persist.Close()`.
func (s *Store) Open() error {
	db, err := os.OpenFile(s.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	s.db = db
	decoder := gob.NewDecoder(db)
	if err := decoder.Decode(&s.store); err != nil && err != io.EOF {
		return err
	}
	return nil
}

// Persists a new key value pair in the store file.
// The value must be a premitive type (string, bool, int, etc.).
func (s *Store) Put(key string, value interface{}) error {
	if s.db == nil {
		return fmt.Errorf("db is not opened")
	}
	s.store[key] = value
	// Clear the file content for the new store map.
	// This could be optimized but for the simple use cases of this package
	// it does not make a huge difference.
	if _, err := s.db.Seek(0, 0); err != nil {
		return err
	}
	if err := s.db.Truncate(0); err != nil {
		return err
	}
	// Encode the new store map and write it to the db file.
	encoder := gob.NewEncoder(s.db)
	return encoder.Encode(s.store)
}

// Gets the value with the specified key from the store.
func (s Store) Get(key string) (interface{}, error) {
	if v, ok := s.store[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("No value is stored with the specified key")
}

// Closes the store, rendering it usable for I/O by other connections.
func (s *Store) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return fmt.Errorf("There is no open connection")
}

// Ensures the function runs without an error, otherwise it panics.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
