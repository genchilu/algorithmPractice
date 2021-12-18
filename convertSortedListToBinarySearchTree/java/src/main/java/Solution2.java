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
class Solution2 {
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

    private ListNode head;

    public TreeNode sortedListToBST(ListNode head) {
        this.head = head;
        var count =0;
        var cur =head;
        while(cur!=null){
            cur=cur.next;
            count++;
        }


        return genBST(1, count);
    }

    public TreeNode genBST(int start, int end) {
        if(start>end) {
            return null;
        }

        var mid = (start+end)/2;
        //System.out.println("head " + head.val + " start " + start + ", end " + end + ", mid " + mid +" fine left");
        var l = genBST(start, mid-1);
        var root = new TreeNode(head.val);
        root.left = l;

        //System.out.println("head " + head.val + " start " + start + ", end " + end + ", mid " + mid +", root " + root.val);

        head=head.next;
        //System.out.println("head " + head.val + " start " + start + ", end " + end + ", mid " + mid +" fine right");
        root.right=genBST(mid+1,end);

        return root;
    }
}