package merkle

import (
	"encoding/hex"

	"github.com/btcsuite/fastsha256"
)

// Node represents a single node in the merkle tree. It can be a leaf
// or non-leaf node.
type Node struct {
	hash []byte
	Data []byte

	Left  *Node
	Right *Node
}

// Hash returns the cached hash
func (n *Node) Hash() []byte {
	return n.hash
}

func (n *Node) String() string {
	return hex.EncodeToString(n.hash)
}

// IsLeaf returns whether this node is a leaf
func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

// NewNode instantiates a new node with the given data and computes and caches the
// hash of the data
func NewNode(data []byte) *Node {
	n := &Node{Data: data}
	sh := fastsha256.Sum256(data)
	n.hash = sh[:]
	return n
}

func generateLeafNodes(data [][]byte) []*Node {
	l := len(data)
	lnodes := make([]*Node, l)

	for i, d := range data {
		lnodes[i] = NewNode(d)
	}

	if (l % 2) != 0 {
		return append(lnodes, lnodes[l-1])
	}
	return lnodes
}

// if ns is odd the last node is repeated.
func generateParentLevel(ns []*Node) []*Node {
	nodes := ns

	if (len(ns) % 2) != 0 {
		nodes = append(nodes, nodes[len(nodes)-1])
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

	return parent
}
