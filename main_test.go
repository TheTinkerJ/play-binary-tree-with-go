package tree

import (
	"testing"
)

func TestSTreeNode_Insert(t *testing.T) {
	bst := NewBST()
	bst.Insert(22)
	bst.Insert(4)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(33)
	bst.Insert(5)
	bst.Insert(6)
	bst.Insert(7)
	bst.PrettyPrint()
	bst.Delete(22)
	bst.PrettyPrint()
}
