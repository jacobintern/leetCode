package Q108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	getTreeSum(root, &sum)
	return sum
}

func getTreeSum(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}

	l := getTreeSum(root.Left, sum)
	r := getTreeSum(root.Right, sum)

	*sum += abs(l - r)

	return l + r + root.Val
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
