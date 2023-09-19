package Q112

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	traversalTree(root, &res)
	return res
}

func traversalTree(root *Node, arr *[]int) {
	if root == nil {
		return
	}

	for _, v := range root.Children {
		traversalTree(v, arr)
	}
	*arr = append(*arr, root.Val)
}
