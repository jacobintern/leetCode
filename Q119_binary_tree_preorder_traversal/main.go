package Q119

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	traversalTree(root, &res)
	return res
}

func traversalTree(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	*arr = append(*arr, root.Val)

	traversalTree(root.Left, arr)
	traversalTree(root.Right, arr)
}
