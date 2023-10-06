package Q125

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	swipe(root)
	return root
}

func swipe(root *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return
	}
	l := root.Left
	r := root.Right
	root.Right, root.Left = l, r

	if root.Left != nil {
		swipe(root.Left)
	}
	if root.Right != nil {
		swipe(root.Right)
	}
}
