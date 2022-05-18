package tree

func (t *Tree) find(path string, bound int, found *Tree) *Tree {
	// fmt.Printf("[%q]	Traverse(path=%q, bound=%d, write=%t)\n", t.Segment, path, bound, write)
	i := len(t.segment)
	j := bound
	for i > 0 && j > 0 {
		i--
		j--
		if t.segment[i] != path[j] {
			i++
			j++
			break
		}
		// fmt.Printf("	i=%d j=%d => %q==%q\n", i, j, t.Segment[i], path[j])
	}
	if i == 0 {
		// fmt.Printf("	path[j-1]=%q Translate(path[j-1])=%d\n", path[j-1], Translate(path[j-1]))
		if j == 0 {
			return t
		}
		if t.wild {
			found = t
		}
		if node := t.children[compress(path[j-1])]; node != nil {

			return node.find(path, j, found)
		}
	}
	return found
}

func (t *Tree) Find(path string) *Tree {
	return t.find(path, len(path), nil)
}
