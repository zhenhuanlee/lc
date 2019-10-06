package main

import "fmt"

// TreeNode tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int = []int{}

func inorderTraversal(root *TreeNode) []int {
	if root.Left != nil {
		inorderTraversal(root.Left)
	}

	fmt.Println(root.Val)
	res = append(res, root.Val)

	if root.Right != nil {
		inorderTraversal(root.Right)
	}

	return res
}

func main() {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	thr := &TreeNode{Val: 3}
	fou := &TreeNode{Val: 4}
	fiv := &TreeNode{Val: 5}
	six := &TreeNode{Val: 6}
	sev := &TreeNode{Val: 7}

	fiv.Left = thr
	fiv.Right = six

	thr.Left = two
	thr.Right = fou

	two.Left = one

	six.Right = sev

	// v :=
	v := inorderTraversal(fiv)
	fmt.Println(v)
}
