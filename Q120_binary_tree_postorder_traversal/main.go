package Q120

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	traversalTree(root, &res)
	return res
}

func traversalTree(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	traversalTree(root.Left, arr)
	traversalTree(root.Right, arr)

	*arr = append(*arr, root.Val)
}
