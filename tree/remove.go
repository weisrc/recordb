package tree

func (t *Tree[T]) Remove(key string) *Tree[T] {
	if node := t.At(key, false); node != nil {
		if parent := node.parent; parent != nil {
			parent.children[node.id()] = nil
			parent.count--
			if !parent.leaf && parent.count == 1 {
				for _, node2 := range parent.children {
					if node2 != nil {
						parent.segment += node2.segment
						parent.swap(node2)
						break
					}
				}
			}
			return node
		} else {
			empty := Empty[T]("")
			empty.swap(t)
			return empty
		}
	}
	return nil
}
