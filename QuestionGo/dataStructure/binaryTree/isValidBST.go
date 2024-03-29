package binaryTree

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val >= root.Val {

		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	if !isValidBST(root.Left) || !isValidBST(root.Right) {
		return false
	}
	return true
}
