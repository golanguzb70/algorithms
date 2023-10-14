#include <iostream>



// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

/*
    Algorithm

1. Start from last gain sum of head.next nodes,  put next = nullptr and return it 
2. When node value equals to 0 it calculating from scratch. 

Be sure that previes sums don't get lost

example: 
head = [0,3,1,0,4,5,2,0]
recursive - 1: 
for->1: sum = 3,  head = 1, 0, 4, 5, 2, 0
for->2: sum = 4, head = 0, 4, 5, 2, 0 -> break because val = 0;
recursive - 2:
for->1: sum = 4, head=5, 2, 0
for->2: sum = 9, head= 2, 0
for->3: sum = 11, head = 0 -> break because val = 0;
recursive - 3:

time complexity -> O(n)
space complexity -> O(number of recursions)
*/
class Solution {
public:
    ListNode* mergeNodes(ListNode* head) {
        int sum = 0;
        head = head->next;
        if (head == nullptr) {
            return head;
        }
        while(head->val != 0) {
            sum+=head->val;
            head = head->next;
        }
        head->val = sum;
        head->next = mergeNodes(head);
        return head;
    }
};

