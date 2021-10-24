import java.util.HashMap;
import java.util.Map;

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

    Map<TreeNode, Integer> map = new HashMap<>();
    public int rob(TreeNode root) {
        if (!map.containsKey(root)) {

            int a = 0;
            int b = root.val;
            if (root.left != null) {
                a += rob(root.left);
                if (root.left.left != null) {
                    b += rob(root.left.left);
                }
                if (root.left.right != null) {
                    b += rob(root.left.right);
                }
            }
            if (root.right != null) {
                a += rob(root.right);
                if (root.right.left != null) {
                    b += rob(root.right.left);
                }
                if (root.right.right != null) {
                    b += rob(root.right.right);
                }
            }

            if (a > b) {
                map.put(root, a);
            } else {
                map.put(root, b);
            }
        }

        return map.get(root);
    }
}