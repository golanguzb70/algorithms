#include <iostream>

using namespace std;

/*
Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes with a value in the inclusive range [low, high].

Algorithm: I use recursive algorithm to get sum of left and right nodes which are in range of [low, high]

f(root, low, high) {
    if root == nil {
        return 0
    }

    result := 0 
    if root.value >= low && root.value <= high {
        result += root.value;
    }

    return res + f(root.left, low, high) + f(root.right,low, high)
}   

dry run: 
       ROOT   high=15 low=7     response 32
depth 1           10   return 22 + 10
                /    \ 
depth 2       5       15   return 7 + 15
             / \        \    
depth 3     3   7        18    return: 7

*/


struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};


class Solution {
public:
    int rangeSumBST(TreeNode* root, int low, int high) {
        if (root == nullptr) {
            return 0;
        }
        return (root->val >= low && root->val <= high ? root->val:0) + rangeSumBST(root->left, low, high) + rangeSumBST(root->right, low, high);
    }
};


int main() {


    return 0;
}