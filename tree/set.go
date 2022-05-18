package tree

func (t *Tree) Set(path string, value string, wild bool) *Tree {
	node := t.set(path, len(path))
	node.Value = value
	node.wild = wild
	return node
}

func (t *Tree) Append(node *Tree) *Tree {
	// fmt.Printf("[%q]	Append(node=[%q])\n", t.Segment, node.Segment)
	t.children[node.id()] = node
	node.parent = t
	return node
}

func (t *Tree) swap(node *Tree) *Tree {
	tmp := *t
	*t = *node
	*node = tmp
	return node
}

func (t *Tree) id() byte {
	return compress(t.segment[len(t.segment)-1])
}

func (t *Tree) fork(at int, insert string) *Tree {
	// fmt.Printf("[%q]	Branch(at=%d, insert=%q)\n", t.Segment, at, insert)
	node := Empty(insert)
	return t.split(at).Append(node)
}

func (t *Tree) split(at int) *Tree {
	// fmt.Printf("[%q]	Split(at=%d)\n", t.Segment, at)
	root := Empty(t.segment[at:])
	t.segment = t.segment[:at]
	t.Append(t.swap(root))
	return t
}

func (t *Tree) set(path string, bound int) *Tree {
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
	if j == 0 {
		if i == 0 {
			return t
		}
		return t.split(i)

	} else if i == 0 {
		// fmt.Printf("	path[j-1]=%q Translate(path[j-1])=%d\n", path[j-1], Translate(path[j-1]))
		if node := t.children[compress(path[j-1])]; node != nil {
			return node.set(path, j)
		}
		return t.Append(Empty(path[:j]))

	}
	return t.fork(i, path[:j])
}
