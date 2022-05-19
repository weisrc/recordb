package tree_test

import (
	"testing"

	"github.com/weisrc/recordb/tree"
)

func TestRemoveBasic(t *testing.T) {
	root := tree.Root[int]()
	root.Set("a.example.com.", 0, true)
	root.Set("b.example.com.", 0, true)
	root.Set("c.example.com.", 0, true)
	root.Remove("b.example.com.")
	Assert(t, root, `
	.example.com.
    |a!*
    |c!*
	`)
}

func TestRemoveRoot(t *testing.T) {
	root := tree.Root[int]()
	root.Set("example.com.", 0, true)
	root.Remove("example.com.")
	Assert(t, root, "")
}

func TestRemoveInexistent(t *testing.T) {
	root := tree.Root[int]()
	root.Set("example.com.", 0, true)
	root.Remove("a.example.com.")
	Assert(t, root, "example.com.!*")
}

func TestRemoveMerge(t *testing.T) {
	root := tree.Root[int]()
	root.Set("a.example.com.", 0, true)
	root.Set("b.example.com.", 0, true)
	root.Remove("b.example.com.")
	Assert(t, root, `
	a.example.com.!*
	`)
}

func TestRemoveMergeNegative(t *testing.T) {
	root := tree.Root[int]()
	root.Set("a.example.com.", 0, true)
	root.Set("b.example.com.", 0, true)
	root.Set("c.example.com.", 0, true)
	root.Remove("c.example.com.")
	Assert(t, root, `
	.example.com.
	|a!*
	|b!*
	`)
}
