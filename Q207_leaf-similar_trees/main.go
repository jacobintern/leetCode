package Q207

import "reflect"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var arr1, arr2 []int
	traverse(root1, &arr1)
	traverse(root2, &arr2)

	return reflect.DeepEqual(arr1, arr2)
}

func traverse(root *TreeNode, arr *[]int) {
	if root.Left == nil && root.Right == nil {
		*arr = append(*arr, root.Val)
	}

	if root.Left != nil {
		traverse(root.Left, arr)
	}

	if root.Right != nil {
		traverse(root.Right, arr)
	}
}
