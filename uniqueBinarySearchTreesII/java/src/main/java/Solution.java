import java.util.ArrayList;
import java.util.List;

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

    public List<TreeNode> generateTrees(int n) {
        List<TreeNode> res = new ArrayList<>();
        return this.generateTrees(1,n);
    }

    public List<TreeNode> generateTrees(int s, int e) {
        List<TreeNode> res = new ArrayList<>();

        if(s>e) {
            return res;
        }

        if ((e-s+1)==1) {
            res.add(new TreeNode(s));
            return res;
        }

        for(int i=s;i<=e;i++){
            var lres = generateTrees(s, i-1);
            var rres = generateTrees(i+1, e);

            if(lres.size()==0 && rres.size()==0) {
                var root = new TreeNode(i);
                res.add(root);
            } else if(lres.size()==0) {
                for(var r:rres) {
                    var root = new TreeNode(i);
                    root.right = r;
                    res.add(root);
                }
            } else if(rres.size() == 0) {
                for(var l:lres) {
                    var root = new TreeNode(i);
                    root.left = l;
                    res.add(root);
                }
            } else {
                for(var r:rres) {
                    for(var l:lres) {
                        var root = new TreeNode(i);
                        root.left = l;
                        root.right = r;
                        res.add(root);
                    }
                }
            }
        }
        return res;
    }
}