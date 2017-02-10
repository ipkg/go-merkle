package merkle

// Tree represents a merkle tree
type Tree struct {
	// levels contains all the levels with the leaf nodes at the top i.e. 0 index
	// and the root node at the end of the slice
	levels [][]*Node
}

// GenerateTree generates a merkle tree from the given data.
func GenerateTree(data [][]byte) *Tree {
	xtras := 0

	leafs, added := generateLeafNodes(data)
	if added {
		xtras++
	}

	tree := &Tree{levels: [][]*Node{leafs}}
	if len(data) == 1 {
		return tree
	}

	for {
		ll := tree.levels[len(tree.levels)-1]
		pnodes, added := generateParentLevel(ll)
		if added {
			xtras++
		}
		tree.levels = append(tree.levels, pnodes)

		if len(pnodes) == 1 {
			break
		}
	}
	return tree
}

// Root returns the root node of the tree
func (t *Tree) Root() *Node {
	rl := t.levels[len(t.levels)-1]
	return rl[0]
}

// Height returns the height of the tree
func (t *Tree) Height() int {
	return len(t.levels)
}

// Leafs returns theh leaf nodes
func (t *Tree) Leafs() []*Node {
	return t.levels[0]
}

// If data is odd, then the last item is repeated
func generateLeafNodes(data [][]byte) ([]*Node, bool) {
	l := len(data)
	lnodes := make([]*Node, l)

	for i, d := range data {
		lnodes[i] = NewNode(d)
	}

	if (l % 2) != 0 {
		//return append(lnodes, NewNode(nil))
		return append(lnodes, lnodes[l-1]), true
	}
	return lnodes, false
}

// If ns is odd the last node is repeated.
func generateParentLevel(ns []*Node) ([]*Node, bool) {
	nodes := ns

	added := false
	if (len(ns) % 2) != 0 {
		nodes = append(nodes, nodes[len(nodes)-1])
		added = true
		//nodes = append(nodes, NewNode(nil))
	}

	ln := len(nodes)
	parent := make([]*Node, ln/2)

	for i := 0; i < ln; i += 2 {
		data := append(nodes[i].Hash(), nodes[i+1].Hash()...)
		nn := NewNode(data)
		nn.Left = nodes[i]
		nn.Right = nodes[i+1]
		parent[i/2] = nn
	}

	return parent, added
}
