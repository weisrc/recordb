package tree_test

import (
	"testing"

	"github.com/weisrc/recordb/tree"
)

func TestFind(t *testing.T) {
	root := tree.Empty("google.com.")
	root.Set("google.ca.", "hello", true)
	root.Set("d.google.ca.", "nope", true)
	root.PrettyPrint(0)
	got := root.Find("ad.google.ca.")
	if got == nil {
		t.Error(`expected node but got nil`)
	} else if got.Value != "hello" {
		t.Errorf(`want "hello", got %q`, got.Value)
	}
}
