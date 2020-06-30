# persist
A simple key value pair storage for simple use cases, it can
store premitive values such as strings, integers and booleans.
"persist" uses a file to read and store it's values, and for putting new
values it clears the whole file and writes the key-value pairs with the
newly added pair.

### Example:

```go
package main

import (
  "fmt"
  
  "github.com/msal4/persist"
)

func main() {
  // Open db file
  if err := persist.Open("path/to/mydb.db")); err != nil {
    panic(err)
  }
  // Always close the file
  defer persist.Close()
  // Store a value with the key "key"
  if err := persist.Put("key", "value"); err != nil {
    fmt.Println(err)
  }
  // Get the value
  var value string
  persist.Get("key", &value)

  fmt.Println(value) // Output: value
}

```

### Note:
This package is meant for small use cases, if you intend to write to it extensively or store large pieces of data, I recommend using something like boltdb.
