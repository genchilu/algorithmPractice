import javax.swing.*;
import java.util.HashMap;

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

    public void reorderList(ListNode head) {
        if(head.next==null) {
            return;
        }

        HashMap<ListNode, ListNode> parentMap = new HashMap<>();

        ListNode res = head;
        ListNode par = head;
        ListNode cur = head.next;
        ListNode tail = null;

        while(cur!=null) {
            tail = cur;
            parentMap.put(cur, par);
            par = cur;
            cur=cur.next;
        }


        while(head!=tail && head.next != tail) {
            ListNode tmpTail = parentMap.get(tail);
            ListNode tmpHead = head.next;

            head.next = tail;
            tail.next = tmpHead;

            head = tmpHead;
            tail = tmpTail;
        }

        tail.next = null;

    }
}