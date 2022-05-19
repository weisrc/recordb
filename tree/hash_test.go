package tree_test

import (
	"testing"

	"github.com/weisrc/recordb/tree"
)

func TestHashCaseInsentivity(t *testing.T) {
	root := tree.Root[int]()
	root.Set("a.example.com.", 0, true)
	root.Set("b.example.com.", 0, true)
	root.Set("c.example.com.", 0, true)
	root.Set("C.example.com.", 0, true)
	Assert(t, root, `
	.example.com.
    |a!*
    |b!*
    |C!*
	`)
}
