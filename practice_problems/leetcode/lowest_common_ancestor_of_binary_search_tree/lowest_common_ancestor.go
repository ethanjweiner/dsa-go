package practice_problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// BST-Specific (assumes that LCA exists)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

// For any binary tree
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	leftLCA := lowestCommonAncestor(root.Left, p, q)

// 	if leftLCA != nil {
// 		return leftLCA
// 	}

// 	rightLCA := lowestCommonAncestor(root.Right, p, q)

// 	if rightLCA != nil {
// 		return rightLCA
// 	}

// 	if search(root, p) && search(root, q) {
// 		return root
// 	}

// 	return nil
// }

// func search(root *TreeNode, n *TreeNode) bool {
// 	switch {
// 	case root == nil:
// 		return false
// 	case n.Val < root.Val:
// 		return search(root.Left, n)
// 	case n.Val > root.Val:
// 		return search(root.Right, n)
// 	}

// 	return true
// }
