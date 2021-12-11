package Q41

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var r []int

func inorderTraversal(root *TreeNode) []int {
	r = []int{}
	sortByLeftTree(root)
	return r
}

func sortByLeftTree(root *TreeNode) {
	if root == nil {
		return
	}

	sortByLeftTree(root.Left)
	r = append(r, root.Val)
	sortByLeftTree(root.Right)
}
