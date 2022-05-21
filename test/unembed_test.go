package unembed

import (
	"embed"
	"testing"
)

//go:embed README.md LICENSE unembed_test.go unembed.go
var eft embed.FS

func TestUnembed(t *testing.T) {
	err := Unembed(eft, "test")
	if err != nil {
		t.Fatal(err)
	}
}
