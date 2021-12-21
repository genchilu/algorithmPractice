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

    public class CountResult {
        public int count;
        public boolean isUnival;
        public CountResult(int count, boolean isUnival) {
            this.count = count;
            this.isUnival = isUnival;
        }
    }

    public int countUnivalSubtrees(TreeNode root) {
        var c = countUnival(root);

        return c.count;
    }

    public CountResult countUnival(TreeNode root) {
        if(root == null) {
            return new CountResult(0, false);
        }
        if(root.right==null && root.left==null) {
            return new CountResult(1, true);
        }

        CountResult lc = countUnival(root.left);
        CountResult rc = countUnival(root.right);

        if(root.right==null) {
            if(root.val == root.left.val && lc.isUnival) {
                return new CountResult(lc.count+1, true);
            } else {
                lc.isUnival = false;
                return lc;
            }
        }

        if(root.left==null) {
            if(root.val==root.right.val && rc.isUnival) {
                return new CountResult(rc.count+1, true);
            } else {
                rc.isUnival=false;
                return rc;
            }
        }

        if(root.val == root.left.val && root.val == root.right.val && lc.isUnival && rc.isUnival) {
            return new CountResult(lc.count+rc.count+1, true);
        } else {
            return new CountResult(lc.count+rc.count, false);
        }
    }
}