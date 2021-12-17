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

    public ListNode reverseBetween(ListNode head, int left, int right) {
        ListNode preL = null;
        ListNode afterR = null;
        ListNode l = null;
        ListNode r = null;
        ListNode cur = head;

        int count = 1;
        while(count<left) {
            preL = cur;
            cur=cur.next;
            count++;
        }
        l = cur;

        while(count < right) {
            cur = cur.next;
            count++;
        }
        r=cur;

        afterR = r.next;
        r.next = null;


        var  pre = l;
        cur = l.next;
        pre.next = null;

        while(cur!=null) {
            var tmp = cur.next;
            cur.next = pre;
            pre = cur;
            cur = tmp;
        }

        link(preL, afterR, r,l);

        if(preL == null) {
            return r;
        }

        return head;
    }

    public void link(ListNode preL, ListNode afterR, ListNode head, ListNode tail) {
        if (preL!=null) {
            preL.next = head;
        }

        tail.next = afterR;
    }
}