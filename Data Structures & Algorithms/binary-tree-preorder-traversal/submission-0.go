/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    ans := []int{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node==nil { return }
		ans = append(ans, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return ans
}
