package practice_problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// Invert in parallel
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
