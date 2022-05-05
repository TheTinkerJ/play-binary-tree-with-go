package tree

type TreeNode struct {
	Data  int
	left  *TreeNode
	right *TreeNode
}

func (root *TreeNode) leftMax() (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	parent, current := root, root.left
	for current != nil && current.right != nil {
		current = current.right
	}
	return parent, current
}

func (root *TreeNode) rightMin() (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	parent, current := root, root.right
	for current != nil && current.left != nil {
		parent = current
		current = current.left
	}
	return parent, current
}
