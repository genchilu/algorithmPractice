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
    public TreeNode deleteNode(TreeNode root, int key) {
        SearchRes searchRes = search(null, root, key);
        if(searchRes == null) {
            return root;
        }

        if(searchRes.parentTree==null) {
            if(searchRes.targetTree.left == null && searchRes.targetTree.right == null) {
                return null;
            }

            if(searchRes.targetTree.left == null) {
                return searchRes.targetTree.right;
            }

            if(searchRes.targetTree.right == null) {
                return searchRes.targetTree.left;
            }

            TreeNode replaceTree = extractMin(searchRes.targetTree, searchRes.targetTree.right);

            replaceTree.left = searchRes.targetTree.left;
            replaceTree.right = searchRes.targetTree.right;

            return replaceTree;
        }

        if(searchRes.targetTree.left == null && searchRes.targetTree.right == null) {
            if(searchRes.parentTree.left != null && searchRes.parentTree.left.val == searchRes.targetTree.val) {
                searchRes.parentTree.left = null;
            } else {
                searchRes.parentTree.right = null;
            }
            return root;
        }

        if(searchRes.targetTree.left == null) {
            if(searchRes.parentTree.left != null && searchRes.parentTree.left.val == searchRes.targetTree.val) {
                searchRes.parentTree.left = searchRes.targetTree.right;
            } else {
                searchRes.parentTree.right = searchRes.targetTree.right;
            }
            return root;
        }

        if(searchRes.targetTree.right == null) {
            if(searchRes.parentTree.left != null &&  searchRes.parentTree.left.val == searchRes.targetTree.val) {
                searchRes.parentTree.left = searchRes.targetTree.left;
            } else {
                searchRes.parentTree.right = searchRes.targetTree.left;
            }
            return root;
        }

        TreeNode replaceTree = extractMin(searchRes.targetTree, searchRes.targetTree.right);
        replaceTree.left = searchRes.targetTree.left;
        replaceTree.right = searchRes.targetTree.right;

        if(searchRes.parentTree.left != null &&  searchRes.parentTree.left.val == searchRes.targetTree.val) {
            searchRes.parentTree.left = replaceTree;
        } else {
            searchRes.parentTree.right = replaceTree;
        }

        return root;
    }

    class SearchRes {
        TreeNode parentTree;
        TreeNode targetTree;

        public SearchRes(TreeNode parentTree, TreeNode targetTree) {
            this.parentTree = parentTree;
            this.targetTree = targetTree;
        }
    }

    public SearchRes search(TreeNode parent, TreeNode tree, int target) {
        if(tree == null) {
            return null;
        }

        if(tree.val==target) {
            return new SearchRes(parent, tree);
        }


        SearchRes leftRes = search(tree, tree.left, target);
        if(leftRes!=null) {
            return leftRes;
        }


        SearchRes rightRes = search(tree, tree.right, target);
        if(rightRes!=null) {
            return rightRes;
        }

        return null;
    }

    public TreeNode extractMin(TreeNode parent, TreeNode tree) {
        if(tree == null) {
            return null;
        }

        TreeNode min = extractMin(tree, tree.left);

        if(min == null) {
            if(parent.left.val == tree.val) {
                parent.left = tree.right;
            } else {
                parent.right = tree.right;
            }
            tree.right = null;
            return tree;
        }

        return min;
    }
}