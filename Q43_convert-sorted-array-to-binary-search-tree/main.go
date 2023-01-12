package Q43

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return ToBST(nums, 0, len(nums)-1)
}

func ToBST(nums []int, i, j int) *TreeNode {
	if i > j {
		return nil
	}

	sum := i + (j-i)/2

	l := ToBST(nums, i, sum-1)
	r := ToBST(nums, sum+1, j)
	return &TreeNode{Val: nums[sum], Left: l, Right: r}
}
