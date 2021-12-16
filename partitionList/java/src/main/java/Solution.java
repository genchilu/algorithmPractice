class Solution {

public class ListNode {
    int val;
    ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
}

    public ListNode partition(ListNode head, int x) {

        if(head == null) {
            return head;
        }


        ListNode cur = head;
        ListNode lh = null;
        ListNode lt = null;
        ListNode rh = head;
        ListNode pre = null;

        while(cur != null && cur.val<x) {
            if(lh==null) {
                lh =cur;
                lt = cur;
            } else {
                lt.next = cur;
                lt = lt.next;
            }
            cur = cur.next;
            lt.next = null;
        }

        if(cur==null) {
            return lh;
        }


        pre = cur;
        rh = pre;
        cur=cur.next;


        while(cur!=null) {
            if(cur.val<x){
                if(lh==null) {
                    lh =cur;
                    lt = cur;
                } else {
                    lt.next = cur;
                    lt = lt.next;
                }
                cur=cur.next;
                pre.next=cur;
                lt.next=null;
            } else {
                pre=cur;
                cur=cur.next;
            }
        }

        if(lh == null) {
            return head;
        }

        lt.next = rh;

//         cur = lh;
//         while(cur!=null){
//             System.out.println(cur.val);
//             cur=cur.next;
//         }


//         System.out.println("=======");
//         cur = rh;
//         while(cur!=null){
//             System.out.println(cur.val);
//             cur=cur.next;
//         }
        return lh;
    }
}