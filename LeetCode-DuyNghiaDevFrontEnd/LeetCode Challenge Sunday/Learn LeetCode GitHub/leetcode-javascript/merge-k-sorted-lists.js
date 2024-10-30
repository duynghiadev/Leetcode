/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode[]} lists
 * @return {ListNode}
 */
var mergeKLists = function (lists) {
    var ret = null;
    for (var i = 0; i < lists.length; i++) {
        ret = mergeTwoLists(ret, lists[i]);
    }
    return ret;
};

var mergeTwoLists = function (l1, l2) {
    var ret = new ListNode(-1);
    var cur = ret;

    while (l1 && l2) {
        if (l1.val <= l2.val) {
            cur.next = l1;
            l1 = l1.next;
        } else {
            cur.next = l2;
            l2 = l2.next;
        }
        cur = cur.next;
    }

    if (l1) {
        cur.next = l1;
    }

    if (l2) {
        cur.next = l2;
    }

    return ret.next;
};