package tree_test

import (
	"testing"

	"github.com/weisrc/recordb/tree"
)

func TestSetMerge(t *testing.T) {
	root := tree.Root[int]()
	root.Set(tree.Hash("a.example.com."), 0, true)
	root.Set(tree.Hash("b.example.com."), 0, true)
	root.Set(tree.Hash("c.example.com."), 0, true)
	Assert(t, root, `
	.example.com.
    |a!*
    |b!*
    |c!*
	`)
}

func TestSetFork(t *testing.T) {
	root := tree.Root[int]()
	root.Set(tree.Hash("example.com."), 0, true)
	root.Set(tree.Hash("com."), 0, true)
	Assert(t, root, `
	com.!*
    |example.!*
	`)
}

func TestSetExact(t *testing.T) {
	root := tree.Root[int]()
	root.Set(tree.Hash("example.com."), 0, true)
	root.Set(tree.Hash("example.com."), 0, true)
	Assert(t, root, `
	example.com.!*
	`)
}

func TestSetExtend(t *testing.T) {
	root := tree.Root[int]()
	root.Set(tree.Hash("example.com."), 0, true)
	root.Set(tree.Hash("a.example.com."), 0, true)
	root.Set(tree.Hash("aa.example.com."), 0, true)
	Assert(t, root, `
	example.com.!*
	|a.!*
	||a!*
	`)
}
