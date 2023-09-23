package Q100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res, m, maxCount := []int{}, make(map[int]int), 0

	loopTree(root, m)

	for k, v := range m {
		if v > maxCount {
			res = []int{k}
			maxCount = v
		} else if v == maxCount {
			res = append(res, k)
		}
	}

	return res
}

func loopTree(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}
	m[root.Val]++
	loopTree(root.Left, m)
	loopTree(root.Right, m)
}
