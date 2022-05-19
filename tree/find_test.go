package tree_test

import (
	"testing"

	"github.com/weisrc/recordb/tree"
)

func TestFind(t *testing.T) {
	root := tree.Empty[string]("google.com.")
	root.Set("google.ca.", "asdf", true)
	root.Set("d.google.ca.", "asdf", true)
	root.PrettyPrint(0)
	got, ok := root.Get("ad.google.ca.", true)
	if !ok {
		t.Error(`expected node but got nil`)
	} else if got != "hello" {
		t.Errorf(`want "hello", got %q`, got)
	}
}
