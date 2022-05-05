package tree

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// RenderElement 特殊渲染字符
type RenderElement string

const (
	RdrEmpty RenderElement = " "
	RdrPlain RenderElement = "━"
	RdrUpper RenderElement = "┻"
	RdrLeft  RenderElement = "┏"
	RdrRight RenderElement = "┓"
)

// TreeNodeType 节点类型
type TreeNodeType int

const (
	NodeRoot  TreeNodeType = 0 // 根节点
	NodeLeft  TreeNodeType = 1 // 左孩子
	NodeRight TreeNodeType = 2 // 右孩子
)

type PrintTreeNode struct {
	Content  []string       // 节点内容 stage1:拷贝
	NodeType TreeNodeType   // 节点类型 stage1:拷贝
	left     *PrintTreeNode // 左孩子 stage1:拷贝
	right    *PrintTreeNode // 右孩子 stage1:拷贝
	LeftPos  int            // 内容占位左边界 stage2:中序扫描
	RightPos int            // 内容占位右边界 stage2:中序扫描
}

// NewPrintTreeNode :TreeNode->PrintTreeNode
func NewPrintTreeNode(root *TreeNode) *PrintTreeNode {
	return &PrintTreeNode{
		Content: strings.Split(strconv.Itoa(root.Data), ""),
	}
}

// BuildPrintTreeNode  stage1 先序递归的树拷贝
func BuildPrintTreeNode(root *TreeNode) *PrintTreeNode {
	if root == nil {
		return nil
	}

	ptn := NewPrintTreeNode(root)
	ptn.NodeType = NodeRoot

	if root.left != nil {
		ptn.left = BuildPrintTreeNode(root.left)
		if ptn.left != nil {
			ptn.left.NodeType = NodeLeft
		}
	}

	if root.right != nil {
		ptn.right = BuildPrintTreeNode(root.right)
		if ptn.right != nil {
			ptn.right.NodeType = NodeRight
		}
	}
	return ptn
}

// inOrderGenDuration stage2 非递归中序计算每一个节点的行占位信息
func (root *PrintTreeNode) inOrderGenDuration() int {
	// >> 非模板代码
	forward := 0
	// ***********
	current := root
	stack := list.New()
	for current != nil || stack.Len() != 0 {
		for current != nil {
			stack.PushFront(current)
			current = current.left
		}
		if stack.Len() != 0 {
			currentTmp, _ := stack.Remove(stack.Front()).(*PrintTreeNode)
			// >> 非模板代码,中序节点操作逻辑区*****************************
			contentLength := len(currentTmp.Content)
			currentTmp.LeftPos = forward
			currentTmp.RightPos = forward + contentLength - 1
			forward = forward + contentLength
			// *******************************************************
			current = currentTmp.right
		}
	}
	return forward
}

// levelOrderHandle 层序遍历的思路,进行一行一行的绘制
func (root *PrintTreeNode) levelOrderHandle(width int) [][]string {
	view := make([][]string, 0)
	queue := list.New()
	queue.PushBack(root)
	depth := 0
	for queue.Len() != 0 {
		levelSize := queue.Len()
		numberRow := make([]string, width)
		iconRow, iconFilled := make([]string, width), false
		for i := 0; i < width; i++ {
			numberRow[i] = string(RdrEmpty)
			iconRow[i] = string(RdrEmpty)
		}

		for levelSize > 0 {
			currentTmp, _ := queue.Remove(queue.Front()).(*PrintTreeNode)
			// >> 层序遍历,非模板代码区
			// 只要有下层节点就需要渲染 iconRow
			if currentTmp.left != nil || currentTmp.right != nil {
				iconFilled = true
			}
			TreeNodeHandle(currentTmp, numberRow, iconRow)
			// ****************
			if currentTmp.left != nil {
				queue.PushBack(currentTmp.left)
			}
			if currentTmp.right != nil {
				queue.PushBack(currentTmp.right)
			}
			levelSize--
		}
		// >> 层序遍历,非模板代码区
		view = append(view, numberRow)
		if iconFilled {
			view = append(view, iconRow)
		}
		// ****************
		depth++
	}
	return view
}

func TreeNodeHandle(ptn *PrintTreeNode, numberRow []string, iconRow []string) {
	content := ptn.Content
	for i := ptn.LeftPos; i <= ptn.RightPos; i++ {
		numberRow[i] = content[i-ptn.LeftPos]
	}

	if ptn.left == nil && ptn.right == nil {
		return
	}

	leftBdr, rightBdr := ptn.LeftPos, ptn.LeftPos
	if ptn.NodeType == NodeRight {
		leftBdr = ptn.RightPos
		rightBdr = ptn.RightPos
	}
	if ptn.left != nil {
		leftBdr = ptn.left.LeftPos
	}
	if ptn.right != nil {
		rightBdr = ptn.right.RightPos
	}
	// LEFT/RIGHT/PLAIN
	iconRow[leftBdr] = string(RdrLeft)
	iconRow[rightBdr] = string(RdrRight)
	for i := leftBdr + 1; i < rightBdr; i++ {
		iconRow[i] = string(RdrPlain)
	}
	//UPPER
	if ptn.NodeType == NodeLeft || ptn.NodeType == NodeRoot {
		iconRow[ptn.LeftPos] = string(RdrUpper)
	} else if ptn.NodeType == NodeRight {
		iconRow[ptn.RightPos] = string(RdrUpper)
	}
}

func (t *TreeNode) prettyPrint() {
	if t == nil {
		fmt.Println("Tree should not be Empty")
		return
	}
	ptn := BuildPrintTreeNode(t)
	width := ptn.inOrderGenDuration()
	view := ptn.levelOrderHandle(width)
	height := len(view)
	for j := 0; j < width; j++ {
		fmt.Print("-")
	}
	fmt.Println()
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(view[i][j])
		}
		fmt.Println()
	}
	for j := 0; j < width; j++ {
		fmt.Print("-")
	}
	fmt.Println()
	fmt.Println()
}
