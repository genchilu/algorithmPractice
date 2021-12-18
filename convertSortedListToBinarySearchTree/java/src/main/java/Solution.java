/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */

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
    public class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

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

    public TreeNode sortedListToBST(ListNode head) {
        List<ListNode> nodes = new ArrayList<>();

        var cur = head;
        while(cur!=null) {
            nodes.add(cur);
            cur = cur.next;
        }

        return genBST(nodes);
    }

    public TreeNode genBST(List<ListNode> nodes) {
        if(nodes.size()==0) {
            return null;
        }

        if(nodes.size()==1) {
            return new TreeNode(nodes.get(0).val);
        }

        var midx = (0+nodes.size())/2;
        var root = new TreeNode(nodes.get(midx).val);
        root.left = genBST(nodes.subList(0, midx));
        root.right = genBST(nodes.subList(midx+1, nodes.size()));

        return root;
    }
}