# unembed
Take a filesystem from go:embed and unpack it to a directory someplace.

```go
package main

import (
	"embed"
)

//go:embed README.md LICENSE unembed_test.go unembed.go
var eft embed.FS

func main() {
	err := Unembed(eft, "test2")
	if err != nil {
		t.Fatal(err)
	}
}

```
