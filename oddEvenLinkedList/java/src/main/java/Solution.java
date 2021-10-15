



class Solution {
    public class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    public ListNode oddEvenList(ListNode head) {
        if(head == null || head.next==null) {
            return head;
        }

        ListNode oh = head;
        ListNode eh = head.next;
        ListNode ot = oh;
        ListNode et = eh;

        while (true) {
            if(et.next!=null) {
                ot.next = et.next;
                ot = ot.next;

                if(ot.next != null) {
                    et.next = ot.next;
                    et = et.next;
                } else {
                    et.next = null;
                    break;
                }
            } else {
                ot.next = null;
                break;
            }
        }

        ot.next=eh;

        return head;
    }
}