
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
    public int longestConsecutive(TreeNode root) {

        var res = countConsecutive(root);
        return res.maxCount;
    }

    class CountRes {
        int maxCount;
        int consecutiveCount;

        public CountRes(int maxCount, int consecutiveCount) {
            this.maxCount = maxCount;
            this.consecutiveCount = consecutiveCount;
        }

    }

    public CountRes countConsecutive(TreeNode root) {
        if(root == null) {
            return new CountRes(0,0);
        }

        var res = new CountRes(1, 1);

        if(root.left!=null) {
            var lres = countConsecutive(root.left);
            if(root.val == root.left.val-1) {
                res.consecutiveCount = lres.consecutiveCount+1;
            }
            res.maxCount = lres.maxCount;
        }

        if(root.right!=null) {
            var rres = countConsecutive(root.right);
            if(root.val == root.right.val-1 && (rres.consecutiveCount+1)>res.consecutiveCount) {
                res.consecutiveCount = rres.consecutiveCount+1;
            }
            if(rres.maxCount> res.maxCount) {
                res.maxCount = rres.maxCount;
            }
        }

        if (res.consecutiveCount>= res.maxCount) {
            res.maxCount =res.consecutiveCount;
        }


        return res;
    }
}