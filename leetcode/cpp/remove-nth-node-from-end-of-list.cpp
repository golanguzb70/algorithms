/**
 * Definition for singly-linked list.
 */
 
#include <iostream>
#include <vector>
#include <string>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        // 0. reverse the head
        // 1. create link to head
        // 2. Loop through the linked list and if nth node make make the next node to N+1th node


        ListNode* temp = head;
        
    

        if(head->next==nullptr) {
            return head->next;
        }

        ListNode* lastnode=new ListNode(0);
        ListNode* result = lastnode;


        while(head !=nullptr) 
        {
            if(head->val==n){
                head = head->next;
                continue;
            }

            if (lastnode->val==0) {
                lastnode->val=head->val;
            } else {
                ListNode* newNode = new ListNode(head->val);
                lastnode->next = newNode;
                lastnode=lastnode->next;
            }

            head=head->next;
        }
        
        head=result;

        return head;
    }

    ListNode* reverse(ListNode* head, ListNode* last) {
        ListNode* current = new ListNode(head->val);
        if(head->next==nullptr) {
            return current;
        }

        ListNode* last = reverse(head->next, last);
        last->next=current;

        return last;
    }    
};


int main(int argc, char const *argv[])
{
    ListNode* head = new ListNode(1);
    ListNode* current = head;
    
    // Add remaining nodes
    current->next = new ListNode(2);
    current = current->next;
    current->next = new ListNode(3);
    current = current->next;
    current->next = new ListNode(4);
    current = current->next;
    current->next = new ListNode(5);

    // Print the linked list
   

    Solution s;
    ListNode* last;
    // ListNode* result = s.removeNthFromEnd(head, 3);
    ListNode* result = s.reverse(head, last);

    current = last;
    while (current != nullptr) {
        std::cout << current->val << " ";
        current = current->next;
    }
    std::cout << std::endl;

    return 0;
}
