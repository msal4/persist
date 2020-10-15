# persist
A key-value pair storage for simple use cases, used to
store premitive values.

### Example:

```go
package main

import (
  "fmt"
  
  "github.com/msal4/persist"
)

func main() {
	// Create a new Store
	s := persist.NewStore("mydb.db")
	// Open db file
	persist.Must(s.Open())
	// Always close the file
	defer s.Close()
	// Store a value with the key "key"
	if err := s.Put("key", "value"); err != nil {
		fmt.Println(err)
	}
	// Get the value
	v, _ := s.Get("key")
	fmt.Println(v) // Output: value
}

```

### Note:
This package is meant for small use cases, if you intend to write to it extensively or store large pieces of data, I recommend using something like boltdb.
