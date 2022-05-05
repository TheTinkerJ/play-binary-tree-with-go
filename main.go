package tree

func main() {
	t7 := &TreeNode{Data: 7}
	//t6 := &TreeNode{data: 6}
	t5 := &TreeNode{Data: 555}
	t4 := &TreeNode{Data: 4}
	t3 := &TreeNode{Data: 333, right: t7}
	t2 := &TreeNode{Data: 22, left: t4, right: t5}
	t1 := &TreeNode{Data: 11111, left: t2, right: t3}
	t1.prettyPrint()
}
