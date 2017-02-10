package merkle

import (
	"fmt"
	"testing"
)

func prepData(c int) [][]byte {
	d := make([][]byte, c)
	for i := 0; i < c; i++ {
		d[i] = []byte(fmt.Sprintf("data%d", i))
	}
	return d
}

func TestMerkleTreeEven(t *testing.T) {
	pd := prepData(10)

	tree := GenerateTree(pd)

	levels := tree.levels
	ll := len(levels)
	for i := 1; i < ll; i++ {
		for _, v := range levels[i] {
			if v.Left == nil || v.Right == nil {
				t.Error("left/right should not be nil")
			}
		}
	}

	if len(tree.Leafs()) != 10 {
		t.Fatal("wrong leaves")
	}

	if tree.Root() == nil {
		t.Fatal("root is nil")
	}
}

func TestMerkleTreeOdd(t *testing.T) {
	pd := prepData(11)
	tree := GenerateTree(pd)

	levels := tree.levels
	ll := len(levels)
	for i := 1; i < ll; i++ {
		for _, v := range levels[i] {
			if v.Left == nil || v.Right == nil {
				t.Error("left/right should not be nil")
			}
		}
	}

	if len(tree.Leafs()) != 12 {
		t.Fatal("wrong leaves")
	}

}
