package tree

func (t *Tree[T]) Set(key string, value T, wild bool) {
	node := t.set(key, len(key)) // TODO: normalize key input for all
	node.Value = value
	node.Wild = wild
	node.leaf = true
}

func (t *Tree[T]) Append(node *Tree[T]) *Tree[T] {
	// fmt.Printf("[%q]	Append(node=[%q])\n", t.Segment, node.Segment)
	t.children[node.id()] = node
	t.count++
	node.parent = t
	return node
}

func (t *Tree[T]) swap(node *Tree[T]) *Tree[T] {
	tmp := *t
	*t = *node
	*node = tmp
	return node
}

func (t *Tree[T]) id() byte {
	return hash(t.segment[len(t.segment)-1])
}

func (t *Tree[T]) fork(at int, insert string) *Tree[T] {
	// fmt.Printf("[%q]	Branch(at=%d, insert=%q)\n", t.Segment, at, insert)
	node := Empty[T](insert)
	return t.split(at).Append(node)
}

func (t *Tree[T]) split(at int) *Tree[T] {
	// fmt.Printf("[%q]	Split(at=%d)\n", t.Segment, at)
	root := Empty[T](t.segment[at:])
	t.segment = t.segment[:at]
	t.Append(t.swap(root))
	return t
}

func (t *Tree[T]) set(key string, bound int) *Tree[T] {
	// fmt.Printf("[%q]	Traverse(path=%q, bound=%d, write=%t)\n", t.Segment, path, bound, write)
	i := len(t.segment)
	j := bound
	for i > 0 && j > 0 {
		i--
		j--
		if t.segment[i] != key[j] {
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
		if node := t.children[hash(key[j-1])]; node != nil {
			return node.set(key, j)
		}
		if t.count == 0 && !t.leaf {
			t.segment += key[:j]
			return t
		}
		return t.Append(Empty[T](key[:j]))
	}
	return t.fork(i, key[:j])
}
