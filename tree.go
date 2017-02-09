package merkle

// Tree represents a merkle tree
type Tree struct {
	// levels contains all the levels with the leaf nodes at the top
	levels [][]*Node
}

// GenerateTree generates a merkle tree from the given data.
func GenerateTree(data [][]byte) *Tree {
	levels := [][]*Node{generateLeafNodes(data)}
	for {
		ll := levels[len(levels)-1]
		pnodes := generateParentLevel(ll)
		levels = append(levels, pnodes)

		if len(pnodes) == 1 {
			break
		}
	}

	return &Tree{levels: levels}
}

// Root returns the root node of the tree
func (t *Tree) Root() *Node {
	rl := t.levels[len(t.levels)-1]
	return rl[0]
}
