# passhash

The package `passhash` implements password hashing.

```go
package main

import (
	"fmt"

	"github.com/ricoberger/gocommon/passhash"
)

func main() {
	password := "password"

	hash, err := passhash.HashString(password)
	if err != nil {
		return
	}

	fmt.Printf("Password: '%s'\nHash: '%s'", password, hash)

	if passhash.MatchString(hash, password) {
		fmt.Printf("The password matchs the hash")
	} else {
		fmt.Printf("The password did not match the hash")
	}
}
```
