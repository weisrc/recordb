package tree_test

import (
	"strings"
	"testing"

	"github.com/weisrc/recordb/tree"
)

func Assert[T any](t *testing.T, node *tree.Tree[T], s string) {
	sb := strings.Builder{}
	for _, v := range strings.Split(strings.Trim(s, " \t\n"), "\n") {
		sb.WriteString(strings.Trim(v, " \t\n") + "\n")
	}
	want := sb.String()
	got := node.String()
	if want != got {
		t.Errorf("\nGOT:\n%sWANT:\n%s", got, want)
	}
}
