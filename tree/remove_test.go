package tree_test

import (
	"testing"

	"github.com/weisrc/recordb/tree"
)

func TestRemoveBasic(t *testing.T) {
	root := tree.Root[int]()
	root.Set(tree.Hash("a.example.com."), 0, true)
	root.Set(tree.Hash("b.example.com."), 0, true)
	root.Set(tree.Hash("c.example.com."), 0, true)
	root.Remove(tree.Hash("b.example.com."))
	Assert(t, root, `
	.example.com.
    |a!*
    |c!*
	`)
}

func TestRemoveRoot(t *testing.T) {
	root := tree.Root[int]()
	root.Set(tree.Hash("example.com."), 0, true)
	root.Remove(tree.Hash("example.com."))
	Assert(t, root, "")
}

func TestRemoveInexistent(t *testing.T) {
	root := tree.Root[int]()
	root.Set(tree.Hash("example.com."), 0, true)
	root.Remove(tree.Hash("a.example.com."))
	Assert(t, root, "example.com.!*")
}

func TestRemoveMerge(t *testing.T) {
	root := tree.Root[int]()
	root.Set(tree.Hash("a.example.com."), 0, true)
	root.Set(tree.Hash("b.example.com."), 0, true)
	root.Remove(tree.Hash("b.example.com."))
	Assert(t, root, `
	a.example.com.!*
	`)
}

func TestRemoveMergeNegative(t *testing.T) {
	root := tree.Root[int]()
	root.Set(tree.Hash("a.example.com."), 0, true)
	root.Set(tree.Hash("b.example.com."), 0, true)
	root.Set(tree.Hash("c.example.com."), 0, true)
	root.Remove(tree.Hash("c.example.com."))
	Assert(t, root, `
	.example.com.
	|a!*
	|b!*
	`)
}
