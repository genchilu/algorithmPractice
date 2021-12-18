import java.util.Arrays;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {

    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode() {
        }

        TreeNode(int val) {
            this.val = val;
        }

        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    public TreeNode buildTree(int[] inorder, int[] postorder) {
        if (inorder.length==0) {
            return null;
        }

        var rootVal = postorder[postorder.length-1];

        var idx = 0;
        while(inorder[idx]!=rootVal) {
            idx++;
        }

        var linorder = Arrays.copyOfRange(inorder, 0, idx);
        var rinorder = Arrays.copyOfRange(inorder, idx+1, inorder.length);

        var lpostorder = Arrays.copyOfRange(postorder, 0, idx);
        var rpostorder = Arrays.copyOfRange(postorder, idx, inorder.length-1);

        var root = new TreeNode(rootVal);
        root.left = buildTree(linorder, lpostorder);
        root.right = buildTree(rinorder, rpostorder);

        return root;
    }
}