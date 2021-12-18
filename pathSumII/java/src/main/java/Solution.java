import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

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
        TreeNode() {}
        TreeNode(int val) { this.val = val; }
        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    public List<List<Integer>> pathSum(TreeNode root, int targetSum) {
        var res = new ArrayList<List<Integer>>();
        dfs(root, targetSum, new ArrayList<Integer>(), res);

        return res;
    }

    public void dfs(TreeNode node, int target, List<Integer> vals, List<List<Integer>> res) {
        if(node==null) {
            return;
        }
        if(node.left==null && node.right==null && node.val == target) {
            vals.add(node.val);
            res.add(vals);
        }

        if(node.left!=null) {
            var cloneVal = vals.stream().collect(Collectors.toList());
            cloneVal.add(node.val);
            dfs(node.left, target-node.val, cloneVal, res);
        }

        if(node.right!=null) {
            var cloneVal = vals.stream().collect(Collectors.toList());
            cloneVal.add(node.val);
            dfs(node.right, target-node.val, cloneVal, res);
        }
    }
}