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

}

func TestMerkleDiff(t *testing.T) {
	pd := prepData(11)
	src := GenerateTree(pd)

	pd = prepData(5)
	dst := GenerateTree(pd)

	//slvl := len(src.levels) - len(dst.levels)

	//dl:=dst.levels[slvl]

	dumpTree(src)
	fmt.Println("")
	dumpTree(dst)

}

func dumpTree(t1 *Tree) {
	for i, level := range t1.levels {
		fmt.Println("Level", i)
		for _, l := range level {
			fmt.Println(" ", l)
		}
	}
}
