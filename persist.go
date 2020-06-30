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
	"reflect"
)

var (
	// The db file to be used for reading and storing the key value pairs.
	db *os.File
	// The decoded map containing the key-value pairs in the db file.
	store = map[string]interface{}{}
)

// Opens the store file with the specified path.
// If the `path` does not exist it will be created otherwise
// the existing file will be used.
// The connection must be closed using `persist.Close()`.
func Open(path string) error {
	var err error
	db, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(db)
	if err := decoder.Decode(&store); err != nil && err != io.EOF {
		return err
	}
	return err
}

// Persists a new key value pair in the store file.
// The value must be a premitive type (string, bool, int, etc.).
func Put(key string, value interface{}) error {
	store[key] = value
	// Clear the file content for the new store map.
	// This could be optimized but for the simple use cases of this package
	// it does not make a huge difference.
	if _, err := db.Seek(0, 0); err != nil {
		return err
	}
	if err := db.Truncate(0); err != nil {
		return err
	}
	// Encode the new store map and write it to the db file.
	encoder := gob.NewEncoder(db)
	return encoder.Encode(store)
}

// Gets the value with the specified key from the store.
// The value must be a pointer to a premitive value (string, int, bool, etc.)
// otherwise an error will be returned.
func Get(key string, value interface{}) error {
	if v, ok := store[key]; ok {
		val := reflect.ValueOf(value)
		// Check if the value is not a pointer.
		if val.Kind() != reflect.Ptr {
			return fmt.Errorf("persist.Get: value must be a pointer")
		}
		// Set the value from the store map to the value pointer.
		val.Elem().Set(reflect.ValueOf(v))
		return nil
	}
	return fmt.Errorf("No value is stored with the specified key")
}

// Closes the store, rendering it usable for I/O by other connections.
func Close() error {
	if db != nil {
		return db.Close()
	}
	return fmt.Errorf("There is no open connection")
}

// Ensures the function runs without an error, otherwise it panics.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
