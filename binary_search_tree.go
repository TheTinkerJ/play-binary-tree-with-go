package tree

// BSTreeNode 二叉搜索树
// - 查找
// - 插入
// - 删除
type BSTreeNode struct {
	root *TreeNode
}

func (bst *BSTreeNode) PrettyPrint() {
	bst.root.prettyPrint()
}

func NewBST() *BSTreeNode {
	return &BSTreeNode{}
}

// Insert BST插入
func (bst *BSTreeNode) Insert(target int) *TreeNode {
	if bst.root == nil {
		bst.root = &TreeNode{
			Data: target,
		}
		return bst.root
	}
	current := bst.root
	for current != nil {
		if current.Data > target {
			if current.left == nil {
				current.left = &TreeNode{Data: target}
				return current.left
			} else {
				current = current.left
			}
		} else if current.Data < target {
			if current.right == nil {
				current.right = &TreeNode{Data: target}
				return current.right
			} else {
				current = current.right
			}
		} else {
			break
		}
	}
	return current
}

// Search BST查找
func (bst *BSTreeNode) Search(target int) *TreeNode {
	current := bst.root
	for current != nil {
		if current.Data == target {
			return current
		} else if current.Data > target {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil
}

// Delete BST删除
func (bst *BSTreeNode) Delete(target int) {
	// 第一步, 节点查找: 找到这个要被删除的节点,以及这个节点的父节点
	var parent *TreeNode
	current := bst.root
	for current != nil && current.Data != target {
		parent = current
		if current.Data > target {
			current = current.left
		} else if current.Data < target {
			current = current.right
		}
	}

	// 第二步, 节点删除
	// (1) 没找到
	if current == nil {
		return
	}
	// (2) 找到了
	// - 待删除节点同时存在左右孩子节点
	//  找到右最小节点,把节点内容互换一下
	//  然后删除这个找到的节点
	if current.left != nil && current.right != nil {
		rightParent, rightMinNode := current.rightMin()
		current.Data = rightMinNode.Data
		// 待删除节点替换
		parent = rightParent
		current = rightMinNode
	}
	// - 但删除节点如果只有左右半边,使用child来记录
	var child *TreeNode
	if current.left != nil {
		child = current.left
	} else if current.right != nil {
		child = current.right
	}
	// 三种情况的统一处理
	// 叶子节点 child == nil
	// 单孩子节点 child
	// 双孩子节点 current 已经完成替换, 这个时候要删除的节点可能是 单孩子节点或者叶子节点
	if parent == nil {
		bst.root = child
	} else if parent.left == current {
		parent.left = child
	} else if parent.right == current {
		parent.right = child
	}
	// 妙啊!
}
