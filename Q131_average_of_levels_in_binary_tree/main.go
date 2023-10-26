package Q131

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type content struct {
	value int
	count int
}

var m map[int]*content

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	m = make(map[int]*content)
	avg(root, 0)

	for i := 1; i <= len(m); i++ {
		res = append(res, float64(m[i].value)/float64(m[i].count))
	}
	return res
}

func avg(root *TreeNode, level int) {
	if root == nil {
		return
	}

	level = level + 1
	if v, exist := m[level]; exist {
		v.value += root.Val
		v.count += 1
	} else {
		m[level] = &content{value: root.Val, count: 1}
	}

	avg(root.Left, level)
	avg(root.Right, level)
}
