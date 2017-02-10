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
	if data == nil {
		n.hash = make([]byte, 32)
	} else {
		sh := fastsha256.Sum256(data)
		n.hash = sh[:]
	}
	return n
}
