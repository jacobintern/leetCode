package Q24

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := TreeDepth(root.Left, true)
	r := TreeDepth(root.Right, false)

	if l == 1 || (r > 1 && r < l) {
		return r
	}
	return l

}

func TreeDepth(root *TreeNode, p bool) int {
	if root == nil {
		return 1
	}

	if p {
		fmt.Println("val: ", root.Val)
		res := TreeDepth(root.Left, true) + 1
		return res
	}
	fmt.Println("val: ", root.Val)
	res := TreeDepth(root.Right, false) + 1
	return res
}
