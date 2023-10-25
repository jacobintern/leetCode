package Q131

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[int]map[int]int

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	m = make(map[int]map[int]int)

	avg(root, 0)
	return res
}

func avg(root *TreeNode, level int) {
	if root == nil {
		return
	}

	level = level + 1
	m[level] = root.Val

	avg(root.Left, level)
	avg(root.Right, level)
}
