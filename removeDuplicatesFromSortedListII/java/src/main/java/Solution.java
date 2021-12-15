import java.util.HashSet;
import java.util.Set;

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
class Solution {
    public class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    public ListNode deleteDuplicates(ListNode head) {

        if(head == null) {
            return head;
        }

        Set<Integer> set = new HashSet<>();

        var pre = head;
        var cur = pre.next;

        while(cur != null) {
            if(cur.val == pre.val) {
                set.add(cur.val);
            }
            pre = pre.next;
            cur = cur.next;
        }


        while(head !=null && set.contains(head.val)) {
            head = head.next;
        }


        if(head==null) {
            return head;
        }

        pre = head;
        cur = pre.next;
        while(cur!=null){
            if(set.contains(cur.val)) {
                pre.next = cur.next;
                cur = cur.next;
            } else {
                cur = cur.next;
                pre = pre.next;
            }
        }

        return head;
    }
}