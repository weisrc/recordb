package tree

import (
	"fmt"
	"strings"
)

func (t *Tree[T]) String() string {
	var sb strings.Builder
	t.string(&sb, 0)
	return sb.String()
}

func (t *Tree[T]) string(sb *strings.Builder, tab int) {
	leaf := ""
	wild := ""
	if t.leaf {
		leaf = "!"
	}
	if t.Wild {
		wild = "*"
	}
	sb.WriteString(fmt.Sprintf("%s%s%s%s\n", strings.Repeat("|", tab), t.segment, leaf, wild))
	for _, v := range t.children {
		if v != nil {
			v.string(sb, tab+1)
		}
	}
}
