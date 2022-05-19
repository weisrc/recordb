package tree

func (t *Tree[T]) get(key string, bound int, wild bool, found *Tree[T]) *Tree[T] {
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
	if i == 0 {
		// fmt.Printf("	path[j-1]=%q Translate(path[j-1])=%d\n", path[j-1], Translate(path[j-1]))
		if j == 0 {
			return t
		}
		if wild && t.Wild {
			found = t
		}
		if node := t.children[hash(key[j-1])]; node != nil {

			return node.get(key, j, wild, found)
		}
	}
	return found
}

func (t *Tree[T]) At(key string, wild bool) *Tree[T] {
	return t.get(key, len(key), wild, nil)
}

func (t *Tree[T]) Get(key string, wild bool) (value T, ok bool) {
	if node := t.At(key, wild); node != nil {
		return node.Value, true
	}
	return
}
